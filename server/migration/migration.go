package main

import (
	"fmt"
	"log"
	"os"
	"simple-dashboard-server/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	migrateTables = []interface{}{
		&model.User{},
	}
)

// GetMigrateTables get migrate table list
func GetMigrateTables() []interface{} {
	return migrateTables
}

var _ = godotenv.Load()

func main() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port)

	if len(migrateTables) > 0 {
		gormWrite, err := gorm.Open(postgres.Open(dns), &gorm.Config{SkipDefaultTransaction: true})
		if err != nil {
			log.Fatal(err)
		}
		tx := gormWrite.Begin()
		if err := tx.AutoMigrate(migrateTables...); err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		tx.Commit()
	}
}
