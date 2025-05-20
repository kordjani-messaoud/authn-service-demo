package productuc

import (
	"authn-service-demo/domain/entities"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CreateProductRequest struct {
	Name  string
	Price float32
}

type CreateProductResponse struct {
	Product *entities.Product
}

type CreateProductUseCase struct {
	DataStore ProductDataStore
}

func NewCreateProductUseCase(ds ProductDataStore) *CreateProductUseCase {
	return &CreateProductUseCase{
		DataStore: ds,
	}
}

func (uc *CreateProductUseCase) CreateProduct(ctx context.Context,
	req CreateProductRequest) (*CreateProductResponse, error) {
	err := uc.validate(req)
	if err != nil {
		return nil, err
	}

	var product = entities.Product{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
		Name:      req.Name,
		Price:     req.Price,
	}

	err = uc.DataStore.Store(&product)

	return &CreateProductResponse{Product: &product}, err
}

func (uc *CreateProductUseCase) validate(req CreateProductRequest) error {
	if len(req.Name) >= 3 &&
		len(req.Name) <= 15 &&
		req.Price > 0 {
		return nil
	} else {
		return fmt.Errorf("[productuc] create product request not valid: name: %v, price: %v ", req.Name, req.Price)
	}
}
