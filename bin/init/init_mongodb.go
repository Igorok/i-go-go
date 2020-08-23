package main

import (
	"i-go-go/bin/testdata"
	"i-go-go/service_layer"
	"log"
)

func main() {
	if err := service_layer.CfgInit(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	db := service_layer.MongoInit("delivery")

	testdata.InsertClient(db, "clients")
	testdata.InsertEstablishments(db, "establishments")
	testdata.InsertProducts(db, "products")
	testdata.InsertUsers(db, "users_system")
}
