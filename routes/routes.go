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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/", handlers.Home)
	router.Static("/uploads", "./uploads")

	// Swagger
	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	// =========================
	// AUTH
	// =========================

	authRoutes := router.Group("/auth")

	{
		authRoutes.POST(
			"/register",
			auth.RegisterHandler,
		)

		authRoutes.POST(
			"/login",
			auth.LoginHandler,
		)
	}

	// =========================
	// USERS
	// =========================

	router.GET(
		"/users",
		users.GetUsers,
	)

	router.POST(
		"/users",
		users.CreateUser,
	)

	router.PATCH(
		"/users/:id",
		users.UpdateUser,
	)

	router.DELETE(
		"/users/:id",
		users.DeleteUser,
	)

	// =========================
	// MATERIALS
	// =========================

	router.GET(
		"/materials",
		materials.GetMaterials,
	)

	router.POST(
		"/materials",
		auth.AuthMiddleware(),
		auth.RequireRole("trader"),
		materials.CreateMaterials,
	)

	router.PATCH(
		"/materials/:id",
		auth.AuthMiddleware(),
		materials.UpdateMaterial,
	)

	router.DELETE(
		"/materials/:id",
		auth.AuthMiddleware(),
		materials.DeleteMaterial,
	)

	// =========================
	// CATALOGUE
	// =========================

	router.GET(
		"/catalogue",
		catalogue.GetCatalogues,
	)

	router.POST(
		"/catalogue",
		catalogue.CreateCatalogue,
	)

	router.PATCH(
		"/catalogue/:id",
		catalogue.UpdateCatalogue,
	)

	router.DELETE(
		"/catalogue/:id",
		catalogue.DeleteCatalogue,
	)

	// =========================
	// UPCYCLED PRODUCTS
	// =========================

	// Upcycled products
	router.GET(
		"/upcycled-products",
		upcycledproducts.GetUpcycledProducts,
	)

	router.POST(
		"/upcycled-products",
		auth.AuthMiddleware(),
		auth.RequireRole("upcycler"),
		upcycledproducts.CreateUpcycledProducts,
	)

	router.PATCH(
		"/upcycled-products/:id",
		auth.AuthMiddleware(),
		auth.RequireRole("upcycler"),
		upcycledproducts.UpdateUpcycledProducts,
	)

	router.DELETE(
		"/upcycled-products/:id",
		auth.AuthMiddleware(),
		auth.RequireRole("upcycler"),
		upcycledproducts.DeleteUpcycledProducts,
	)
	// =========================
	// PAYMENTS
	// =========================

	router.GET(
		"/payments",
		payment.GetPayment,
	)

	router.POST(
		"/payments",
		auth.AuthMiddleware(),
		payment.CreatePayment,
	)

	// =========================
	// DARAJA
	// =========================

	router.POST(
		"/daraja/stkpush",
		daraja.STKPush,
	)

	router.POST(
		"/daraja/callback",
		payment.Callback,
	)
}
