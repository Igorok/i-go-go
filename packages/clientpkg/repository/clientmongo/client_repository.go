package clientmongo

import (
	"context"
	"delivery-go/entities/client/cliententity"
	"delivery-go/entities/client/clientmongo"
	"delivery-go/entities/user/userentity"

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
