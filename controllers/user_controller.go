package controllers

import (
	"go-api-server/models"
	"go-api-server/repositories"
	"go-api-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignupStep1(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    otp := services.GenerateOTP()
    user.OTP = otp

    if err := repositories.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    emailService, err := services.NewEmailService()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize email service"})
        return
    }

    err = emailService.SendEmail(user.Email, "Your OTP", otp)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP email"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "OTP sent to email"})
}


func SignupStep2(c *gin.Context) {
    var input struct {
        Email string
        OTP   string
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := repositories.FindUserByEmail(input.Email)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if user.OTP != input.OTP {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
        return
    }

    user.Verified = true
    if err := repositories.UpdateUser(&user); err != nil {  // Pass a pointer here
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User verified"})
}
