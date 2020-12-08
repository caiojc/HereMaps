package models

import "fmt"

//Point : Model com a latitude e a longitude
type Point struct {
	Latitude  float64
	Longitude float64
}

//ToString : Converte em string o Point passado
func (point *Point) ToString() string {
	return fmt.Sprintf("%f,%f", point.Latitude, point.Longitude)
}
