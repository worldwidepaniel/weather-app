package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/worldwidepaniel/weather-app/internal/config"
	"github.com/worldwidepaniel/weather-app/internal/handlers"
)

func ConfigRoutes() {
	router := gin.Default()

	router.GET("/auth", handlers.AuthHandler)

	router.Run(config.PORT)
}
