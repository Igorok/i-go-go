package testdata

import "i-go-go/entities_layer/product/productentity"

// GetTestProducts return data of test products
func GetTestProducts() []productentity.Product {
	TestProducts := []productentity.Product{
		productentity.Product{
			ID:       "5e874f4b327272d07e537a4d",
			IDCl:     "5e874f4b327272d07e537a44",
			IDEst:    "5e874f4b327272d07e537a46",
			Name:     "Steak New York",
			Category: "steak",
			Price:    200,
			Status:   "active",
		},
		productentity.Product{
			ID:       "5e874f4b327272d07e537a4c",
			IDCl:     "5e874f4b327272d07e537a44",
			IDEst:    "5e874f4b327272d07e537a46",
			Name:     "Red wine",
			Category: "drinks",
			Price:    150,
			Status:   "active",
		},
		productentity.Product{
			ID:       "5e874f4b327272d07e537a4b",
			IDCl:     "5e874f4b327272d07e537a44",
			IDEst:    "5e874f4b327272d07e537a46",
			Name:     "Salad Cesar",
			Category: "salad",
			Price:    100,
			Status:   "active",
		},
		productentity.Product{
			ID:       "5e874f4b327272d07e537a4a",
			IDCl:     "5e874f4b327272d07e537a44",
			IDEst:    "5e874f4b327272d07e537a45",
			Name:     "Cola",
			Category: "drinks",
			Price:    50,
			Status:   "active",
		},
		productentity.Product{
			ID:       "5e874f4b327272d07e537a49",
			IDCl:     "5e874f4b327272d07e537a44",
			IDEst:    "5e874f4b327272d07e537a45",
			Name:     "Chicken Barbecue",
			Category: "pizza",
			Price:    200,
			Status:   "active",
		},
		productentity.Product{
			ID:       "5e874f4b327272d07e537a48",
			IDCl:     "5e874f4b327272d07e537a44",
			IDEst:    "5e874f4b327272d07e537a45",
			Name:     "Four Cheese",
			Category: "pizza",
			Price:    150,
			Status:   "active",
		},
		productentity.Product{
			ID:       "5e874f4b327272d07e537a47",
			IDCl:     "5e874f4b327272d07e537a44",
			IDEst:    "5e874f4b327272d07e537a45",
			Name:     "Pepperoni",
			Category: "pizza",
			Price:    100,
			Status:   "active",
		},
	}

	return TestProducts
}
