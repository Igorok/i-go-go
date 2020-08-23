package productmongo

import (
	"context"
	"delivery-go/entities_layer/product/productentity"
	"delivery-go/entities_layer/product/productmongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductRepository implement of productpkg.Repository
type ProductRepository struct {
	collProducts *mongo.Collection
}

// NewProductRepository is constructor for ProductRepository
func NewProductRepository(db *mongo.Database, prCollName string) *ProductRepository {
	return &ProductRepository{
		collProducts: db.Collection(prCollName),
	}
}

// GetProducts contain requests to get list of products
func (r ProductRepository) GetProducts(ctx context.Context) ([]*productentity.Product, error) {
	q := bson.M{
		"status": "active",
	}

	cur, err := r.collProducts.Find(ctx, q)
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}

	out := make([]*productentity.Product, 0)

	for cur.Next(ctx) {
		p := new(productmongo.Product)
		err := cur.Decode(p)
		if err != nil {
			return nil, err
		}

		pm := productmongo.ToEntity(p)
		out = append(out, pm)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

// GetProduct - detail of product
func (r ProductRepository) GetProduct(ctx context.Context, id string) (productentity.Product, error) {
	pm := productentity.Product{}

	ID, eId := primitive.ObjectIDFromHex(id)
	if eId != nil {
		return pm, eId
	}

	q := bson.M{
		"status": "active",
		"_id":    ID,
	}

	product := productmongo.Product{}
	err := r.collProducts.FindOne(ctx, q).Decode(&product)
	if err != nil {
		return pm, err
	}

	pm = (*productmongo.ToEntity(&product))

	return pm, nil
}
