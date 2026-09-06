package handler

import (
	"api_gateway_microservice/internal/client"
	jwt_lib "api_gateway_microservice/internal/lib/jwt"
	auth_service "api_gateway_microservice/internal/service/auth"
	current_user_service "api_gateway_microservice/internal/service/current_user"
	user_service "api_gateway_microservice/internal/service/user"
	auth_handler "api_gateway_microservice/internal/transport/rest/handler/auth"
	current_user_handler "api_gateway_microservice/internal/transport/rest/handler/current_user"
	user_handler "api_gateway_microservice/internal/transport/rest/handler/user"
	http_middleware "api_gateway_microservice/internal/transport/rest/middleware"
	"api_gateway_microservice/pkg/logger"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func NewRouter(
	e *echo.Echo,
	log *logger.Logger,
	jwtManager *jwt_lib.JWTManager,
	grpcClients *client.GrpcClients,
) {
	// Services
	authService := auth_service.New(log, jwtManager, grpcClients)
	userService := user_service.New(log, grpcClients)
	currentUserService := current_user_service.New(log, grpcClients)

	// Handlers
	authHandler := auth_handler.New(log, authService)
	userHandler := user_handler.New(log, userService)
	currentUserHandler := current_user_handler.New(log, currentUserService)

	jwtMiddleware := http_middleware.NewJwtMiddleware(log, jwtManager)
	e.Use(jwtMiddleware.Process)

	// Base middlewares
	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLogger())

	// Auth
	e.POST("/api/auth/login", authHandler.Login)
	e.POST("/api/auth/register", authHandler.Register)
	e.DELETE("/api/auth/logout", authHandler.Logout)

	// User
	e.GET("/api/users/:id", userHandler.GetUserById)

	// Current User
	e.GET("/api/me", currentUserHandler.GetCurrentUser)
}
