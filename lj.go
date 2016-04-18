// LennardJones simulates molecular dynamics with the Lennard Jones potential.
package main

import (
	//"fmt"
	"github.com/quells/LennardJonesGo/sim"
)

// Globals holds global simulation constants
type Globals struct {
	N           int
	L, M, T0, h float64
}

func main() {
	g := Globals{
		N: 2048, L: 12.6992084,//6.3496042,//3.1748021,//4.2323167,
		T0: 0.0, M: 48.0,
		h: 0.01,
	}

	Rs := sim.InitPositionFCC(g.N, g.L)
	Vs := sim.InitVelocity(g.N, g.T0, g.M)

	T := 1000
	for t := 1; t <= T; t++ {
		Rs, Vs = sim.TimeStepParallel(Rs, Vs, g.L, g.M, g.h)
		//fmt.Println(sim.TotalEnergy(Rs, Vs, g.L, g.M))
		//fmt.Println(sim.Temperature(Vs, g.M, g.N))
	}
}
