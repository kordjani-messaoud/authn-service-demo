package handlers

import (
	"authn-service-demo/domain/entities"
	"authn-service-demo/use_cases/productuc"
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, req productuc.CreateProductRequest) (*productuc.CreateProductResponse, error)
}

type GetProductsUseCase interface {
	GetProducts(ctx context.Context) []entities.Product
}

func CreateProductHandler(useCase CreateProductUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()

		var req = productuc.CreateProductRequest{}

		err := c.BodyParser(&req)
		if err != nil {
			return errors.Join(err, errors.New("[producthandler] could not parse req"))
		}

		response, err := useCase.CreateProduct(ctx, req)

		if err != nil {
			return errors.Join(err, errors.New("[producthandler] could not create product"))
		}

		return c.Status(fiber.StatusCreated).JSON(response)
	}
}

func GetProductsHandler(useCase GetProductsUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()
		products := useCase.GetProducts(ctx)

		return c.Status(fiber.StatusOK).JSON(products)
	}
}
