package productuc

import (
	"authn-service-demo/domain/entities"
	"context"
)

type GetProductsUseCase struct {
	dataStore ProductDataStore
}

func NewGetProductsUseCase(ds ProductDataStore) *GetProductsUseCase {
	return &GetProductsUseCase{
		dataStore: ds,
	}
}

func (uc *GetProductsUseCase) GetProducts(ctx context.Context) []entities.Product {

	return uc.dataStore.GetAll()
}
