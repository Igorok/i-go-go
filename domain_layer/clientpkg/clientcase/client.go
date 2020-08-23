package clientcase

import (
	"context"
	"i-go-go/entities_layer/client/cliententity"
	"i-go-go/entities_layer/user/userentity"
	"i-go-go/domain_layer/clientpkg"
)

type ClientUseCase struct {
	clientRepo clientpkg.Repository
}

func NewClientUseCase(clientRepo clientpkg.Repository) *ClientUseCase {
	return &ClientUseCase{
		clientRepo: clientRepo,
	}
}

func (c ClientUseCase) CreateClient(ctx context.Context, user *userentity.UserSystem, name, email, phoneNumber string) error {
	cm := &cliententity.Client{
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
	}

	return c.clientRepo.CreateClient(ctx, user, cm)
}
