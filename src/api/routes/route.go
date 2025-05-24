package routes

import (
	"authn-service-demo/api/handlers"
	"authn-service-demo/api/middlewares"
	"authn-service-demo/infrastructure/datastores"
	"authn-service-demo/infrastructure/identity"
	"authn-service-demo/use_cases/productuc"
	"authn-service-demo/use_cases/usermgmtuc"

	"github.com/gofiber/fiber/v2"
)

func InitPublicRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to My Application backend. Yes Again, sorry, i'm testing FluxCD webhook reciever.  ༼つಠ益ಠ༽つ ─=≡ΣO))"))
	})

	grp := app.Group("/api/v1")

	identityManager := identity.NewIdentityManager()
	registerUseCase := usermgmtuc.NewRegisterUseCase(identityManager)

	grp.Post("/user", handlers.RegisterHandler(registerUseCase))
}

func InitProtectedRoute(app *fiber.App) {
	grp := app.Group("/api/v1")

	productDataStore := datastores.NewProductDataStore()
	createProductUseCase := productuc.NewCreateProductUseCase(productDataStore)
	grp.Post("/products", middlewares.NewRequiresRealmRole("admin"),
		handlers.CreateProductHandler(createProductUseCase))

	getProductsUseCase := productuc.NewGetProductsUseCase(productDataStore)
	grp.Get("/products", middlewares.NewRequiresRealmRole("viewer"),
		handlers.GetProductsHandler(getProductsUseCase))
}
