package main

import (
	"delivery-go/bin/testdata"
	"delivery-go/utils"
	"log"
)

func main() {
	if err := utils.CfgInit(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	db := utils.MongoInit("delivery")

	testdata.InsertClient(db, "clients")
	testdata.InsertEstablishments(db, "establishments")
	testdata.InsertProducts(db, "products")
	testdata.InsertUsers(db, "users_system")
}
