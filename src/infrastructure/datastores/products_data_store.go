package datastores

import (
	"authn-service-demo/domain/entities"
	"sort"
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

func (ds *ProductDataStore) GetAll() []entities.Product {
	all := make([]entities.Product, 0, len(ds.Products))
	for _, product := range ds.Products {
		all = append(all, product)
	}
	// Sort all slice in asceding order according to date of creating
	// You just need to to define the sorting login in the nameles function and sort.Slice() will take it from here
	sort.Slice(all, func(i, j int) bool {
		return all[i].CreatedAt.Before(all[j].CreatedAt)
	})
	return all
}
