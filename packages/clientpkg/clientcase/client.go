package clientcase

import (
	"context"
	"delivery-go/entities/client/cliententity"
	"delivery-go/entities/user/userentity"
	"delivery-go/packages/clientpkg"
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
