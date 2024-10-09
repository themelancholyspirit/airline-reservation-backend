package storage

import (
	"context"

	"gorm.io/gorm"

	"github.com/themelancholyspirit/airline-reservation-system/types"
)

type PostgreStorage struct {
	db *gorm.DB
}

func NewPostgreStorage(db *gorm.DB) *PostgreStorage {
	return &PostgreStorage{db: db}
}

// User Storage Implementation
func (s *PostgreStorage) CreateUser(ctx context.Context, user types.User) error {
	return s.db.WithContext(ctx).Create(&user).Error
}

func (s *PostgreStorage) GetUser(ctx context.Context, id string) (types.User, error) {
	var user types.User
	err := s.db.WithContext(ctx).First(&user, "id = ?", id).Error
	return user, err
}

func (s *PostgreStorage) GetUserByEmail(ctx context.Context, email string) (types.User, error) {
	var user types.User
	err := s.db.WithContext(ctx).First(&user, "email = ?", email).Error
	return user, err
}

func (s *PostgreStorage) UpdateUser(ctx context.Context, user types.UserUpdateRequest) error {
	return s.db.WithContext(ctx).Model(&types.User{}).Where("email = ?", user.Email).Updates(user).Error
}

func (s *PostgreStorage) DeleteUser(ctx context.Context, id string) error {
	return s.db.WithContext(ctx).Delete(&types.User{}, "id = ?", id).Error
}

func (s *PostgreStorage) ListUsers(ctx context.Context) ([]types.User, error) {
	var users []types.User
	err := s.db.WithContext(ctx).Find(&users).Error
	return users, err
}

// Flight Storage Implementation
func (s *PostgreStorage) CreateFlight(ctx context.Context, flight types.Flight) error {
	return s.db.WithContext(ctx).Create(&flight).Error
}

func (s *PostgreStorage) GetFlight(ctx context.Context, id string) (types.Flight, error) {
	var flight types.Flight
	err := s.db.WithContext(ctx).First(&flight, "id = ?", id).Error
	return flight, err
}

func (s *PostgreStorage) UpdateFlight(ctx context.Context, id string, flight types.Flight) error {
	return s.db.WithContext(ctx).Model(&types.Flight{}).Where("id = ?", id).Updates(flight).Error
}

func (s *PostgreStorage) DeleteFlight(ctx context.Context, id string) error {
	return s.db.WithContext(ctx).Delete(&types.Flight{}, "id = ?", id).Error
}

func (s *PostgreStorage) ListFlights(ctx context.Context) ([]types.Flight, error) {
	var flights []types.Flight
	err := s.db.WithContext(ctx).Find(&flights).Error
	return flights, err
}

// Booking Storage Implementation
func (s *PostgreStorage) CreateBooking(ctx context.Context, booking types.Booking) error {
	return s.db.WithContext(ctx).Create(&booking).Error
}

func (s *PostgreStorage) GetBooking(ctx context.Context, id string) (types.Booking, error) {
	var booking types.Booking
	err := s.db.WithContext(ctx).First(&booking, "id = ?", id).Error
	return booking, err
}

func (s *PostgreStorage) UpdateBooking(ctx context.Context, id string, booking types.Booking) error {
	return s.db.WithContext(ctx).Model(&types.Booking{}).Where("id = ?", id).Updates(booking).Error
}

func (s *PostgreStorage) DeleteBooking(ctx context.Context, id string) error {
	return s.db.WithContext(ctx).Delete(&types.Booking{}, "id = ?", id).Error
}

func (s *PostgreStorage) ListBookings(ctx context.Context) ([]types.Booking, error) {
	var bookings []types.Booking
	err := s.db.WithContext(ctx).Find(&bookings).Error
	return bookings, err
}

// Reservation Storage Implementation
func (s *PostgreStorage) CreateReservation(ctx context.Context, reservation types.Reservation) error {
	return s.db.WithContext(ctx).Create(&reservation).Error
}

func (s *PostgreStorage) GetReservation(ctx context.Context, id string) (types.Reservation, error) {
	var reservation types.Reservation
	err := s.db.WithContext(ctx).First(&reservation, "id = ?", id).Error
	return reservation, err
}

func (s *PostgreStorage) UpdateReservation(ctx context.Context, id string, reservation types.Reservation) error {
	return s.db.WithContext(ctx).Model(&types.Reservation{}).Where("id = ?", id).Updates(reservation).Error
}

func (s *PostgreStorage) DeleteReservation(ctx context.Context, id string) error {
	return s.db.WithContext(ctx).Delete(&types.Reservation{}, "id = ?", id).Error
}

func (s *PostgreStorage) ListReservations(ctx context.Context) ([]types.Reservation, error) {
	var reservations []types.Reservation
	err := s.db.WithContext(ctx).Find(&reservations).Error
	return reservations, err
}

// Payment Storage Implementation
func (s *PostgreStorage) CreatePayment(ctx context.Context, payment types.Payment) error {
	return s.db.WithContext(ctx).Create(&payment).Error
}

func (s *PostgreStorage) GetPayment(ctx context.Context, id string) (types.Payment, error) {
	var payment types.Payment
	err := s.db.WithContext(ctx).First(&payment, "id = ?", id).Error
	return payment, err
}

func (s *PostgreStorage) UpdatePayment(ctx context.Context, id string, payment types.Payment) error {
	return s.db.WithContext(ctx).Model(&types.Payment{}).Where("id = ?", id).Updates(payment).Error
}

func (s *PostgreStorage) DeletePayment(ctx context.Context, id string) error {
	return s.db.WithContext(ctx).Delete(&types.Payment{}, "id = ?", id).Error
}

func (s *PostgreStorage) ListPayments(ctx context.Context) ([]types.Payment, error) {
	var payments []types.Payment
	err := s.db.WithContext(ctx).Find(&payments).Error
	return payments, err
}
