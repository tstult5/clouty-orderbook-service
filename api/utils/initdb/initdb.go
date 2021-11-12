package initdb

import (
	"log"
  "time"
	"github.com/jinzhu/gorm"
	"github.com/tstult5/clouty-orderbook-service/api/models"
)

var orders = []models.Order{
	models.Order{
		OrderName: "CloutyOrder 1",
    OrderPrice: 2.01,
    OrderType: "buy",
		CreatedAt:   time.Now(),
	},
	models.Order{
    OrderName: "CloutyOrder 1",
    OrderPrice: 2.05,
    OrderType: "sell",
		CreatedAt:   time.Now(),
	},
}


func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Order{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Order{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range orders {
		err = db.Debug().Model(&models.Order{}).Create(&orders[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
