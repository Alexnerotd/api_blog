package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbData struct {
	Db_user     string
	Db_password string
	Db_host     string
	Db_name     string
}

func ConnectDB() (*gorm.DB, error) {

	err := godotenv.Load("C:/Users/jesal/programing/projects/apis/blog/.env")
	if err != nil {
		log.Fatalf("Error loading file .env: %v", err)
	}

	dbData := &DbData{Db_user: os.Getenv("DB_USER"), Db_host: os.Getenv("DB_HOST"),
		Db_password: os.Getenv("DB_PASSWORD"), Db_name: os.Getenv("DB_NAME")}
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?parseTime=True`, dbData.Db_user, dbData.Db_password, dbData.Db_host, dbData.Db_name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
