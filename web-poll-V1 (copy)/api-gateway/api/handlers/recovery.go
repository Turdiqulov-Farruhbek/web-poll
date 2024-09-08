package handlers

import (
	"auth-service/config"
	"auth-service/models"
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"regexp"
	"time"

	pb "auth-service/genprotos"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
)

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

func (h *HTTPHandler) SendConfirmationCode(email string) error {
	code, err := generateConfirmationCode()
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.Load().SENDER_EMAIL)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Tasdiqlash kodi")
	m.SetBody("text/plain", fmt.Sprintf("Sizning tasdiqlash kodingiz: %d", code))

	d := gomail.NewDialer("smtp.gmail.com", 587, config.Load().SENDER_EMAIL, config.Load().APP_PASSWORD)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	err = rdb.Set(context.Background(), email, code, 3*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("server error storing confirmation code in Redis", "dtc", err.Error())
	}
	return nil
}

func generateConfirmationCode() (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}

// ForgotPassword godoc
// @Summary Forgot passwrod
// @Description Sends a confirmation code to email recovery password
// @Tags password-recovery
// @Accept json
// @Produce json
// @Param credentials body models.ForgotPasswordReq true "User login credentials"
// @Success 200 {object} string ""
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Page not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /forgot-password [POST]
func (h *HTTPHandler) ForgotPassword(c *gin.Context) {
	var req models.ForgotPasswordReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	if !config.IsValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	user, err := h.User.Profile(context.Background(), &pb.GetProfileReq{Email: req.Email})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "details": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	err = h.SendConfirmationCode(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error sending confirmation code to email", "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Confirmation code sent to your email. Please use your code within 3 minutes."})
}

// RecoverPassword godoc
// @Summary Recover password (Use this one after sending verification code)
// @Description Verifies the code and updates the password
// @Tags password-recovery
// @Accept json
// @Produce json
// @Param request body models.RecoverPasswordReq true "Recover Password Request"
// @Success 200 {object} string "Password successfully updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Incorrect verification code"
// @Failure 404 {object} string "Verification code expired or email not found"
// @Failure 500 {object} string "Error updating password"
// @Router /recover-password [post]
func (h *HTTPHandler) RecoverPassword(c *gin.Context) {
	var req models.RecoverPasswordReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	if req.Email == "" || req.Code == "" || req.NewPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email, code, and new password are required fields."})
		return
	}

	if len(req.NewPassword) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters long"})
		return
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(req.NewPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must contain at least one number"})
		return
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(req.NewPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must contain at least one uppercase letter"})
		return
	}
	if !regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(req.NewPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must contain at least one special character"})
		return
	}

	storedCode, err := rdb.Get(context.Background(), req.Email).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Verification code expired or email not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "This email not found in a recovery requests!"})
		return
	}

	if storedCode != req.Code {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect verification code"})
		return
	}

	_, err = h.User.UpdatePassword(context.Background(), &pb.UpdatePasswordReq{Email: req.Email, NewPassword: req.NewPassword})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating password", "details": err.Error()})
		return
	}
	rdb.Del(context.Background(), req.Email)

	c.JSON(http.StatusOK, gin.H{"message": "Password successfully updated"})
}

// SendCodeAgain godoc
// @Summary Sends code again if user didn't recieve the code
// @Description Sends a confirmation code to email recovery password again
// @Tags password-recovery
// @Accept json
// @Produce json
// @Param credentials body pb.ByEmail true "User login credentials"
// @Success 200 {object} string ""
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Page not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /resverify [POST]
func (h *HTTPHandler) SendCodeAgain(c *gin.Context) {
	var req pb.ByEmail
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	if !config.IsValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	user, err := h.User.GetByEmail(c, &pb.ByEmail{Email: req.Email})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "details": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	err = h.SendConfirmationCode(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error sending confirmation code to email", "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Confirmation code sent to your email. Please use your code within 3 minutes."})
}
