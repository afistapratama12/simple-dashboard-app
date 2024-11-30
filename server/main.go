package main

import (
	"simple-dashboard-server/api/handler"
	"simple-dashboard-server/config"
	"simple-dashboard-server/middleware"
	"simple-dashboard-server/repository"
	"simple-dashboard-server/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

var _ = godotenv.Load()

func main() {
	cfg := config.InitConfig()

	dialer := gomail.NewDialer("smtp.gmail.com", 587, cfg.Env.EmailUsername, cfg.Env.AppPass)

	userRepo := repository.NewUserRepo(cfg.DB)
	notifRepo := repository.NewNotifRepo(cfg.Env, dialer)

	authSvc := service.NewAuthService(cfg.Env, userRepo, notifRepo)
	userSvc := service.NewUserService(userRepo)

	authHandler := handler.NewAuthHandler(authSvc, userSvc)

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	v1 := r.Group("/v1")
	{
		v1.POST("/register", authHandler.Register)
		v1.POST("/login", authHandler.Login)

		v1.POST("/verify-email", authHandler.VerifyEmail)

		v1.GET("/validate-token", middleware.AuthMiddleware(), authHandler.ValidateToken)
		v1.POST("/notif-forgot-password", authHandler.NotifForgotPassword)
		v1.POST("/reset-password", authHandler.ResetPassword)

		v1.PUT("/users/edit", middleware.AuthMiddleware(), authHandler.EditUserLogin)
		v1.GET("/users/profile", middleware.AuthMiddleware(), authHandler.ProfileUserLogin)
	}

	r.Run(":8080")
}
