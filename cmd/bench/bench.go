// LennardJones simulates molecular dynamics with the Lennard Jones potential.
package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/quells/LennardJonesGo/sim"
)

// Globals holds global simulation constants
type Globals struct {
	N           int
	L, M, T0, h float64
}

func main() {
	g := Globals{
		N: 2048, L: 12.6992084,
		T0: 0.2, M: 48.0,
		h: 0.01,
	}

	Rs := sim.InitPositionFCC(g.N, g.L)
	Vs := sim.InitVelocity(g.N, g.T0, g.M)

	numSteps := 100
	start := time.Now()
	for t := 1; t <= numSteps; t++ {
		Rs, Vs = sim.TimeStepParallel(Rs, Vs, g.L, g.M, g.h)
	}
	elapsed := time.Since(start)
	numCPUs := runtime.NumCPU()
	fmt.Printf("%v for %d time steps with %d particles using %d CPUs\n", elapsed, numSteps, g.N, numCPUs)
}
