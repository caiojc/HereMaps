package main

import (
	"HereMaps/engine"
	"HereMaps/models"
)

func main() {
	polygon := models.Polygon{
		Points: []models.Point{
			models.Point{Latitude: 37.7397, Longitude: -121.4252},
			models.Point{Latitude: 37.7974, Longitude: -121.2161},
			models.Point{Latitude: 37.6391, Longitude: -120.9969},
			models.Point{Latitude: 37.7397, Longitude: -121.4252},
		},
	}

	engine.GetMapWithPolygon("map.jpg", polygon.ToString())
}
