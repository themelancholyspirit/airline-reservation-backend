package types

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	ID        string  `json:"id" gorm:"primaryKey"`
	BookingID string  `json:"booking_id"`
	Amount    float64 `json:"amount"`
	Method    string  `json:"method"`
	Status    string  `json:"status"`
}
