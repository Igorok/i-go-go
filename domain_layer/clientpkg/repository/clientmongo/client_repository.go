package clientmongo

import (
	"context"
	"i-go-go/entities_layer/client/cliententity"
	"i-go-go/entities_layer/client/clientmongo"
	"i-go-go/entities_layer/user/userentity"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClietnRepository struct {
	clientCollection *mongo.Collection
}

func NewClientRepository(db *mongo.Database) *ClietnRepository {
	return &ClietnRepository{
		clientCollection: db.Collection("clients"),
	}
}

func (r ClietnRepository) CreateClient(ctx context.Context, user *userentity.UserSystem, cm *cliententity.Client) error {
	cr := clientmongo.ToClient(cm)

	res, err := r.clientCollection.InsertOne(ctx, cr)
	if err != nil {
		return err
	}

	cm.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}
