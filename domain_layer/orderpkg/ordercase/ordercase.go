package ordercase

import (
	"context"
	"delivery-go/entities_layer/order/orderentity"
	"delivery-go/entities_layer/user/userentity"
	"delivery-go/domain_layer/orderpkg"
	"time"

	"gopkg.in/go-playground/validator.v8"
)

var validate *validator.Validate

type OrderCase struct {
	orderRepo orderpkg.Repository
}

func NewOrderCase(orderRepo orderpkg.Repository) *OrderCase {
	return &OrderCase{
		orderRepo: orderRepo,
	}
}

func (o *OrderCase) GetOrder(ctx context.Context, user *userentity.UserSystem, id string) (*orderentity.Order, error) {
	return &orderentity.Order{}, nil
}

// GetOrders - operator or admin of establishment should see list of orders
func (o *OrderCase) GetOrders(
	ctx context.Context,
	user *userentity.UserSystem,
	skip, limit int,
) ([]*orderentity.Order, error) {
	if user == nil {
		return nil, orderpkg.ErrOrderNotFound
	}

	return o.orderRepo.GetOrders(ctx, user, skip, limit)
}

// CreateOrder - user of system or unknown customer could make order
func (o *OrderCase) CreateOrder(
	ctx context.Context,
	user *userentity.UserSystem,
	cart []orderentity.CartItem,
	dateCustomer time.Time,
	address string,
	comment string,
) (*orderentity.Order, error) {
	if cart == nil || len(cart) == 0 {
		return nil, orderpkg.ErrOrderWrongData
	}

	// get ids
	productIds := make([]string, len(cart))
	cartByProduct := make(map[string]orderentity.CartItem)

	for i, pr := range cart {
		productIds[i] = pr.IDPr
		cartByProduct[pr.IDPr] = pr
	}

	// find products
	products, err := o.orderRepo.GetProducts(ctx, productIds)
	if err != nil {
		return nil, err
	}

	// calculate price
	price := 0
	finalCart := make([]orderentity.CartItem, len(cart))

	for i, pr := range products {
		item := cartByProduct[pr.ID]
		item.Price = pr.Price

		price += item.Price * item.Count

		finalCart[i] = item
	}

	// insert order
	IDCl := (*products[0]).IDCl
	IDEst := (*products[0]).IDEst

	order := orderentity.Order{
		IDCl:         IDCl,
		IDEst:        IDEst,
		Cart:         finalCart,
		Price:        price,
		DateCreate:   time.Now(),
		DateCustomer: dateCustomer,
		Address:      address,
		Comment:      comment,
		Status:       orderpkg.StatusCreate,
	}

	if user != nil {
		order.Status = orderpkg.StatusApprove
	}

	// validation
	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)
	errs := validate.Struct(order)

	if errs != nil {
		return nil, orderpkg.ErrOrderWrongData
	}

	// insert order
	_, err = o.orderRepo.CreateOrder(ctx, &order)
	if err != nil {
		return nil, err
	}

	IDUsr := ""
	if user != nil {
		IDUsr = user.ID
	}

	// insert logs
	orderLogs := make([]*orderentity.OrderLog, 0)

	orderLogs = append(orderLogs, &orderentity.OrderLog{
		IDOrd:  order.ID,
		IDUsr:  IDUsr,
		Date:   time.Now(),
		Status: orderpkg.StatusCreate,
	})

	if user != nil {
		orderLogs = append(orderLogs, &orderentity.OrderLog{
			IDOrd:  order.ID,
			IDUsr:  IDUsr,
			Date:   time.Now(),
			Status: orderpkg.StatusApprove,
		})
	}

	_, err = o.orderRepo.InsertLogs(ctx, orderLogs)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *OrderCase) UpdateStatus(ctx context.Context, user *userentity.UserSystem, status string) (*orderentity.Order, error) {
	return &orderentity.Order{}, nil
}
