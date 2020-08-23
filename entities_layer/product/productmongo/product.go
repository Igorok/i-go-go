package productmongo

import (
	"i-go-go/entities_layer/product/productentity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product is object of product for organisation
type Product struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	IDCl     primitive.ObjectID `bson:"_id_client"`
	IDEst    primitive.ObjectID `bson:"_ids_estb"`
	Name     string
	Category string
	Price    int
	Status   string
}

func ToEntity(p *Product) *productentity.Product {
	return &productentity.Product{
		ID:       p.ID.Hex(),
		IDCl:     p.IDCl.Hex(),
		IDEst:    p.IDEst.Hex(),
		Name:     p.Name,
		Category: p.Category,
		Price:    p.Price,
		Status:   p.Status,
	}
}

func toProduct(pm *productentity.Product) *Product {
	id, _ := primitive.ObjectIDFromHex(pm.ID)
	idCl, _ := primitive.ObjectIDFromHex(pm.IDCl)
	idEst, _ := primitive.ObjectIDFromHex(pm.IDEst)

	return &Product{
		ID:       id,
		IDCl:     idCl,
		IDEst:    idEst,
		Name:     pm.Name,
		Category: pm.Category,
		Price:    pm.Price,
		Status:   pm.Status,
	}
}
