package data

import (
	"math"
)

func NewEntity(name string, mass float64, radius float64, velocity Vec2, coords Point) *Entity {
	return &Entity{
		Name:     name,
		Mass:     mass,
		Radius:   radius,
		Density:  mass / (4 / 3) * math.Pi * math.Pow(radius, 3.0),
		Velocity: velocity,
		Coords:   coords,
	}
}

type Entity struct {
	Name                  string
	Mass, Radius, Density float64
	Velocity              Vec2
	Coords                Point
}

func Merge(e1 *Entity, e2 *Entity) *Entity {
	name := e1.Name + " & " + e2.Name
	mass := e1.Mass + e2.Mass
	radius := e1.Radius + e2.Radius

	relativeMass1 := e1.Mass / mass
	relativeMass2 := e2.Mass / mass

	relativeVelocity1 := e1.Velocity.TimesScalar(relativeMass1)
	relativeVelocity2 := e2.Velocity.TimesScalar(relativeMass2)
	velocity := relativeVelocity1.Plus(&relativeVelocity2)

	relativeCoords1 := e1.Coords.Times(relativeMass1)
	relativeCoords2 := e2.Coords.Times(relativeMass2)
	coords := relativeCoords1.Plus(&relativeCoords2)

	return NewEntity(name, mass, radius, velocity, coords)
}
