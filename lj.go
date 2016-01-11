// LennardJones simulates molecular dynamics with the Lennard Jones potential.
package main

import (
	// "fmt"
	"github.com/quells/LennardJones/sim"
)

type Globals struct {
	N           int
	L, M, T0, h float64
}

func main() {
	g := Globals{
		N: 64, L: 4.2323167,
		T0: 0.728, M: 48.0,
		h: 0.01,
	}

	Rs := sim.InitPositionCubic(g.N, g.L)
	Vs := sim.InitVelocity(g.N, g.T0, g.M)

	T := 1000
	for t := 1; t < T; t++ {
		Rs, Vs = sim.TimeStep(Rs, Vs, g.L, g.M, g.h)
	}
}
