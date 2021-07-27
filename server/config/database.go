package Config

import (
	"fmt"
	"github.com/zahlekhan/retailer/server/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
)

// DBConfig represents db configuration

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func buildDBConfig() *DBConfig {
	port, _ := strconv.Atoi(os.Getenv("TestDbPort"))
	dbConfig := DBConfig{
		Host:     os.Getenv("TestDbHost"),
		Port:     port,
		User:     os.Getenv("TestDbUser"),
		Password: os.Getenv("TestDbPassword"),
		DBName:   os.Getenv("TestDbName"),
	}
	return &dbConfig
}

func dbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func Migrations() error {

	DB, err := ConnectDB()
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

// ConnectDB opens a connection to the database
func ConnectDB() (*gorm.DB, error) {

	DB, err := gorm.Open(mysql.Open(dbURL(buildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	return DB, err
}
