package main

import (
	"context"
	"i-go-go/entities_layer/client/clientmongo"
	"i-go-go/entities_layer/product/productmongo"
	"i-go-go/entities_layer/user/usermongo"
	"i-go-go/service_layer"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func insertClient(db *mongo.Database) *clientmongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := &clientmongo.Client{
		Name:        "test",
		Email:       "test@test.tst",
		PhoneNumber: "12345678",
	}
	res, err := db.Collection("clients").InsertOne(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	client.ID = res.InsertedID.(primitive.ObjectID)

	return client
}

func getClientByID(db *mongo.Database, ID string) *clientmongo.Client {
	client := new(clientmongo.Client)
	clientID, e1 := primitive.ObjectIDFromHex(ID)
	if e1 != nil {
		log.Fatal(e1)
	}

	q := bson.M{
		"_id": clientID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := db.Collection("clients").FindOne(ctx, q).Decode(client)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func insertEstablishments(db *mongo.Database, clientID primitive.ObjectID) *[]clientmongo.Establishment {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	establishments := []interface{}{
		clientmongo.Establishment{
			IDCl:        clientID,
			Name:        "pizza",
			Email:       "pizza@test.tst",
			PhoneNumber: "12345678",
		},
		clientmongo.Establishment{
			IDCl:        clientID,
			Name:        "restaurant",
			Email:       "restaurant@test.tst",
			PhoneNumber: "87654321",
		},
	}

	res, err := db.Collection("establishments").InsertMany(ctx, establishments)
	if err != nil {
		log.Fatal(err)
	}

	var estArr []clientmongo.Establishment
	for key, mongoID := range res.InsertedIDs {
		est := establishments[key].(clientmongo.Establishment)
		est.ID = mongoID.(primitive.ObjectID)
		estArr = append(estArr, est)
	}

	// return &IDs
	return &estArr
}

func insertUsers(db *mongo.Database, eArr *[]clientmongo.Establishment) []usermongo.UserSystem {
	hashSalt := viper.GetString("app.hash_salt")
	usrArr := []usermongo.UserSystem{}

	usrArr = append(usrArr, usermongo.UserSystem{
		IDCl:        (*eArr)[0].IDCl,
		IDsEst:      []primitive.ObjectID{(*eArr)[0].ID, (*eArr)[1].ID},
		Login:       "admin",
		Password:    service_layer.HashPwd("admin", hashSalt),
		Email:       "admin@admin.ru",
		PhoneNumber: "12345678",
		Name:        "Admin",
		Surname:     "Test",
		Patronymic:  "Default",
		Roles:       []string{"ADMIN"},
	})

	for ind, est := range *eArr {
		operatorLogin := service_layer.ReplaceWhitespace("operator "+est.Name, "_")
		operatorPassword := service_layer.HashPwd(operatorLogin, hashSalt)

		courierLogin := service_layer.ReplaceWhitespace("courier "+est.Name, "_")
		courierPassword := service_layer.HashPwd(courierLogin, hashSalt)

		operator := usermongo.UserSystem{
			IDCl:        est.IDCl,
			IDsEst:      []primitive.ObjectID{est.ID},
			Login:       operatorLogin,
			Password:    operatorPassword,
			Email:       est.Email,
			PhoneNumber: est.PhoneNumber,
			Name:        "Operator " + strconv.Itoa(ind),
			Surname:     "Test " + strconv.Itoa(ind),
			Patronymic:  "Default " + strconv.Itoa(ind),
			Roles:       []string{"OPERATOR"},
		}
		courier := usermongo.UserSystem{
			IDCl:        est.IDCl,
			IDsEst:      []primitive.ObjectID{est.ID},
			Login:       courierLogin,
			Password:    courierPassword,
			Email:       est.Email,
			PhoneNumber: est.PhoneNumber,
			Name:        "Courier " + strconv.Itoa(ind),
			Surname:     "Test " + strconv.Itoa(ind),
			Patronymic:  "Default " + strconv.Itoa(ind),
			Roles:       []string{"COURIER"},
		}

		usrArr = append(usrArr, operator)
		usrArr = append(usrArr, courier)
	}

	insArr := []interface{}{}
	for _, usr := range usrArr {
		insArr = append(insArr, usr)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := db.Collection("users_system").InsertMany(ctx, insArr)
	if err != nil {
		log.Fatal(err)
	}

	for ind, mongoID := range res.InsertedIDs {
		usrArr[ind].ID = mongoID.(primitive.ObjectID)
	}

	return usrArr
}

func insertProducts(db *mongo.Database, eArr *[]clientmongo.Establishment) []productmongo.Product {
	type ProdData struct {
		Name     string
		Category string
		Price    int
	}

	defaultPizza := []ProdData{
		ProdData{
			Name:     "Pepperoni",
			Category: "pizza",
			Price:    100,
		},
		ProdData{
			Name:     "Four Cheese",
			Category: "pizza",
			Price:    150,
		},
		ProdData{
			Name:     "Chicken Barbecue",
			Category: "pizza",
			Price:    200,
		},
		ProdData{
			Name:     "Cola",
			Category: "drinks",
			Price:    50,
		},
	}

	defaultRestaurant := []ProdData{
		ProdData{
			Name:     "Salad Cesar",
			Category: "salad",
			Price:    100,
		},
		ProdData{
			Name:     "Red wine",
			Category: "drinks",
			Price:    150,
		},
		ProdData{
			Name:     "Steak New York",
			Category: "steak",
			Price:    200,
		},
	}

	insArr := []interface{}{}
	for _, est := range *eArr {
		if est.Name == "pizza" {
			for _, pr := range defaultPizza {
				product := productmongo.Product{
					IDCl:     est.IDCl,
					IDEst:    est.ID,
					Name:     pr.Name,
					Category: pr.Category,
					Price:    pr.Price,
					Status:   "active",
				}
				insArr = append(insArr, product)
			}
		} else {
			for _, pr := range defaultRestaurant {
				product := productmongo.Product{
					IDCl:     est.IDCl,
					IDEst:    est.ID,
					Name:     pr.Name,
					Category: pr.Category,
					Price:    pr.Price,
					Status:   "active",
				}
				insArr = append(insArr, product)
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := db.Collection("products").InsertMany(ctx, insArr)
	if err != nil {
		log.Fatal(err)
	}

	var products []productmongo.Product
	for ind, mongoID := range res.InsertedIDs {
		pr := insArr[ind].(productmongo.Product)
		pr.ID = mongoID.(primitive.ObjectID)
		products = append(products, pr)
	}

	return products
}

func main_v0() {
	if err := service_layer.CfgInit(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	db := service_layer.MongoInit("")

	defaultClient := insertClient(db)
	fmt.Println("defaultClient", defaultClient)

	establishmentsArr := insertEstablishments(db, defaultClient.ID)
	fmt.Println("establishmentsArr", establishmentsArr)

	productsArr := insertProducts(db, establishmentsArr)
	fmt.Println("productsArr", productsArr)

	users := insertUsers(db, establishmentsArr)
	fmt.Println("users", users)
}
