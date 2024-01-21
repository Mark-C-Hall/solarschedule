package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mark-c-hall/solarschedule/internal/model"
)

func GetSunTimes(c *gin.Context) {
	city := c.Query("city")
	state := c.Query("state")
	zip := c.Query("zip")

	// Validation: Either city and state should be provided, or zip
	if (city == "" || state == "") && zip == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide either city and state or zip code",
		})
		return
	}

	// Dummy data for the purpose of this example
	location := model.Location{Latitude: 36.7201600, Longitude: -4.4203400}
	date := time.Now()
	sunrise := "06:24:00"
	sunset := "18:46:00"

	// Instantiate the NewSunTimes model
	sunTimes := model.NewSunTimes(date, sunrise, sunset, location)

	// Return the response with the sunrise and sunset times
	c.JSON(http.StatusOK, gin.H{
		"city":    city,
		"state":   state,
		"zip":     zip,
		"sunrise": sunTimes.Sunrise,
		"sunset":  sunTimes.Sunset,
	})
}
