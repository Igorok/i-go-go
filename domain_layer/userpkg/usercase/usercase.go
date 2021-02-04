package usercase

import (
	"context"
	"fmt"
	"i-go-go/domain_layer/userpkg"
	"i-go-go/entities_layer/user/userentity"
	"i-go-go/service_layer"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *userentity.UserSystem `json:"user"`
}

type UserCase struct {
	userRepo       userpkg.Repository
	hashSalt       string
	signingKey     string
	signingKeyByte []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo userpkg.Repository,
	hashSalt string,
	signingKey string,
	tokenTTLSeconds time.Duration) *UserCase {

	return &UserCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		signingKeyByte: []byte(signingKey),
		expireDuration: time.Second * tokenTTLSeconds,
	}
}

func (uc *UserCase) SignIn(ctx context.Context, login, password string) (string, error) {
	password = service_layer.HashPwd(password, uc.hashSalt)

	user, err := uc.userRepo.GetUser(ctx, login, password)
	if err != nil {
		return "", userpkg.ErrUserNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(uc.expireDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(uc.signingKeyByte)
}

func (uc *UserCase) ParseToken(ctx context.Context, accessToken string) (*userentity.UserSystem, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return uc.signingKeyByte, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, userpkg.ErrInvalidAccessToken
}
