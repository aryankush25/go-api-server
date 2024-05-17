package routes

import (
	"go-api-server/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
    userGroup := router.Group("/users")
    {
        userGroup.POST("/signup-step1", controllers.SignupStep1)
        userGroup.POST("/signup-step2", controllers.SignupStep2)
    }
}
