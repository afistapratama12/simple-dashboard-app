package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		DBHost: getEnv("DB_HOST", "localhost"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASS", "postgres"),
		DBName: getEnv("DB_NAME", "postgres"),
		DBPort: getEnv("DB_PORT", "5432"),

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

var DB *gorm.DB

func InitDatabase(env ENV) *gorm.DB {
	if DB != nil {
		return DB
	}

	dns := "host=" + env.DBHost + " user=" + env.DBUser + " password=" + env.DBPass + " dbname=" + env.DBName + " port=" + env.DBPort + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	return db
}
