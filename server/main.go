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

	_ "simple-dashboard-server/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var _ = godotenv.Load()

// @title Service / Server Simple Dashboard
// @version 1.0
// @description This is a service / server simple dashboard
// @contact.name Afista pratama
// @contact.url https://linkedin.com/in/afistapratama
// @contact.email pratama.otori.12@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @value Bearer <token>
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

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": gin.H{
				"server": true,
				"database": func() bool {
					sqlDB, err := cfg.DB.DB()
					if err != nil {
						return false
					}

					err = sqlDB.Ping()
					return err == nil
				}(),
			},
		})
	})

	// add swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
