package location

import "fmt"

// Location is a struct that represents a location

type Location struct {

	// Latitude is the latitude of the location

	Latitude float64

	// Longitude is the longitude of the location

	Longitude float64
}

func (loc Location) String() string {
	return fmt.Sprintf("Location{Latitude: %f, Longitude: %f}", loc.Latitude, loc.Longitude)
}package location

import "fmt"

// Location is a struct that represents a location

type Location struct {

	// Latitude is the latitude of the location

	Latitude float64

	// Longitude is the longitude of the location

	Longitude float64
}

func (loc Location) String() string {
	return fmt.Sprintf("Location{Latitude: %f, Longitude: %f}", loc.Latitude, loc.Longitude)
}