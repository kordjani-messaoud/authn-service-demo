package datastores

import (
	"authn-service-demo/domain/entities"
	"sync"

	"github.com/google/uuid"
)

type ProductDataStore struct {
	Products map[uuid.UUID]entities.Product
	sync.Mutex
}

func NewProductDataStore() *ProductDataStore {
	return &ProductDataStore{
		Products: make(map[uuid.UUID]entities.Product),
	}
}

func (ds *ProductDataStore) Store(product *entities.Product) error {
	ds.Lock()
	ds.Products[product.Id] = *product
	ds.Unlock()

	return nil
}
