package main

import (
	"go-api-server/config"
	"go-api-server/models"
	"go-api-server/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    config.ConnectDB()
    config.DB.AutoMigrate(&models.User{}, &models.Organization{})

    router := gin.Default()
    routes.UserRoutes(router)
    routes.OrganizationRoutes(router)

    router.Run(os.Getenv("PORT"))
}
