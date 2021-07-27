package models

import (
	"errors"
	Config "github.com/zahlekhan/retailer/server/config"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Order struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	CustomerID uint           `json:"customer_id"`
	Customer   Customer       `json:"-"`
	Products   []Product      `gorm:"many2many:order_book;" json:"products"`
	Status     string         `gorm:"type:enum('order placed', 'pending', 'failed');default:'pending'" json:"status"`
}

func CreateOrder(o *Order) error {
	db, err := Config.ConnectDB()
	if err != nil {
		return err
	}
	err = db.Create(&o).Error
	if err != nil {
		return err
	}
	return nil
}

func FindOrderByID(o *Order, oid string) error {
	db, err := Config.ConnectDB()
	if err != nil {
		return err
	}
	err = db.Model(Order{}).Where("id = ?", oid).Preload(clause.Associations).First(&o).Error
	if err != nil {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("order not found")
	}
	return err
}

func UpdateOrder(oid uint, status string) error {
	db, err := Config.ConnectDB()
	if err != nil {
		return err
	}
	db = db.Model(Order{}).Where("id= ?", oid).Updates(map[string]interface{}{
		"status": status,
	})

	err = db.Error
	if err != nil {
		return err
	}

	return nil
}
