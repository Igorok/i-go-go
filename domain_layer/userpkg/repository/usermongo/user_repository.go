package usermongo

import (
	"context"
	"i-go-go/entities_layer/user/userentity"
	"i-go-go/entities_layer/user/usermongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository mongo db repository
type UserRepository struct {
	collUsers *mongo.Collection
}

// NewUserRepository - constructor for UserRepository
func NewUserRepository(db *mongo.Database, collName string) *UserRepository {
	return &UserRepository{
		collUsers: db.Collection(collName),
	}
}

// GetUser - get user by login and password
func (r UserRepository) GetUser(ctx context.Context, login, password string) (*userentity.UserSystem, error) {
	user := new(usermongo.UserSystem)
	err := r.collUsers.FindOne(ctx, bson.M{
		"login":    login,
		"password": password,
	}).Decode(user)

	if err != nil {
		return nil, err
	}

	return usermongo.ToEntityUserSystem(user), nil
}
