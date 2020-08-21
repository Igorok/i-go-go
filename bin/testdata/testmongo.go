package testdata

import (
	"context"
	"delivery-go/entities/client/clientmongo"
	"delivery-go/entities/product/productmongo"
	"delivery-go/entities/user/usermongo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertClient insert test client data
func InsertClient(db *mongo.Database, collection string) {
	if collection == "" {
		collection = "clients"
	}

	cm := GetTestClient()
	clID, e := primitive.ObjectIDFromHex(cm.ID)
	if e != nil {
		log.Fatal(e)
	}

	client := &clientmongo.Client{
		ID:          clID,
		Name:        cm.Name,
		Email:       cm.Email,
		PhoneNumber: cm.PhoneNumber,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.Collection(collection).InsertOne(ctx, client)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertEstablishments insert test data
func InsertEstablishments(db *mongo.Database, collection string) {
	if collection == "" {
		collection = "establishments"
	}

	emArray := GetTestEstablishments()
	// establishments := make([]clientmongo.Establishment, len(emArray))
	establishments := make([]interface{}, len(emArray))

	for i, em := range emArray {
		ID, _ := primitive.ObjectIDFromHex(em.ID)
		IDCl, _ := primitive.ObjectIDFromHex(em.IDCl)

		e := clientmongo.Establishment{
			ID:          ID,
			IDCl:        IDCl,
			Name:        em.Name,
			Email:       em.Email,
			PhoneNumber: em.PhoneNumber,
		}
		establishments[i] = e
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.Collection(collection).InsertMany(ctx, establishments)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertProducts insert test data
func InsertProducts(db *mongo.Database, collection string) {
	if collection == "" {
		collection = "products"
	}

	pmArray := GetTestProducts()
	products := make([]interface{}, len(pmArray))

	for i, pm := range pmArray {
		ID, _ := primitive.ObjectIDFromHex(pm.ID)
		IDCl, _ := primitive.ObjectIDFromHex(pm.IDCl)
		IDEst, _ := primitive.ObjectIDFromHex(pm.IDEst)

		product := productmongo.Product{
			ID:       ID,
			IDCl:     IDCl,
			IDEst:    IDEst,
			Name:     pm.Name,
			Category: pm.Category,
			Price:    pm.Price,
			Status:   pm.Status,
		}

		products[i] = product
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.Collection(collection).InsertMany(ctx, products)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertUsers insert test data
func InsertUsers(db *mongo.Database, collection string) {
	if collection == "" {
		collection = "users_system"
	}

	umArray := GetTestUsers()
	users := make([]interface{}, len(umArray))

	for i, um := range umArray {
		ID, _ := primitive.ObjectIDFromHex(um.ID)
		IDCl, _ := primitive.ObjectIDFromHex(um.IDCl)
		IDsEst := make([]primitive.ObjectID, len(um.IDsEst))

		for iID, valID := range um.IDsEst {
			IDsEst[iID], _ = primitive.ObjectIDFromHex(valID)
		}

		user := usermongo.UserSystem{
			ID:          ID,
			IDCl:        IDCl,
			IDsEst:      IDsEst,
			Login:       um.Login,
			Password:    um.Password,
			Email:       um.Email,
			PhoneNumber: um.PhoneNumber,
			Name:        um.Name,
			Surname:     um.Surname,
			Patronymic:  um.Patronymic,
			Roles:       um.Roles,
		}

		users[i] = user
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.Collection(collection).InsertMany(ctx, users)
	if err != nil {
		log.Fatal(err)
	}
}
