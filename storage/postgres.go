package storage

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct{
	Host		string
	Port		string
	Password	string
	User		string
	DBName		string
	SSLMode		string
}

func NewConfig() (*Config) {
	config := &Config{
		Host: os.Getenv("DATABASE_HOST"),
		Port: os.Getenv("DATABASE_PORT"),
		Password: os.Getenv("DATABASE_PW"),
		User: os.Getenv("DATABASE_USER"),
		DBName: os.Getenv("DATABASE_NAME"),
		SSLMode: os.Getenv("DATABASE_SSLMODE")}
	return config
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s password=%s user=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Password, config.User, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}

type User struct {
	name	string
}

func createUser() error {
	return nil
}

func createWallet() error {
	return nil
}