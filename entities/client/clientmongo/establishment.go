package clientmongo

import (
	"delivery-go/entities/client/cliententity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Establishment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	IDCl        primitive.ObjectID `bson:"_id_client"`
	Name        string             `bson:"name"`
	Email       string             `bson:"email"`
	PhoneNumber string             `bson:"phone_number"`
}

func ToEntityEstablishment(e *Establishment) *cliententity.Establishment {
	return &cliententity.Establishment{
		ID:          e.ID.Hex(),
		IDCl:        e.ID.Hex(),
		Name:        e.Name,
		Email:       e.Email,
		PhoneNumber: e.PhoneNumber,
	}
}

func toEstablishment(em *cliententity.Establishment) *Establishment {
	ID, _ := primitive.ObjectIDFromHex(em.ID)
	clientID, _ := primitive.ObjectIDFromHex(em.IDCl)

	return &Establishment{
		ID:          ID,
		IDCl:        clientID,
		Name:        em.Name,
		Email:       em.Email,
		PhoneNumber: em.PhoneNumber,
	}
}
