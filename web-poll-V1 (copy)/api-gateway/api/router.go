package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "auth-service/api/docs"
	"auth-service/api/handlers"
	"auth-service/api/middleware"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(PollConn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	h := handlers.NewHandler(PollConn)

	router.Use(CORSMiddleware())

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// #################### AUTH SERVICE ######################### //
	router.POST("/register", h.Register)
	router.POST("/confirm-registration", h.ConfirmRegistration)
	router.POST("/login", h.Login)
	router.POST("/forgot-password", h.ForgotPassword)
	router.POST("/recover-password", h.RecoverPassword)
	router.POST("/resverify", h.SendCodeAgain)

	protected := router.Group("/", middleware.JWTMiddleware())
	protected.GET("/profile", h.Profile)
	router.GET("/user/:id", h.GetByID)

	// #################### USER SERVICE ######################### //
	for_user := protected.Group("/", middleware.IsUserMiddleware())
	for_user.POST("/send-answers", h.SendAnswers)

	// #################### POLLING SERVICE ######################### //
	for_admin := protected.Group("/", middleware.IsAdminMiddleware())
	for_admin_or_user := protected.Group("/", middleware.IsAdminOrUserMiddleware())

	for_admin.GET(("/results"), h.GetUserResultsInExcel)

	poll := for_admin.Group("/poll")
	{
		poll.POST("/", h.CreatePoll)
		poll.PUT("/:id", h.UpdatePoll)
		poll.DELETE("/:id", h.DeletePoll)
		for_admin_or_user.GET("/poll/:id", h.GetPollByID)
	}

	for_admin_or_user.GET("/polls", h.GetAllPolls)

	question := for_admin.Group("/question")
	{
		question.POST("/", h.CreateQuestion)
		question.PUT("/", h.UpdateQuestion)
		question.DELETE("/:id", h.DeleteQuestion)
		for_admin_or_user.GET("/question/:id", h.GetQuestionByID)
	}
	for_admin_or_user.GET("/questions/:poll_id", h.GetAllQuestions)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
