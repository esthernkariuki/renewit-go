package main

import (
	"log"
	"renewit-go/config"
	"renewit-go/database"
	"renewit-go/internal/catalogue"
	"renewit-go/internal/materials"
	"renewit-go/internal/payment"
	"renewit-go/internal/upcycledproducts"
	"renewit-go/internal/users"
	"renewit-go/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.LoadEnv()
	router := gin.Default()
	database.ConnectDatabase()
	database.DB.AutoMigrate(
		&users.User{},
		&materials.Material{},
		&catalogue.Catalogue{},
		&upcycledproducts.Upcycledproduct{},
		&payment.Payment{},
	)
	routes.RegisterRoutes(router)
	router.Run(":8080")

}
