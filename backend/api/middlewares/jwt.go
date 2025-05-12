package middlewares

import (
	"authn-service-demo/infrastructure/config"
	"authn-service-demo/shared/enums"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
	contribJwt "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	golangJwt "github.com/golang-jwt/jwt/v5"
)

type TokenRetrospector interface {
	RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error)
}

func NewJwtMiddleware(tokenRetrospector TokenRetrospector) fiber.Handler {
	base64str := config.GlobalConfigParams.Keycloak.RealmRS256PublickKey
	publicKey, err := ParseKeycloakPublicKey(base64str)
	if err != nil {
		panic(err)
	}
	return contribJwt.New(contribJwt.Config{
		SigningKey: contribJwt.SigningKey{
			JWTAlg: contribJwt.RS256,
			Key:    publicKey,
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			return SuccessHandler(c, tokenRetrospector)
		},
	})
}

// Check if Keycloak public key is a valid RSA public key
func ParseKeycloakPublicKey(base64str string) (*rsa.PublicKey, error) {
	buf, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		return nil, err
	}

	parseKey, err := x509.ParsePKIXPublicKey(buf)
	if err != nil {
		return nil, err
	}
	// rsa.PublicKey type assertion because paseKey is of type interface
	publicKey, ok := parseKey.(*rsa.PublicKey)

	if ok {
		return publicKey, nil
	}

	return nil, fmt.Errorf("[middlewares]unexpected key type %T", publicKey)
}

func SuccessHandler(c *fiber.Ctx, tokenRetrospector TokenRetrospector) error {
	// jwtToken is stored in the fiber req context and refered to by the key "users"
	// This assume that jwtToken had previously been set in the fiber req context
	jwtToken := c.Locals("user").(*golangJwt.Token)
	claims := jwtToken.Claims.(golangJwt.MapClaims)

	var ctx = c.UserContext()
	var contextWithClaims = context.WithValue(ctx, enums.ContextKeyClaims, claims)
	c.SetUserContext(contextWithClaims)

	rptResult, err := tokenRetrospector.RetrospectToken(ctx, jwtToken.Raw)
	if err != nil {
		panic(err)
	}
	if !*rptResult.Active {
		return c.Status(fiber.StatusUnauthorized).SendString("token is not active")
	}

	return c.Next()

}
