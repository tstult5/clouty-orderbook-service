package models

import (
	"errors"
	"html"
	"strings"
	"time"
	"github.com/jinzhu/gorm"
)

type Order struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	OrderName  string    `gorm:"size:255;not null" json:"ordername"`
  OrderPrice float32 `gorm:"default:0" json:"order_price"`
  OrderType string    `gorm:"default:nil" json:"order_type"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type OrderCreate struct {
  OrderName  string    `gorm:"size:255;not null;unique" json:"ordername"`
  OrderPrice float32 `gorm:"default:0" json:"order_price"`
  OrderType string    `gorm:"default:nil" json:"order_type"`
}

type OrderBook struct {
  OrderBookName string `gorm:"size:255;not null;unique" json:"orderbook_name"`
  buy_price float32
  buy_count uint32
  sell_count uint32
  sell_price float32
  //price, count, and amount.
}

func (o *Order) Prepare() {
	o.ID = 0
	o.OrderName = html.EscapeString(strings.TrimSpace(o.OrderName))
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
}

func (o *Order) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if o.OrderName == "" {
			return errors.New("Required OrderName")
		}
		return nil

	default:
		if o.OrderName == "" {
			return errors.New("Required OrderName")
		}
			return nil
	}
}
// SaveOrder godoc
// @Summary Create new Order
// @Description Create new order sell / buy type
// @Tags orders
// @Param order body models.OrderCreate true "Order Data"
// @Accept  json
// @Produce  json
// @Success 200 {object} Order
// @Router /order/create [post]
func (o *Order) SaveOrder(db *gorm.DB) (*Order, error) {

	var err error
	err = db.Debug().Create(&o).Error
	if err != nil {
		return &Order{}, err
	}
	return o, nil
}

// FindAllOrders godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {object} Order
// @Router /orders [get]
func (u *Order) FindAllOrders(db *gorm.DB) (*[]Order, error) {
	var err error
	orders := []Order{}
	err = db.Debug().Model(&Order{}).Limit(100).Find(&orders).Error
	if err != nil {
		return &[]Order{}, err
	}
	return &orders, err
}


// FindAllOrders godoc
// @Summary Get aggregation of orders by orderName
// @Description Get aggregation of orders by orderName
// @Tags orderBooks
// @Accept  json
// @Produce  json
// @Param ordername path string true "Order Name"
// @Success 200 {object} Order
// @Router /orderbooks/{ordername} [get]
func (u *Order) FindOrderBookByName(db *gorm.DB, ordername string) (*OrderBook, error) {
	var err error
  orderbook := OrderBook{}
  orders := []Order{}
	err = db.Debug().Model(&Order{}).Where("OrderName = ?", ordername).Find(&orders).Error
  //TO DO : Questions on the orderbook model - what structure should the orderbok have ?
	if err != nil {
    for _, order := range orders {
       orderbook.OrderBookName = order.OrderName
       orderbook.buy_count = 5
       orderbook.sell_count =5
       orderbook.sell_price = 2.55
       orderbook.buy_price = 2.53
   }
		return &orderbook, err
	}
	return &orderbook, err
}


// FindOrderByID godoc
// @Summary Get details of order by id
// @Description Get details of order by id
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {object} Order
// @Router /order/{id} [get]
func (o *Order) FindOrderByID(db *gorm.DB, uid uint32) (*Order, error) {
	var err error
	err = db.Debug().Model(Order{}).Where("id = ?", uid).Take(&o).Error
	if err != nil {
		return &Order{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Order{}, errors.New("Order Not Found")
	}
	return o, err
}
