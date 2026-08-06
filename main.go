// @title RenewIt API
// @version 1.0
// @description REST API for RenewIt, connecting traders, upcyclers, and buyers.
// @termsOfService http://swagger.io/terms/

// @contact.name Esther Nyambura Kariuki
// @contact.email esthernnyamburaa@gmail.com

// @license.name MIT

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"fmt"
	"log"
	"os"
	"renewit-go/config"
	"renewit-go/database"
	"renewit-go/internal/catalogue"
	"renewit-go/internal/materials"
	"renewit-go/internal/payment"
	"renewit-go/internal/upcycledproducts"
	"renewit-go/internal/users"
	"renewit-go/routes"
	"strings"
	"time"

	_ "renewit-go/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using environment variables")
	}
	config.LoadEnv()

	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	database.ConnectDatabase()
	database.DB.AutoMigrate(
		&users.User{},
		&materials.Material{},
		&catalogue.Catalogue{},
		&upcycledproducts.Upcycledproduct{},
		&payment.Payment{},
	)
	routes.RegisterRoutes(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(fmt.Sprintf(":%s", port))

}
