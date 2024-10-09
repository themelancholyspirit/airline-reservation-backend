package types

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	ID        string `json:"id" gorm:"primaryKey"`
	BookingID string `json:"booking_id"`
	Seat      string `json:"seat"`
	Status    string `json:"status"`
}
