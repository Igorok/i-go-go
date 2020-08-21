package clientmongo

import (
	"delivery-go/entities/client/cliententity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Email       string             `bson:"email"`
	PhoneNumber string             `bson:"phone_number"`
}

func ToEntityClient(c *Client) *cliententity.Client {
	return &cliententity.Client{
		ID:          c.ID.Hex(),
		Name:        c.Name,
		Email:       c.Email,
		PhoneNumber: c.PhoneNumber,
	}
}

func ToClient(c *cliententity.Client) *Client {
	ID, _ := primitive.ObjectIDFromHex(c.ID)

	return &Client{
		ID:          ID,
		Name:        c.Name,
		Email:       c.Email,
		PhoneNumber: c.PhoneNumber,
	}
}
