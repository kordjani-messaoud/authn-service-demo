package productuc

import (
	"authn-service-demo/domain/entities"
)

type ProductDataStore interface {
	Store(product *entities.Product) error
}
