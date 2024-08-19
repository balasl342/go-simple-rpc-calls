package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type Db struct {
	DB *gorm.DB
}

var Config Db

// Connect DB
func DBconnection() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))
	Config.DB, err = gorm.Open("postgres", connect)
	if err != nil {
		return err
	}
	return nil
}
