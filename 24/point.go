package main

import (
	"math"
)

type Point struct {
	x int
	y int
}

// New создает экземпляр структуры Point
func New(x, y int) *Point {
	return &Point{x, y}
}

// X возвращает значение координаты х
func (p *Point) X() int {
	return p.x
}

// Y возвращает значение координаты у
func (p *Point) Y() int {
	return p.y
}

// Distance вычисляет расстояние между двумя точками по формуле: sqrt((xb - xa)^2 + (ya - yb)^2)
func Distance(p1 *Point, p2 *Point) float64 {
	xDiff := math.Pow(float64(p1.X()-p2.X()), 2)
	yDiff := math.Pow(float64(p1.Y()-p2.Y()), 2)

	d := math.Sqrt(xDiff + yDiff)

	return d
}
