package types

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID        uint      `json:"user_id"`
	FlightID      uint      `json:"flight_id" gorm:"foreignKey:FlightID;references:ID"`
	Status        string    `json:"status"`
	SeatNumber    string    `json:"seat_number"`
	BookingTime   time.Time `json:"booking_time" gorm:"auto_now_add"`
	DepartureCity string    `json:"departure_city"`
	ArrivalCity   string    `json:"arrival_city"`
	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time"`
}
