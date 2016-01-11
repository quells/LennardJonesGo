// LennardJones simulates many bodied particle systems with the Lennard Jones potential.
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
	// fmt.Println("LennardJones v1.0.0")
	g := Globals{
		N: 64, L: 4.2323167,
		T0: 0.728, M: 48.0,
		h: 0.01,
	}

	Rs := sim.InitPositionCubic(g.N, g.L)
	Vs := sim.InitVelocity(g.N, g.T0, g.M)

	T := 1000
	// fmt.Printf("0 %f\n", sim.TotalEnergy(Rs, Vs, g.L, g.M))
	for t := 1; t < T; t++ {
		Rs, Vs = sim.TimeStep(Rs, Vs, g.L, g.M, g.h)
		// fmt.Printf("%d %f\n", t, sim.TotalEnergy(Rs, Vs, g.L, g.M))
		// fmt.Printf("%d %f\n", t, sim.Temperature(Vs, g.M, g.N))
	}
}
