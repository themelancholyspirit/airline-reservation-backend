package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/themelancholyspirit/airline-reservation-system/types"
)

func (s *Server) handleCreateFlight(ctx *gin.Context) {
	var flight types.Flight
	if err := ctx.ShouldBindJSON(&flight); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.Storage.CreateFlight(ctx, flight); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create flight"})
		return
	}

	ctx.JSON(http.StatusCreated, flight)
}

func (s *Server) handleGetFlight(ctx *gin.Context) {
	id := ctx.Param("id")
	flight, err := s.Storage.GetFlight(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}

	ctx.JSON(http.StatusOK, flight)
}

func (s *Server) handleUpdateFlight(ctx *gin.Context) {
	id := ctx.Param("id")
	var flight types.Flight
	if err := ctx.ShouldBindJSON(&flight); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.Storage.UpdateFlight(ctx, id, flight); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update flight"})
		return
	}

	ctx.JSON(http.StatusOK, flight)
}

func (s *Server) handleDeleteFlight(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := s.Storage.DeleteFlight(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete flight"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Flight deleted successfully"})
}

func (s *Server) handleListFlights(ctx *gin.Context) {
	flights, err := s.Storage.ListFlights(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list flights"})
		return
	}

	ctx.JSON(http.StatusOK, flights)
}
