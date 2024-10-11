package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/themelancholyspirit/airline-reservation-system/types"
	"github.com/themelancholyspirit/airline-reservation-system/util"
)

func (s *Server) handleGetProfile(ctx *gin.Context) {
	tokenString, exists := ctx.Get("token")
	fmt.Println("tokenString", tokenString)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found in context"})
		return
	}

	claims, err := util.ValidateToken(tokenString.(string))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	user, err := s.Storage.GetUser(ctx, claims.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user data"})
		return
	}

	profileResponse := types.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	ctx.JSON(http.StatusOK, profileResponse)
}

func (s *Server) handleUpdateProfile(ctx *gin.Context) {
	tokenString, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found in context"})
		return
	}

	claims, err := util.ValidateToken(tokenString.(string))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var updateRequest types.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = s.Storage.UpdateUser(ctx, updateRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user data"})
		return
	}

	profileResponse := types.User{
		ID:    uint(claims.UserID),
		Name:  updateRequest.Name,
		Email: updateRequest.Email,
	}
	ctx.JSON(http.StatusOK, profileResponse)
}

func (s *Server) handleDeleteProfile(ctx *gin.Context) {
	tokenString, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found in context"})
		return
	}

	claims, err := util.ValidateToken(tokenString.(string))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	err = s.Storage.DeleteUser(ctx, claims.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User profile deleted successfully"})
}
