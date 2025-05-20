package handlers

import (
	"context"
	"errors"

	"authn-service-demo/use_cases/usermgmtuc"

	"github.com/gofiber/fiber/v2"
)

type RegisterUseCase interface {
	Register(context.Context, usermgmtuc.RegisterRequest) (*usermgmtuc.RegisterResponse, error)
}

func RegisterHandler(useCase RegisterUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req = usermgmtuc.RegisterRequest{}
		var ctx = c.UserContext()

		err := c.BodyParser(&req)
		if err != nil {
			return errors.Join(err, errors.New("[handlers]: unable to parse request"))
		}

		res, err := useCase.Register(ctx, req)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusCreated).JSON(res)
	}

}
