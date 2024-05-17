package routes

import (
	"go-api-server/controllers"
	"go-api-server/middlewares"

	"github.com/gin-gonic/gin"
)

func OrganizationRoutes(router *gin.Engine) {
    orgGroup := router.Group("/organizations")
    orgGroup.Use(middlewares.AuthMiddleware())
    {
        orgGroup.POST("/", controllers.CreateOrganization)
    }
}
