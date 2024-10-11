package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/themelancholyspirit/airline-reservation-system/types"
	"github.com/themelancholyspirit/airline-reservation-system/util"
)

func (s *Server) handleCreateBooking(ctx *gin.Context) {
	var bookingRequest types.Booking

	// Bind JSON to bookingRequest
	if err := ctx.ShouldBindJSON(&bookingRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the request for debugging
	fmt.Printf("Booking Request after binding: %+v\n", bookingRequest)

	// Set FlightID and BookingTime
	bookingRequest.BookingTime = time.Now() // Set booking time to now if not set

	// Validate FlightID
	if bookingRequest.FlightID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid flight ID"})
		return
	}

	// Get the token from the context
	tokenString, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found in context"})
		return
	}

	// Validate the token and get the claims
	claims, err := util.ValidateToken(tokenString.(string))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Populate the Booking struct
	booking := types.Booking{
		UserID:        claims.UserID,
		FlightID:      bookingRequest.FlightID,
		Status:        bookingRequest.Status,
		SeatNumber:    bookingRequest.SeatNumber,
		BookingTime:   bookingRequest.BookingTime, // Ensure booking time is populated
		DepartureCity: bookingRequest.DepartureCity,
		ArrivalCity:   bookingRequest.ArrivalCity,
		DepartureTime: bookingRequest.DepartureTime,
		ArrivalTime:   bookingRequest.ArrivalTime,
	}

	// Log the booking before insertion
	fmt.Println("Booking before DB insertion:", booking)

	// Create booking in the database
	if err := s.Storage.CreateBooking(ctx, booking); err != nil {
		fmt.Println("Failed to create booking:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	ctx.JSON(http.StatusCreated, booking)
}

func (s *Server) handleGetBooking(ctx *gin.Context) {
	bookingID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	// Get the token from the context
	tokenString, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
		return
	}

	// Validate the token and get the claims
	claims, err := util.ValidateToken(tokenString.(string))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	booking, err := s.Storage.GetBooking(ctx, uint(bookingID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Check if the authenticated user owns the booking
	if booking.UserID != claims.UserID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this booking"})
		return
	}

	ctx.JSON(http.StatusOK, booking)
}

func (s *Server) handleUpdateBooking(ctx *gin.Context) {
	bookingID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	var updateRequest types.Booking

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the token from the context
	tokenString, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
		return
	}

	// Validate the token and get the claims
	claims, err := util.ValidateToken(tokenString.(string))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	booking, err := s.Storage.GetBooking(ctx, uint(bookingID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Check if the authenticated user owns the booking
	if booking.UserID != claims.UserID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this booking"})
		return
	}

	// Only allow updating the status
	booking.Status = updateRequest.Status

	if err := s.Storage.UpdateBooking(ctx, uint(bookingID), booking); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
		return
	}

	ctx.JSON(http.StatusOK, booking)
}

func (s *Server) handleGetUserBookings(ctx *gin.Context) {
	// Get the token from the context
	tokenString, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
		return
	}

	// Validate the token and get the claims
	claims, err := util.ValidateToken(tokenString.(string))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Get all bookings for the authenticated user
	bookings, err := s.Storage.GetBookingsByUserID(ctx, claims.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	// If no bookings found, return an empty array instead of null
	if bookings == nil {
		bookings = []types.Booking{}
	}

	ctx.JSON(http.StatusOK, bookings)
}

func (s *Server) handleCancelBooking(ctx *gin.Context) {
	bookingID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	// Get the token from the context
	tokenString, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
		return
	}

	// Validate the token and get the claims
	claims, err := util.ValidateToken(tokenString.(string))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	booking, err := s.Storage.GetBooking(ctx, uint(bookingID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Check if the authenticated user owns the booking
	if booking.UserID != claims.UserID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to cancel this booking"})
		return
	}

	booking.Status = "cancelled"

	if err := s.Storage.UpdateBooking(ctx, uint(bookingID), booking); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel booking"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Booking cancelled successfully"})
}
