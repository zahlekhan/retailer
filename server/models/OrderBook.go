package models

import (
	Config "github.com/zahlekhan/retailer/server/config"
	"gorm.io/gorm"
)

type OrderBook struct {
	OrderID   uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Quantity  uint
}

func (*OrderBook) TableName() string {
	return "order_book"
}

func BatchCreateOrderBookByOrderID(o *Order, ProductIDs, Quantities []uint) error {
	db, err := Config.ConnectDB()
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		//Create Order
		tx.Create(&o)

		// Batch create entries in OrderBook for given OrderID
		length := len(ProductIDs)
		for idx := 0; idx < length; idx++ {
			ob := OrderBook{
				OrderID:   o.ID,
				ProductID: ProductIDs[idx],
				Quantity:  Quantities[idx],
			}
			if err := tx.Create(&ob).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
