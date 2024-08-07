package api

import (
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

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// #################### AUTH SERVICE ######################### //
	router.POST("/register", h.Register)
	router.POST("/confirm-registration", h.ConfirmRegistration)
	router.POST("/login", h.Login)
	router.POST("/forgot-password", h.ForgotPassword)
	router.POST("/recover-password", h.RecoverPassword)

	protected := router.Group("/", middleware.JWTMiddleware())
	protected.GET("/profile", h.Profile)
	router.GET("/user/:id", h.GetByID)

	// #################### POLLING SERVICE ######################### //
	for_admin := protected.Group("/", middleware.IsAdminMiddleware())

	for_admin.GET(("/results"), h.GetUserResultsInExcel)

	poll := for_admin.Group("/poll")
	{
		poll.POST("/", h.CreatePoll)
		poll.PUT("/", h.UpdatePoll)
		poll.DELETE("/:id", h.DeletePoll)
		poll.GET("/:id", h.GetPollByID)
	}

	for_admin.GET("/polls", h.GetAllPolls)

	question := for_admin.Group("/question")
	{
		question.POST("/", h.CreateQuestion)
		question.PUT("/", h.UpdateQuestion)
		question.DELETE("/:id", h.DeleteQuestion)
		question.GET("/:id", h.GetQuestionByID)
	}
	for_admin.GET("/questions", h.GetAllQuestions)

	return router
}
