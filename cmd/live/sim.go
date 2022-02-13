package main

import (
	"github.com/quells/LennardJonesGo/sim"
)

type Sim struct {
	N, Steps    int
	L, M, T0, h float64

	Rs, Vs [][3]float64
}

func DefaultSim() *Sim {
	const (
		N  = 2048 / 8
		L  = 12.6992084 / 2 * 1.1
		T0 = 0.2
		M  = 48.0
		h  = 0.01
	)
	return NewSim(N, L, M, T0, h, sim.InitPositionFCC)
}

func NewSim(N int, L, M, T0, h float64, initPos sim.InitPosition) *Sim {
	s := Sim{N: N, L: L, M: M, T0: T0, h: h}
	s.Rs = initPos(N, L)
	s.Vs = sim.InitVelocity(N, T0, M)
	return &s
}

func (s *Sim) Step() {
	s.Rs, s.Vs = sim.TimeStepParallel(s.Rs, s.Vs, s.L, s.M, s.h)
	s.Steps++
}

type byZ [][3]float64

func (a byZ) Len() int           { return len(a) }
func (a byZ) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byZ) Less(i, j int) bool { return a[i][2] < a[j][2] }
