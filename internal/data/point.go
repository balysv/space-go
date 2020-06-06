package data

import (
	"math"
)

type Point struct {
	X, Y float64
}

func ZeroPoint() Point {
	return Point{}
}

func (p *Point) DistanceTo(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p.X-p2.X, 2.0) + math.Pow(p.Y-p2.Y, 2.0))
}

func (p *Point) Times(a float64) Point {
	return Point{p.X * a, p.Y * a}
}

func (p *Point) Plus(p2 *Point) Point {
	return Point{p.X + p2.X, p.Y + p2.Y}
}

func (p *Point) Minus(p2 *Point) Point {
	return Point{p.X - p2.X, p.Y - p2.Y}
}

func (p *Point) Div(a float64) Point {
	return Point{p.X / a, p.Y / a}
}

func (p *Point) UnaryMinus() Point {
	return Point{-p.X, - p.Y}
}
