package types

import (
	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	ID               uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	FlightNumber     string  `json:"flightNumber"`
	DepartureAirport string  `json:"departureAirport"`
	ArrivalAirport   string  `json:"arrivalAirport"`
	DepartureTime    string  `json:"departureTime"`
	ArrivalTime      string  `json:"arrivalTime"`
	Capacity         int     `json:"capacity"`
	AvailableSeats   int     `json:"availableSeats"`
	Price            float64 `json:"price"`
	Status           string  `json:"status"`
}

type FlightResponse struct {
	ID               uint    `json:"id"`
	FlightNumber     string  `json:"flightNumber"`
	DepartureAirport string  `json:"departureAirport"`
	ArrivalAirport   string  `json:"arrivalAirport"`
	DepartureTime    string  `json:"departureTime"`
	ArrivalTime      string  `json:"arrivalTime"`
	Capacity         int     `json:"capacity"`
	AvailableSeats   int     `json:"availableSeats"`
	Price            float64 `json:"price"`
	Status           string  `json:"status"`
}
