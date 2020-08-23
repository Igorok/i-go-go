package ordercontroller

import (
	"delivery-go/entities_layer/order/orderentity"
	"delivery-go/domain_layer/orderpkg"
	"delivery-go/domain_layer/userpkg"
	"delivery-go/application_layer/webgraphql/webtypes/ordertype"
	"time"

	"github.com/graphql-go/graphql"
)

// OrderController - web controller
type OrderController struct {
	ordercase orderpkg.UseCase
	usercase  userpkg.UseCase
}

// NewOrderController - contructor for OrderController
func NewOrderController(ordercase orderpkg.UseCase, usercase userpkg.UseCase) *OrderController {
	return &OrderController{
		ordercase: ordercase,
		usercase:  usercase,
	}
}

// CreateOrder - create order
func (controller OrderController) CreateOrder() *graphql.Field {
	return &graphql.Field{
		Type: ordertype.OrderType,
		Args: graphql.FieldConfigArgument{
			"token": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"dateCustomer": &graphql.ArgumentConfig{
				Type: graphql.DateTime,
			},
			"address": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"comment": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"cart": &graphql.ArgumentConfig{
				Type: graphql.NewList(ordertype.CartItemArgType),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			token, _ := p.Args["token"].(string)
			dateCustomer, _ := p.Args["dateCustomer"].(time.Time)
			comment, _ := p.Args["comment"].(string)
			address, _ := p.Args["address"].(string)

			user, _ := controller.usercase.ParseToken(p.Context, token)

			var cart []orderentity.CartItem

			products := p.Args["cart"].([]interface{})
			if products != nil {
				for _, c := range products {
					item := c.(map[string]interface{})
					cart = append(cart, orderentity.CartItem{
						IDPr:  item["idPr"].(string),
						Count: item["count"].(int),
						Price: item["price"].(int),
					})
				}
			}

			order, err := controller.ordercase.CreateOrder(p.Context, user, cart, dateCustomer, address, comment)

			return order, err
		},
	}

}

// GetOrders - get list of orders
func (controller OrderController) GetOrders() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(ordertype.OrderType),
		Args: graphql.FieldConfigArgument{
			"token": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"skip": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			token, _ := p.Args["token"].(string)
			skip, _ := p.Args["skip"].(int)
			limit, _ := p.Args["limit"].(int)

			user, _ := controller.usercase.ParseToken(p.Context, token)

			orders, err := controller.ordercase.GetOrders(p.Context, user, skip, limit)

			return orders, err
		},
	}

}
