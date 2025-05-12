package middlewares

import (
	"authn-service-demo/infrastructure/identity"
	"authn-service-demo/shared/enums"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitFiberMiddlewares(app *fiber.App,
	initPublicRoutes func(app *fiber.App),
	initProtectedRoutes func(app *fiber.App)) {

	app.Use(requestid.New())
	app.Use(logger.New())

	app.Use(func(c *fiber.Ctx) error {
		// get the request id that was add by requestid middleware
		var requestId = c.Locals("requestid")
		// create a new context and add the requestid to it
		var ctx = context.WithValue(context.Background(), enums.ContextKeyRequestId, requestId)

		c.SetUserContext(ctx)
		return c.Next()
	})
	// Route that doesn't require authorization
	initPublicRoutes(app)

	tokenRetrospector := identity.NewIdentityManager()
	app.Use(NewJwtMiddleware(tokenRetrospector))
	// route that require authorization
	initProtectedRoutes(app)

	log.Println("Fiber middlewares initialized")
}
