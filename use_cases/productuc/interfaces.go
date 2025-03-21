package productuc

import (
	"authn-service-demo/domain/entities"
)

type ProductDataStore interface {
	GetAll() []entities.Product
	Store(product *entities.Product) error
}
