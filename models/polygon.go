package models

import "strings"

//Polygon : Model com array de Points para mostrar no mapa
type Polygon struct {
	Points []Point
}

//ToString : Converte em string o Polygon com o seus Points
func (polygon *Polygon) ToString() string {
	var result string
	var points []string
	for _, point := range polygon.Points {
		points = append(points, point.ToString())
	}
	result = strings.Join(points, ",")
	return result
}
