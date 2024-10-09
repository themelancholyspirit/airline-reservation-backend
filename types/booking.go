package types

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID            uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	FlightID      uint   `json:"flightId" gorm:"foreignKey:FlightID"`
	UserID        uint   `json:"userId" gorm:"foreignKey:UserID"`
	SeatNumber    string `json:"seatNumber"`
	BookingTime   string `json:"bookingTime"`
	Status        string `json:"status"`
	DepartureCity string `json:"departureCity"`
	ArrivalCity   string `json:"arrivalCity"`
	DepartureTime string `json:"departureTime"`
	ArrivalTime   string `json:"arrivalTime"`
}
