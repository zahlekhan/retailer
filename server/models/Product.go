package models

import (
	"errors"
	Config "github.com/zahlekhan/retailer/server/config"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string `gorm:"size:100;not null;unique" json:"name"`
	Price    uint   `gorm:"check:(price >=0);not null" json:"price"`
	Quantity uint   `gorm:"check:(price >=0);not null" json:"quantity"`
}

func CreateProduct(p *Product) error {
	db, err := Config.ConnectDB()
	if err != nil {
		return err
	}
	err = db.Create(&p).Error
	if err != nil {
		return err
	}
	return nil
}

func FindAllProducts(products *[]Product) error {
	db, err := Config.ConnectDB()
	if err != nil {
		return err
	}
	err = db.Model(&Product{}).Limit(100).Find(&products).Error
	if err != nil {
		return err
	}
	return err
}

func FindProductByID(p *Product, pid string) error {
	db, err := Config.ConnectDB()
	if err != nil {
		return err
	}
	err = db.Model(Product{}).Where("id = ?", pid).First(&p).Error
	if err != nil {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("product not found")
	}
	return nil
}

func UpdateProduct(pid, price, qty string) error {
	db, err := Config.ConnectDB()
	if err != nil {
		return err
	}
	db = db.Model(Product{}).Where("id= ?", pid).Updates(map[string]interface{}{
		"price":    price,
		"quantity": qty,
	})
	err = db.Error

	if err != nil {
		return err
	}

	return nil
}
