package sim

import (
	"fmt"
	"github.com/quells/LennardJones/space"
	"github.com/quells/LennardJones/vector"
	"math"
)

// PairwiseLennardJonesForce calculates the force vector on particle Ri due to Rj using the Lennard Jones potential.
func PairwiseLennardJonesForce(Ri, Rj [3]float64, L float64) [3]float64 {
	if space.PointsAreEqual(Ri, Rj, L) {
		panic(fmt.Sprintf("%v and %v are equal, the pairwise force is infinite", Ri, Rj))
	}
	r := space.Displacement(Ri, Rj, L)
	mag_r := vector.Length(r)
	f := 4 * (-12*math.Pow(mag_r, -13) + 6*math.Pow(mag_r, -7))
	return vector.Scale(r, f/mag_r)
}

// InternalForce calculates the total force vector on particle Ri due to the other particles in R due to a pairwise force.
func InternalForce(Ri [3]float64, R [][3]float64, L float64) [3]float64 {
	F := [3]float64{0, 0, 0}
	for _, Rj := range R {
		if !space.PointsAreEqual(Ri, Rj, L) {
			F = vector.Sum(F, PairwiseLennardJonesForce(Ri, Rj, L))
		}
	}
	return F
}
