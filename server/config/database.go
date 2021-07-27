package Config

import (
	"fmt"
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

// ConnectDB opens a connection to the database
func ConnectDB() (*gorm.DB, error) {

	DB, err := gorm.Open(mysql.Open(dbURL(buildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	return DB, err
}
