package models

import (
	"errors"
	"html"
	"strings"
	"time"
	"github.com/jinzhu/gorm"
)

type OrderBook struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	OrderBookName  string    `gorm:"size:255;not null;unique" json:"orderbookname"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (o *OrderBook) Prepare() {
	o.ID = 0
	o.OrderBookName = html.EscapeString(strings.TrimSpace(o.OrderBookName))
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
}

func (o *OrderBook) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if o.OrderBookName == "" {
			return errors.New("Required OrderBookName")
		}
		return nil

	default:
		if o.OrderBookName == "" {
			return errors.New("Required OrderBookName")
		}
			return nil
	}
}

func (o *OrderBook) SaveOrderBook(db *gorm.DB) (*OrderBook, error) {

	var err error
	err = db.Debug().Create(&o).Error
	if err != nil {
		return &OrderBook{}, err
	}
	return o, nil
}

// GetOrderBook godoc
// @Summary Get details of all orderbooks
// @Description Get details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {object} Order
// @Router /orders [get]
func (u *OrderBook) FindAllOrderBooks(db *gorm.DB) (*[]OrderBook, error) {
	var err error
	orderbooks := []OrderBook{}
	err = db.Debug().Model(&OrderBook{}).Limit(100).Find(&orderbooks).Error
	if err != nil {
		return &[]OrderBook{}, err
	}
	return &orderbooks, err
}

// GetOrderBook godoc
// @Summary Get details of orderbook by id
// @Description Get details of orderbook by id
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {object} Order
// @Router /orders/{id} [get]
func (o *OrderBook) FindOrderBookByID(db *gorm.DB, uid uint32) (*OrderBook, error) {
	var err error
	err = db.Debug().Model(OrderBook{}).Where("id = ?", uid).Take(&o).Error
	if err != nil {
		return &OrderBook{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &OrderBook{}, errors.New("OrderBook Not Found")
	}
	return o, err
}
