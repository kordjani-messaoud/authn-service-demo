package middlewares

import (
	"authn-service-demo/shared/enums"
	"authn-service-demo/shared/jwt"

	"github.com/gofiber/fiber/v2"
	golangJwt "github.com/golang-jwt/jwt/v5"
)

func NewRequiresRealmRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()
		claims := ctx.Value(enums.ContextKeyClaims).(golangJwt.MapClaims)
		jwtHelper := jwt.NewJwtHelper(claims)
		if !jwtHelper.IsUserInRealmRole(role) {
			return c.Status(fiber.StatusForbidden).SendString("[middlewares]role authorization failed")
		}
		return c.Next()
	}
}
