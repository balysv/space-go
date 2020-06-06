package data

import (
	"math"
	"testing"
)

func roundToOneDecimalPlace(a float64) float64 {
	return math.Round(a*10) / 10
}

func TestGravityOnEarth(t *testing.T) {
	space := NewSpace()
	space.AddEntity(NewEntity("Earth", 5.972e24, 6.371e6, ZeroVec2(), ZeroPoint()))
	space.AddEntity(NewEntity("Apple", 0.01, 0.05, ZeroVec2(), Point{0, 6.372e6}))

	space.Step()

	apple := space.FindByName("Apple")
	roundedAppleVelocity := Vec2{Point{
		X: roundToOneDecimalPlace(apple.Velocity.X),
		Y: roundToOneDecimalPlace(apple.Velocity.Y),
	}}

	gravityOnEarth := Vec2{Point{0, -9.8}}

	if roundedAppleVelocity != gravityOnEarth {
		t.Logf("Gravity on Earth expected to be %v, but got %v", gravityOnEarth, roundedAppleVelocity)
		t.Fail()
	}
}

func TestCollideEarthIntoTheSun(t *testing.T) {
	space := NewSpace()
	space.AddEntity(NewEntity("Earth", 5.972e24, 6.371e6, ZeroVec2(), ZeroPoint()))
	space.AddEntity(NewEntity("Sun", 2e30, 6.957e8, ZeroVec2(), Point{10e5, 10e5}))

	space.Step()

	if len(space.Entities) > 1 {
		t.Logf("Expected a single entity after collision, but got %v", space.Entities)
		t.Fail()
	}

	mergedEntity := space.FindByName("Earth & Sun")

	if mergedEntity == nil {
		t.Logf("Could not find mergedEntity, have %v", space.Entities)
		t.Fail()
	}
}
