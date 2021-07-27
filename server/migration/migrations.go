package migration

import (
	Config "github.com/zahlekhan/retailer/server/config"
	"github.com/zahlekhan/retailer/server/models"
)

func Migrations() error {

	DB, err := Config.ConnectDB()
	if err != nil {
		return err
	}

	err = DB.Migrator().DropTable(&models.Customer{}, &models.Product{}, &models.Order{}, &models.OrderBook{}, "order_products")
	if err != nil {
		return err
	}

	if err = DB.AutoMigrate(&models.Customer{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&models.Product{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&models.Order{}); err != nil {
		return err
	}

	err = DB.SetupJoinTable(&models.Order{}, "Products", &models.OrderBook{})
	if err != nil {
		return err
	}

	if err = DB.AutoMigrate(&models.OrderBook{}); err != nil {
		return err
	}

	return nil
}
