package routes

import (
	"renewit-go/handlers"
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
	router.GET("/users", users.GetUsers)
	router.POST("/users", users.CreateUser)
	router.GET("/materials", materials.GetMaterials)
	router.POST("/materials", materials.CreateMaterials)
	router.PATCH("/materials/:id", materials.UpdateMaterial)
	router.DELETE("/materials/:id", materials.DeleteMaterial)
	router.GET("/catalogue", catalogue.GetCatalogues)
	router.POST("/catalogue", catalogue.CreateCatalogue)
	router.PATCH("/catalogue/:id", catalogue.UpdateCatalogue)
	router.DELETE("/catalogue/:id", catalogue.DeleteCatalogue)
	router.GET("/upcycled-products", upcycledproducts.GetUpcycledProducts)
	router.POST("/upcycled-products", upcycledproducts.CreateUpcycledProducts)
	router.PATCH("/upcycled-products/:id", upcycledproducts.UpdateUpcycledProducts)
	router.DELETE("/upcycled-products/:id", upcycledproducts.DeleteUpcycledProducts)
	router.GET("/payments", payment.GetPayment)
	router.POST("/payments", payment.CreatePayment)
	router.POST("/daraja/stkpush", daraja.STKPush)
	router.POST("/daraja/callback", payment.Callback)
}
