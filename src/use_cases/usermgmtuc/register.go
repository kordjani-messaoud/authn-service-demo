package usermgmtuc

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"github.com/Nerzal/gocloak/v13"
)

type RegisterRequest struct {
	Username     string
	Password     string
	FirstName    string
	LastName     string
	Email        string
	MobileNumber string
}

type RegisterResponse struct {
	User *gocloak.User
}

type RegisterUseCase struct {
	identityManager IdentityManager
}

func NewRegisterUseCase(im IdentityManager) *RegisterUseCase {

	return &RegisterUseCase{
		im,
	}
}

func (uc *RegisterUseCase) Register(ctx context.Context, req RegisterRequest) (*RegisterResponse, error) {

	valid, err := uc.ValidRegisterRequest(req)

	if err != nil {
		return nil, errors.New("[usermgmtuc]: error while register request validation")
	}
	if !valid {
		return nil, errors.New("[usermgmtuc]: register request not valid")
	}

	var user = gocloak.User{
		Username:      gocloak.StringP(req.Username),
		FirstName:     gocloak.StringP(req.FirstName),
		LastName:      gocloak.StringP(req.LastName),
		Email:         gocloak.StringP(req.Email),
		EmailVerified: gocloak.BoolP(true),
		Enabled:       gocloak.BoolP(true),
		Attributes:    &map[string][]string{},
	}

	if strings.TrimSpace(req.MobileNumber) != "" {
		(*user.Attributes)["mobile"] = []string{req.MobileNumber}
	}

	userRes, err := uc.identityManager.CreateUser(ctx, user, req.Password, "viewer")

	return &RegisterResponse{userRes}, err
}

func (uc *RegisterUseCase) ValidRegisterRequest(req RegisterRequest) (bool, error) {

	if len(req.Username) > 0 &&
		len(req.Username) <= 15 &&
		len(req.Password) > 0 &&
		len(req.FirstName) >= 0 &&
		len(req.LastName) <= 30 &&
		len(req.Email) >= 0 {
		if len(req.Email) > 0 {
			valid, err := regexp.Match(`[a-zA-Z0-9+._%+=-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, []byte(req.Email))
			if err != nil {
				return false, err
			}
			return valid, nil
		}
		return true, nil

	}
	return false, nil
}
