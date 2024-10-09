package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/themelancholyspirit/airline-reservation-system/types"
	"github.com/themelancholyspirit/airline-reservation-system/util"
)

func (s *Server) handleSignup(ctx *gin.Context) {
	var signupRequest types.UserSignupRequest
	if err := ctx.ShouldBindJSON(&signupRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := util.HashPassword(signupRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := types.User{
		Name:     signupRequest.Name,
		Email:    signupRequest.Email,
		Password: hashedPassword,
	}

	if err := s.Storage.CreateUser(ctx, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	token, err := util.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"token": token})
}

func (s *Server) handleLogin(ctx *gin.Context) {
	var loginRequest types.UserLoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.Storage.GetUserByEmail(ctx, loginRequest.Email)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !util.CheckPasswordHash(loginRequest.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := util.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
