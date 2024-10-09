package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/themelancholyspirit/airline-reservation-system/types"
	"github.com/themelancholyspirit/airline-reservation-system/util"
)

func (s *Server) handleCreateBooking(ctx *gin.Context) {
	var bookingRequest types.Booking
	if err := ctx.ShouldBindJSON(&bookingRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the token from the context (set by middleware)
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

	// Parse the user ID from the claims
	userID, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	booking := types.Booking{
		UserID:   uint(userID),
		FlightID: bookingRequest.FlightID,
		Status:   "pending",
	}

	if err := s.Storage.CreateBooking(ctx, booking); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	ctx.JSON(http.StatusCreated, booking)
}

func (s *Server) handleGetBooking(ctx *gin.Context) {
	bookingID := ctx.Param("id")

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

	// Parse the user ID from the claims
	userID, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	booking, err := s.Storage.GetBooking(ctx, bookingID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Check if the authenticated user owns the booking
	if booking.UserID != uint(userID) {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this booking"})
		return
	}

	ctx.JSON(http.StatusOK, booking)
}

func (s *Server) handleUpdateBooking(ctx *gin.Context) {
	bookingID := ctx.Param("id")
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

	// Parse the user ID from the claims
	userID, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	booking, err := s.Storage.GetBooking(ctx, bookingID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Check if the authenticated user owns the booking
	if booking.UserID != uint(userID) {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this booking"})
		return
	}

	// Only allow updating the status
	booking.Status = updateRequest.Status

	if err := s.Storage.UpdateBooking(ctx, bookingID, booking); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
		return
	}

	ctx.JSON(http.StatusOK, booking)
}

func (s *Server) handleCancelBooking(ctx *gin.Context) {
	bookingID := ctx.Param("id")

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

	// Parse the user ID from the claims
	userID, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	booking, err := s.Storage.GetBooking(ctx, bookingID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Check if the authenticated user owns the booking
	if booking.UserID != uint(userID) {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to cancel this booking"})
		return
	}

	booking.Status = "cancelled"

	if err := s.Storage.UpdateBooking(ctx, bookingID, booking); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel booking"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Booking cancelled successfully"})
}
