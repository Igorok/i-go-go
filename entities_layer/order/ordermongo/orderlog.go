package ordermongo

import (
	"delivery-go/entities_layer/order/orderentity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderLog entity in mongo
type OrderLog struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	IDOrd  primitive.ObjectID `bson:"_id_order"`
	IDUsr  primitive.ObjectID `bson:"_id_user,omitempty"`
	Date   time.Time
	Status string
}

// ToEntityOrderLog - convert mongod object to entity
func ToEntityOrderLog(ol *OrderLog) *orderentity.OrderLog {
	ID := ""
	if !ol.ID.IsZero() {
		ID = ol.ID.Hex()
	}

	IDOrd := ol.IDOrd.Hex()
	IDUsr := ol.IDUsr.Hex()

	return &orderentity.OrderLog{
		ID:     ID,
		IDOrd:  IDOrd,
		IDUsr:  IDUsr,
		Date:   ol.Date,
		Status: ol.Status,
	}
}

// ToOrderLog - convert to mongo object
func ToOrderLog(ol *orderentity.OrderLog) *OrderLog {
	var ID primitive.ObjectID
	if ol.ID != "" {
		ID, _ = primitive.ObjectIDFromHex(ol.ID)
	}

	IDOrd, _ := primitive.ObjectIDFromHex(ol.IDOrd)
	IDUsr, _ := primitive.ObjectIDFromHex(ol.IDUsr)

	return &OrderLog{
		ID:     ID,
		IDOrd:  IDOrd,
		IDUsr:  IDUsr,
		Date:   ol.Date,
		Status: ol.Status,
	}
}
