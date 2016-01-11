package sim

import (
	"github.com/quells/LennardJones/space"
	"github.com/quells/LennardJones/vector"
	"github.com/quells/LennardJones/verlet"
)

// TimeStep evolves the system by one unit of time using the Velocity Verlet algorithm for molecular dynamics.
func TimeStep(R, V [][3]float64, L, M, h float64) ([][3]float64, [][3]float64) {
	N := len(R)
	A := make([][3]float64, N)
	nR := make([][3]float64, N)
	nV := make([][3]float64, N)
	for i, _ := range R {
		Fi := InternalForce(R[i], R, L)
		A[i] = vector.Scale(Fi, 1.0/M)
		nR[i] = space.PutInBox(verlet.NextR(R[i], V[i], A[i], h), L)
	}
	for i, _ := range R {
		nFi := InternalForce(nR[i], nR, L)
		nAi := vector.Scale(nFi, 1.0/M)
		nV[i] = verlet.NextV(V[i], A[i], nAi, h)
	}
	return nR, nV
}
