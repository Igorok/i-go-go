package usercase_test

import (
	"context"
	"delivery-go/packages/userpkg/repository/usermock"
	"delivery-go/packages/userpkg/usercase"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("../../../utils")
	viper.ReadInConfig()

	userRepo := usermock.UserRepositoryMock{}
	uc := usercase.NewAuthUseCase(
		userRepo,
		viper.GetString("app.hash_salt"),
		[]byte(viper.GetString("app.signing_key")),
		viper.GetDuration("app.token_ttl"),
	)

	// SignIn test
	token, err := uc.SignIn(context.TODO(), "courier_pizza", "courier_pizza")
	assert.Nil(t, err)
	assert.NotEqual(t, token, "")

	// ParseToken test (ctx context.Context, accessToken string)
	umLink, err := uc.ParseToken(context.TODO(), token)
	assert.Nil(t, err)
	assert.NotNil(t, umLink)
	assert.Equal(t, (*umLink).Login, "courier_pizza")
	assert.Equal(t, (*umLink).ID, "5e874f4b327272d07e537a50")
}
