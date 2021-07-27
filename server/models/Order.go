package models

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Order struct {
	gorm.Model
	CustomerID uint      `json:"-"`
	Customer   Customer  `json:"-"`
	Products   []Product `gorm:"many2many:order_products;"`
	Status     string    `gorm:"type:enum('order placed', 'pending', 'failed');default:'pending'"`
}

//func (o *Order) CreateOrder(db *gorm.DB) (*Product, error) {
//
//}

func (o *Order) FindOrderByID(db *gorm.DB, oid uint) (*Order, error) {
	var err error
	err = db.Model(Order{}).Where("id = ?", oid).Preload(clause.Associations).First(&o).Error
	if err != nil {
		return &Order{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Order{}, errors.New("order not found")
	}
	return o, err
}

func (o *Order) UpdateOrder(db *gorm.DB, oid uint, status string) (*Order, error) {
	db = db.Model(Order{}).Where("id= ?", oid).Updates(map[string]interface{}{
		"status": status,
	})

	if db.Error != nil {
		return &Order{}, db.Error
	}
	return o, nil
}
