package types

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	ID        uint    `json:"id" gorm:"primaryKey"`
	BookingID uint    `json:"booking_id"`
	Amount    float64 `json:"amount"`
	Method    string  `json:"method"`
	Status    string  `json:"status"`
}
