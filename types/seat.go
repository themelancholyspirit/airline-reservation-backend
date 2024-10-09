package types

import (
	"gorm.io/gorm"
)

type Seat struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	SeatNumber string `json:"seatNumber"`
	FlightID   uint   `json:"flightId" gorm:"foreignKey:FlightID"`
	IsBooked   bool   `json:"isBooked"`
}
