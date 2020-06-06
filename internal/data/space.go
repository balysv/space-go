package data

import (
	"math"
)

// Space
type Space struct {
	G        float64
	Entities []Entity
}

func NewSpace() *Space {
	return &Space{G: 6.674e-11}
}

func (space *Space) AddEntity(entity *Entity) {
	space.Entities = append(space.Entities, *entity)
}

func (space *Space) FindByName(name string) *Entity {
	for _, entity := range space.Entities {
		if entity.Name != name {
			continue
		}
		return &entity
	}
	return nil
}

func (space *Space) Step() {
	space.applyGravitationalForces()
	space.mergeCollidedEntities()
}

/**
 * Computes and applies gravitation forces between each entity
 */
func (space *Space) applyGravitationalForces() {
	for i := range space.Entities {
		e1 := &space.Entities[i]

		totalForce := ZeroVec2()
		for j := range space.Entities {
			if i == j {
				continue
			}
			e2 := &space.Entities[j]

			gravity := computeGravity(space.G, *e1, *e2)
			direction := UnitVector(&e1.Coords, &e2.Coords)

			force := direction.TimesScalar(gravity)
			totalForce = totalForce.Plus(&force)
		}

		acceleration := totalForce.Div(e1.Mass)
		e1.Velocity = e1.Velocity.Plus(&acceleration)
		e1.Coords = e1.Coords.Plus(&e1.Velocity.Point)
	}
}

func computeGravity(G float64, e1 Entity, e2 Entity) float64 {
	return G * (e1.Mass * e2.Mass) / math.Pow(e1.Coords.DistanceTo(&e2.Coords), 2.0)
}

/**
 * Merges existing entities with the provided ones by checking if position in space of two
 * entities is nearly the same. Combines their mass, radius and velocity vector
 */
func (space *Space) mergeCollidedEntities() {
	var currentEntities []*Entity
	for idx := range space.Entities {
		entity := space.Entities[idx]
		currentEntities = append(currentEntities, &entity)
	}
	var newEntities []*Entity

	for i := range currentEntities {
		for j := range currentEntities {
			e1 := currentEntities[i]
			e2 := currentEntities[j]
			if e1 == nil || e2 == nil  || e1 == e2 {
				continue
			}

			// skip if entities are too far apart
			distance := math.Abs(e1.Coords.DistanceTo(&e2.Coords))
			if distance >= e1.Radius+e2.Radius {
				continue
			}

			// add new entity
			mergedEntity := Merge(e1, e2)
			newEntities = append(newEntities, mergedEntity)
			// mark merged entities for removal
			currentEntities[i] = nil
			currentEntities[j] = nil
		}
	}

	space.Entities = nil

	// Add all entities that haven't merged
	for idx := range currentEntities {
		entity := currentEntities[idx]
		if entity == nil {
			continue
		}
		space.Entities = append(space.Entities, *entity)
	}

	// Add newly merged entities
	for idx := range newEntities {
		space.Entities = append(space.Entities, *newEntities[idx])
	}
}
