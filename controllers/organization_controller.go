package controllers

import (
	"go-api-server/models"
	"go-api-server/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrganization(c *gin.Context) {
    var organization models.Organization
    if err := c.ShouldBindJSON(&organization); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := repositories.CreateOrganization(&organization); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, organization)
}
