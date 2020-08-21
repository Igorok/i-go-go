package usercontroller

import (
	"delivery-go/packages/userpkg"
	"delivery-go/web/webgraphql/webtypes/usertype"

	"github.com/graphql-go/graphql"
)

// UserController - web controller
type UserController struct {
	usercase userpkg.UseCase
}

// NewUserController - contructor for UserController
func NewUserController(usercase userpkg.UseCase) *UserController {
	return &UserController{
		usercase: usercase,
	}
}

// SignIn - login
func (controller UserController) SignIn() *graphql.Field {
	return &graphql.Field{
		Type: usertype.SignInType,
		Args: graphql.FieldConfigArgument{
			"login": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			login, ok := p.Args["login"].(string)
			if !ok {
				return nil, userpkg.ErrUserNotFound
			}
			password, ok := p.Args["password"].(string)
			if !ok {
				return nil, userpkg.ErrUserNotFound
			}

			token, err := controller.usercase.SignIn(p.Context, login, password)
			if err != nil {
				return nil, err
			}

			res := make(map[string]string)
			res["token"] = token
			return res, nil
		},
	}
}

// ParseToken - jwt token parsing
func (controller UserController) ParseToken() *graphql.Field {
	return &graphql.Field{
		Type: usertype.UserType,
		Args: graphql.FieldConfigArgument{
			"token": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			token, ok := p.Args["token"].(string)
			if !ok {
				return nil, userpkg.ErrInvalidAccessToken
			}

			user, err := controller.usercase.ParseToken(p.Context, token)
			if err != nil {
				return nil, err
			}

			return user, err
		},
	}
}
