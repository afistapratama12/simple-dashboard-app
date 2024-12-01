package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	Env ENV
	DB  *gorm.DB
}

type ENV struct {
	DBHost        string
	DBUser        string
	DBPass        string
	DBName        string
	DBPort        string
	DBHostProd    string
	BaseClientURL string
	JWTSecret     string
	AppENV        string
	EmailUsername string
	FromEmail     string
	AppPass       string
}

var _ = godotenv.Load()

func InitConfig() Config {
	var env = ENV{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPass:     getEnv("DB_PASS", "postgres"),
		DBName:     getEnv("DB_NAME", "postgres"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBHostProd: getEnv("DB_HOST_PROD", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),

		BaseClientURL: getEnv("BASE_CLIENT_URL", "http://localhost:3000"),
		JWTSecret:     getEnv("JWT_SECRET", ""),
		AppENV:        getEnv("APP_ENV", "development"),
		FromEmail:     getEnv("FROM_EMAIL", ""),
		EmailUsername: getEnv("EMAIL_USERNAME", ""),
		AppPass:       getEnv("APP_PASSWORD", ""),
	}

	db := InitDatabase(env)

	return Config{
		Env: env,
		DB:  db,
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func InitDatabase(env ENV) *gorm.DB {
	if env.AppENV == "production" {
		db, err := sql.Open("pgx", env.DBHostProd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			panic(err)
		}

		defer db.Close()

		if err := db.Ping(); err != nil {
			fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
			panic(err)
		}

		gormDB, err := gorm.Open(postgres.New(postgres.Config{
			Conn: db,
		}))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable open sql using gorm: %v\n", err)
			panic(err)
		}

		return gormDB
	} else {
		dns := "host=" + env.DBHost + " user=" + env.DBUser + " password=" + env.DBPass + " dbname=" + env.DBName + " port=" + env.DBPort + " sslmode=disable TimeZone=Asia/Jakarta"
		db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			panic(err)
		}
		return db
	}

}
