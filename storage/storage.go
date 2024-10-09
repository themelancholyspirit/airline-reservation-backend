package storage

import (
	"context"

	"github.com/themelancholyspirit/airline-reservation-system/types"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user types.User) error
	GetUser(ctx context.Context, id string) (types.User, error)
	UpdateUser(ctx context.Context, user types.UserUpdateRequest) error
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context) ([]types.User, error)
	GetUserByEmail(ctx context.Context, email string) (types.User, error)
}

type FlightStorage interface {
	CreateFlight(ctx context.Context, flight types.Flight) error
	GetFlight(ctx context.Context, id string) (types.Flight, error)
	UpdateFlight(ctx context.Context, id string, flight types.Flight) error
	DeleteFlight(ctx context.Context, id string) error
	ListFlights(ctx context.Context) ([]types.Flight, error)
}

type BookingStorage interface {
	CreateBooking(ctx context.Context, booking types.Booking) error
	GetBooking(ctx context.Context, id string) (types.Booking, error)
	UpdateBooking(ctx context.Context, id string, booking types.Booking) error
	DeleteBooking(ctx context.Context, id string) error
	ListBookings(ctx context.Context) ([]types.Booking, error)
}

type ReservationStorage interface {
	CreateReservation(ctx context.Context, reservation types.Reservation) error
	GetReservation(ctx context.Context, id string) (types.Reservation, error)
	UpdateReservation(ctx context.Context, id string, reservation types.Reservation) error
	DeleteReservation(ctx context.Context, id string) error
	ListReservations(ctx context.Context) ([]types.Reservation, error)
}

type PaymentStorage interface {
	CreatePayment(ctx context.Context, payment types.Payment) error
	GetPayment(ctx context.Context, id string) (types.Payment, error)
	UpdatePayment(ctx context.Context, id string, payment types.Payment) error
	DeletePayment(ctx context.Context, id string) error
	ListPayments(ctx context.Context) ([]types.Payment, error)
}

type Storage interface {
	UserStorage
	FlightStorage
	BookingStorage
	ReservationStorage
	PaymentStorage
}
