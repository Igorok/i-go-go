package testdata

import "delivery-go/entities_layer/user/userentity"

// GetTestUsers return data of test products
func GetTestUsers() []userentity.UserSystem {
	TestUsers := []userentity.UserSystem{
		userentity.UserSystem{
			ID:          "5e874f4b327272d07e537a52",
			IDCl:        "5e874f4b327272d07e537a44",
			IDsEst:      []string{"5e874f4b327272d07e537a46"},
			Login:       "courier_restaurant",
			Password:    "f87c0b481d65def48a15a8728f3e3ac8fd51d095",
			Email:       "restaurant@test.tst",
			PhoneNumber: "87654321",
			Name:        "Courier 1",
			Surname:     "Test 1",
			Patronymic:  "Default 1",
			Roles:       []string{"COURIER"},
		},

		userentity.UserSystem{
			ID:          "5e874f4b327272d07e537a51",
			IDCl:        "5e874f4b327272d07e537a44",
			IDsEst:      []string{"5e874f4b327272d07e537a46"},
			Login:       "operator_restaurant",
			Password:    "d73b886f5d386c67c41698329b1a0055acb4543b",
			Email:       "restaurant@test.tst",
			PhoneNumber: "87654321",
			Name:        "Operator 1",
			Surname:     "Test 1",
			Patronymic:  "Default 1",
			Roles:       []string{"OPERATOR"},
		},
		userentity.UserSystem{
			ID:          "5e874f4b327272d07e537a50",
			IDCl:        "5e874f4b327272d07e537a44",
			IDsEst:      []string{"5e874f4b327272d07e537a45"},
			Login:       "courier_pizza",
			Password:    "a4a165fc8fca19f70e88fad37d5476ba1d0b7415",
			Email:       "pizza@test.tst",
			PhoneNumber: "12345678",
			Name:        "Courier 0",
			Surname:     "Test 0",
			Patronymic:  "Default 0",
			Roles:       []string{"COURIER"},
		},

		userentity.UserSystem{
			ID:          "5e874f4b327272d07e537a4f",
			IDCl:        "5e874f4b327272d07e537a44",
			IDsEst:      []string{"5e874f4b327272d07e537a45"},
			Login:       "operator_pizza",
			Password:    "5b8c79681c19a8132409da8987452fb2ff95ce5a",
			Email:       "pizza@test.tst",
			PhoneNumber: "12345678",
			Name:        "Operator 0",
			Surname:     "Test 0",
			Patronymic:  "Default 0",
			Roles:       []string{"OPERATOR"},
		},

		userentity.UserSystem{
			ID:   "5e874f4b327272d07e537a4e",
			IDCl: "5e874f4b327272d07e537a44",
			IDsEst: []string{
				"5e874f4b327272d07e537a45",
				"5e874f4b327272d07e537a46",
			},
			Login:       "admin",
			Password:    "b67a983d2035d6a4a642fc160b000dd70b62586d",
			Email:       "admin@admin.ru",
			PhoneNumber: "12345678",
			Name:        "Admin",
			Surname:     "Test",
			Patronymic:  "Default",
			Roles:       []string{"ADMIN"},
		},
	}

	return TestUsers
}
