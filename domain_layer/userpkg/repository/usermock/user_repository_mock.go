package usermock

import (
	"context"
	"delivery-go/bin/testdata"
	"delivery-go/entities_layer/user/userentity"
)

// UserRepositoryMock
type UserRepositoryMock struct{}

// GetUser - get user by login and password
func (r UserRepositoryMock) GetUser(ctx context.Context, login, password string) (*userentity.UserSystem, error) {
	umArr := testdata.GetTestUsers()
	um := userentity.UserSystem{}

	for _, u := range umArr {
		if u.Login == login && u.Password == password {
			um = u
			break
		}
	}

	return &um, nil
}
