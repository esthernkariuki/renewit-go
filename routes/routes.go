package routes

import (
	"renewit-go/handlers"
	"renewit-go/internal/auth"
	"renewit-go/internal/catalogue"
	"renewit-go/internal/daraja"
	"renewit-go/internal/materials"
	"renewit-go/internal/payment"
	"renewit-go/internal/upcycledproducts"
	"renewit-go/internal/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/", handlers.Home)

	// Auth routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", auth.RegisterHandler)
		authRoutes.POST("/login", auth.LoginHandler)
	}

	// User routes
	router.GET("/users", users.GetUsers)
	router.POST("/users", users.CreateUser)
	router.PATCH("/users/:id", users.UpdateUser)
	router.DELETE("/users/:id", users.DeleteUser)

	// Material routes
	router.GET("/materials", materials.GetMaterials)
	router.POST("/materials", auth.AuthMiddleware(), auth.RequireRole("Trader"), materials.CreateMaterials)
	router.PATCH("/materials/:id", auth.AuthMiddleware(), materials.UpdateMaterial)
	router.DELETE("/materials/:id", auth.AuthMiddleware(), materials.DeleteMaterial)

	// Catalogue routes
	router.GET("/catalogue", catalogue.GetCatalogues)
	router.POST("/catalogue", catalogue.CreateCatalogue)
	router.PATCH("/catalogue/:id", catalogue.UpdateCatalogue)
	router.DELETE("/catalogue/:id", catalogue.DeleteCatalogue)

	// Upcycled products
	router.GET("/upcycled-products", upcycledproducts.GetUpcycledProducts)
	router.POST("/upcycled-products", auth.AuthMiddleware(), auth.RequireRole("Upcycler"), upcycledproducts.CreateUpcycledProducts)
	router.PATCH("/upcycled-products/:id", auth.RequireRole("Upcycler"), upcycledproducts.UpdateUpcycledProducts)
	router.DELETE("/upcycled-products/:id", auth.RequireRole("Upcycler"), upcycledproducts.DeleteUpcycledProducts)

	// Payments
	router.GET("/payments", payment.GetPayment)
	router.POST("/payments", payment.CreatePayment)

	// Daraja
	router.POST("/daraja/stkpush", daraja.STKPush)
	router.POST("/daraja/callback", payment.Callback)
}
