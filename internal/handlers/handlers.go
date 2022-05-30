package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/worldwidepaniel/weather-app/internal/auth"
)

func AuthHandler(c *gin.Context) {
	token, err := auth.GenerateJWT()
	if err != nil {
		fmt.Printf("Something went wrong while handling auth: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"jwt": token,
		})

	}
}
