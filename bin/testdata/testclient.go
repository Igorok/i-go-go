package testdata

import "delivery-go/entities/client/cliententity"

// GetTestClient return example of client data
func GetTestClient() cliententity.Client {
	TestClient := cliententity.Client{
		ID:          "5e874f4b327272d07e537a44",
		Name:        "test",
		Email:       "test@test.tst",
		PhoneNumber: "12345678",
	}

	return TestClient
}

// GetTestEstablishments return example of client data
func GetTestEstablishments() []cliententity.Establishment {
	TestEstablishments := []cliententity.Establishment{
		cliententity.Establishment{
			ID:          "5e874f4b327272d07e537a46",
			IDCl:        "5e874f4b327272d07e537a44",
			Name:        "pizza",
			Email:       "pizza@test.tst",
			PhoneNumber: "12345678",
		},
		cliententity.Establishment{
			ID:          "5e874f4b327272d07e537a45",
			IDCl:        "5e874f4b327272d07e537a44",
			Name:        "restaurant",
			Email:       "restaurant@test.tst",
			PhoneNumber: "87654321",
		},
	}

	return TestEstablishments
}
