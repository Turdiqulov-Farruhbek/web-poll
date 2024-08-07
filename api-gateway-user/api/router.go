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
	protected := router.Group("/", middleware.JWTMiddleware())

	protected.POST("/send-answers", h.SendAnswers)

	return router
}
