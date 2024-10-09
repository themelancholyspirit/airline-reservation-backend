package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/themelancholyspirit/airline-reservation-system/types"
)

// Config holds the database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgreDB creates and returns a new PostgreSQL database connection
func NewPostgreDB(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(&types.User{}, &types.Flight{}, &types.Booking{}, &types.Reservation{}, &types.Payment{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto-migrate schema: %w", err)
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
