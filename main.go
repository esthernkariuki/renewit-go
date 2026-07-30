package main

import (
	"renewit-go/config"
	"renewit-go/database"
	"renewit-go/internal/catalogue"
	"renewit-go/internal/materials"
	"renewit-go/internal/payment"
	"renewit-go/internal/upcycledproducts"
	"renewit-go/internal/users"
	"renewit-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
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
