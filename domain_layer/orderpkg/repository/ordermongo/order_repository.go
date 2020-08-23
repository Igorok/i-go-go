package ordermongo

import (
	"context"
	"delivery-go/entities_layer/order/orderentity"
	"delivery-go/entities_layer/order/ordermongo"
	"delivery-go/entities_layer/product/productentity"
	"delivery-go/entities_layer/product/productmongo"
	"delivery-go/entities_layer/user/userentity"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// OrderRepository implement of orderpkg.Repository
type OrderRepository struct {
	collOrders    *mongo.Collection
	collOrderLogs *mongo.Collection
	collProducts  *mongo.Collection
}

// NewOrderRepository is constructor for OrderRepository
func NewOrderRepository(db *mongo.Database, ordCollName, ordCollLogName, ordCollProdName string) *OrderRepository {
	return &OrderRepository{
		collOrders:    db.Collection(ordCollName),
		collOrderLogs: db.Collection(ordCollLogName),
		collProducts:  db.Collection(ordCollProdName),
	}
}

func (r *OrderRepository) GetOrder(ctx context.Context, user *userentity.UserSystem, id string) (*orderentity.Order, error) {
	return &orderentity.Order{}, nil
}

// GetOrders - get list of orders
func (r *OrderRepository) GetOrders(
	ctx context.Context,
	user *userentity.UserSystem,
	skip, limit int,
) ([]*orderentity.Order, error) {
	IDsEst := make([]primitive.ObjectID, len(user.IDsEst))
	for i, id := range user.IDsEst {
		IDsEst[i], _ = primitive.ObjectIDFromHex(id)
	}

	fmt.Println("IDsEst", IDsEst)

	q := bson.M{
		"_id_estb": bson.M{
			"$in": IDsEst,
		},
	}

	opts := options.Find()
	opts.SetSort(bson.M{"datecreate": -1})

	cur, err := r.collOrders.Find(ctx, q, opts)
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}

	out := make([]*orderentity.Order, 0)

	for cur.Next(ctx) {
		o := new(ordermongo.Order)
		err := cur.Decode(o)
		if err != nil {
			return nil, err
		}

		ord := ordermongo.ToEntityOrder(o)
		out = append(out, ord)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	fmt.Println("out", out)

	return out, nil
}

// CreateOrder - insert new order
func (r *OrderRepository) CreateOrder(ctx context.Context, order *orderentity.Order) (*orderentity.Order, error) {
	orderItem := ordermongo.ToOrder(order)

	res, err := r.collOrders.InsertOne(ctx, orderItem)
	if err != nil {
		return nil, err
	}

	order.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return order, nil
}

func (r *OrderRepository) UpdateStatus(
	ctx context.Context,
	user *userentity.UserSystem,
	id, status string,
) (*orderentity.Order, error) {
	return &orderentity.Order{}, nil
}

// GetProducts - get info about products
func (r *OrderRepository) GetProducts(ctx context.Context, ids []string) ([]*productentity.Product, error) {
	IDs := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		ID, _ := primitive.ObjectIDFromHex(id)
		IDs[i] = ID
	}
	q := bson.M{
		"status": "active",
		"_id": bson.M{
			"$in": IDs,
		},
	}
	opts := options.Find()
	opts.SetSort(bson.M{"price": 1})

	cur, err := r.collProducts.Find(ctx, q, opts)
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

// InsertLogs - save event of order
func (r *OrderRepository) InsertLogs(ctx context.Context, orderLogs []*orderentity.OrderLog) ([]*orderentity.OrderLog, error) {

	oLogs := make([]interface{}, len(orderLogs))

	for i, oEntity := range orderLogs {
		oLogs[i] = ordermongo.ToOrderLog(oEntity)
	}

	res, err := r.collOrderLogs.InsertMany(ctx, oLogs)
	if err != nil {
		return nil, err
	}

	for i, mongoID := range res.InsertedIDs {
		orderLogs[i].ID = mongoID.(primitive.ObjectID).Hex()
	}

	return orderLogs, nil
}
