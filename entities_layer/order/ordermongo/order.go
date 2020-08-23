package ordermongo

import (
	"delivery-go/entities_layer/order/orderentity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order - model in mongo
type Order struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	IDCl         primitive.ObjectID `bson:"_id_client"`
	IDEst        primitive.ObjectID `bson:"_id_estb"`
	Cart         []CartItem
	Price        int
	DateCreate   time.Time
	DateCustomer time.Time
	Comment      string
	Address      string
	Status       string
}

// CartItem - model in mongo
type CartItem struct {
	IDPr  primitive.ObjectID `bson:"_id_prod"`
	Count int
	Price int
}

func ToEntityOrder(o *Order) *orderentity.Order {
	cartItems := make([]orderentity.CartItem, len(o.Cart))
	for i, ci := range o.Cart {
		cartItems[i] = orderentity.CartItem{
			IDPr:  ci.IDPr.Hex(),
			Count: ci.Count,
			Price: ci.Price,
		}
	}

	ID := ""
	if !o.ID.IsZero() {
		ID = o.ID.Hex()
	}

	return &orderentity.Order{
		ID:           ID,
		IDCl:         o.IDCl.Hex(),
		IDEst:        o.IDEst.Hex(),
		Cart:         cartItems,
		Price:        o.Price,
		DateCreate:   o.DateCreate,
		DateCustomer: o.DateCustomer,
		// DateCreate:   o.DateCreate.Format("2006-01-02 15:04:05"),
		// DateCustomer: o.DateCustomer.Format("2006-01-02 15:04:05"),
		Comment: o.Comment,
		Status:  o.Status,
	}
}

func ToOrder(om *orderentity.Order) *Order {
	IDCl, _ := primitive.ObjectIDFromHex(om.IDCl)
	IDEst, _ := primitive.ObjectIDFromHex(om.IDEst)

	var ID primitive.ObjectID
	if om.ID != "" {
		ID, _ = primitive.ObjectIDFromHex(om.ID)
	}

	cartItems := make([]CartItem, len(om.Cart))
	for i, ci := range om.Cart {
		IDPr, _ := primitive.ObjectIDFromHex(ci.IDPr)
		cartItems[i] = CartItem{
			IDPr:  IDPr,
			Count: ci.Count,
			Price: ci.Price,
		}
	}

	// tLayout := "2006-01-02 15:04:05"
	// dateCreate, _ := time.Parse(tLayout, om.DateCreate)
	// dateCustomer, _ := time.Parse(tLayout, om.DateCustomer)

	return &Order{
		ID:           ID,
		IDCl:         IDCl,
		IDEst:        IDEst,
		Cart:         cartItems,
		Price:        om.Price,
		DateCreate:   om.DateCreate,
		DateCustomer: om.DateCustomer,
		Address:      om.Address,
		Comment:      om.Comment,
		Status:       om.Status,
	}
}
