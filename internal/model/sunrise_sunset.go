package model

import "time"

// Location represents geographical coordinates with latitude and longitude
type Location struct {
	Latitude  float64
	Longitude float64
}

// SunTimes holds the information about the sunrise and sunset times for a given location and date
type SunTimes struct {
	Date     time.Time
	Sunrise  string
	Sunset   string
	Location Location
}

// NewSunTimes creates a new SunTimes instance
func NewSunTimes(date time.Time, sunrise, sunset string, location Location) *SunTimes {
	return &SunTimes{
		Date:     date,
		Sunrise:  sunrise,
		Sunset:   sunset,
		Location: location,
	}
}
