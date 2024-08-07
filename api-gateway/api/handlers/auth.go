package handlers

import (
	"auth-service/api/token"
	"auth-service/config"
	"auth-service/models"
	"context"
	"net/http"

	pb "auth-service/genprotos"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	_ "github.com/swaggo/swag"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email, username, and password
// @Tags registration
// @Accept json
// @Produce json
// @Param user body pb.RegisterReqForSwagger true "User registration request"
// @Success 201 {object} token.Tokens "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /register [post]
func (h *HTTPHandler) Register(c *gin.Context) {
	var req pb.RegisterReqForSwagger
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	if !config.IsValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	if _, err := h.User.IsEmailExists(c, &pb.IsEmailExistsReq{Email: req.Email}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered", "details": err.Error()})
		return
	}

	if err := config.IsValidPassword(req.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := config.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "err": err.Error()})
	}

	req.Password = string(hashedPassword)

	var gender int
	if req.Gender == "male" {
		gender = 0
	} else if req.Gender == "female" {
		gender = 1
	} else {
		c.JSON(400, "Invalid gender format")
	}

	_, err = h.User.Register(c,
		&pb.RegisterReq{
			Password:          req.Password,
			Name:              req.Name,
			Email:             req.Email,
			Surname:           req.Surname,
			Gender:            pb.GenderType(gender),
			PhoneNumber:       req.PhoneNumber,
			WorkingExperience: req.WorkingExperience,
			LevelType:             req.LevelType,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "err": err.Error()})
		return
	}

	err = h.SendConfirmationCode(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error sending confirmation code", "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Your account has been registered. Please check your email for a confirmation link. You have 3 minutes to confirm your account."})
}

// ConfirmRegistration godoc
// @Summary Confirm registration with code
// @Description Confirms a user's registration using the code sent to their email.
// @Tags registration
// @Accept json
// @Produce json
// @Param confirmation body models.ConfirmRegistrationReq true "Confirmation request"
// @Success 200 {object} token.Tokens "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Incorrect verification code"
// @Failure 404 {object} string "Verification code expired or email not found"
// @Router /confirm-registration [post]
func (h *HTTPHandler) ConfirmRegistration(c *gin.Context) {
	var req models.ConfirmRegistrationReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	storedCode, err := rdb.Get(context.Background(), req.Email).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Verification code expired or email not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "This email not found in confirmation requests!"})
		return
	}

	if storedCode != req.Code {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect verification code"})
		return
	}

	_, err = h.User.ConfirmUser(c, &pb.ConfirmUserReq{Email: req.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error confirming user", "details": err.Error()})
		return
	}

	user, err := h.User.Profile(c, &pb.GetProfileReq{Email: req.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user", "details": err.Error()})
		return
	}

	tokens := token.GenerateJWTToken(user.Id, user.Email, user.Role)

	rdb.Del(context.Background(), req.Email)

	c.JSON(http.StatusOK, tokens)
}

// Login godoc
// @Summary Login a user
// @Description Authenticate user with email and password
// @Tags login
// @Accept json
// @Produce json
// @Param credentials body models.LoginReq true "User login credentials"
// @Success 200 {object} token.Tokens "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Invalid email or password"
// @Router /login [post]
func (h *HTTPHandler) Login(c *gin.Context) {
	req := models.LoginReq{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	user, err := h.User.Profile(c, &pb.GetProfileReq{Email: req.Email})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User registered with this email not found"})
		return
	}

	if !config.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !user.IsConfirmed {
		err = h.SendConfirmationCode(req.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error sending confirmation code", "err": err.Error()})
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Your account is not verified. Please check your email for a confirmation link."})
		return
	}

	if user.Role == "banned" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are banned"})
		return
	}

	tokens := token.GenerateJWTToken(user.Id, user.Email, user.Role)

	c.JSON(http.StatusOK, tokens)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get the profile of the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.GetProfileResp
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /profile [get]
func (h *HTTPHandler) Profile(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	email := claims.(jwt.MapClaims)["email"].(string)
	user, err := h.User.Profile(c, &pb.GetProfileReq{Email: email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Server error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *HTTPHandler) GetByID(c *gin.Context) {
	id := &models.GetProfileByIdReq{ID: c.Param("id")}
	user, err := h.User.GetByID(c, &pb.GetProfileByIdReq{Id: id.ID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Couldn't get the user": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
