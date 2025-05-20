package usermgmtuc

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type IdentityManager interface {
	CreateUser(ctx context.Context, user gocloak.User, password string, role string) (*gocloak.User, error)
}
