package main

import (
	"fmt"
	"go_game/location"
)

func main() {
	fmt.Println("Hello, World!")

	loc := location.Location{Latitude: 37.7749, Longitude: -122.4194}

	fmt.Println(loc.String())
}