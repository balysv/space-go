package data

import (
	"math"
)

type Vec2 struct {
	Point
}

func ZeroVec2() Vec2 {
	return Vec2{ZeroPoint()}
}

func UnitVector(p1 *Point, p2 *Point) Vec2 {
	vector := Vec2{Point{p2.X - p1.X, p2.Y - p1.Y}}
	return vector.Unit()
}

func (v *Vec2) Unit() Vec2 {
	return Vec2{v.Point.Div(v.Magnitude())}
}

func (v *Vec2) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2.0) + math.Pow(v.Y, 2.0))
}

func (v *Vec2) AngleTo(v2 *Vec2) float64 {
	return math.Acos(v.Times(v2) / (v.Magnitude() * v2.Magnitude()))
}

func (v *Vec2) Times(v2 *Vec2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v *Vec2) TimesScalar(a float64) Vec2 {
	return Vec2{v.Point.Times(a)}
}

func (v *Vec2) Plus(v2 *Vec2) Vec2 {
	return Vec2{v.Point.Plus(&v2.Point)}
}

func (v *Vec2) Minus(v2 *Vec2) Vec2 {
	return Vec2{v.Point.Minus(&v2.Point)}
}

func (v *Vec2) Div(a float64) Vec2 {
	return Vec2{v.Point.Div(a)}
}

func (v *Vec2) UnaryMinus() Vec2 {
	return Vec2{v.Point.UnaryMinus()}
}
