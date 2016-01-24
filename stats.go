package gopher_typer

import "math/rand"

type stats struct {
	LevelsCompleted int
	LevelsAttempted int
	Dollars         int
	TotalEarned     int
	CpuUpgrades     int
	GoVersion       float32
	Lives           int
	Garbage         int
	GarbageFreq     int
}

func NewStats() stats {
	return stats{Lives: 3, CpuUpgrades: 1, GarbageFreq: 10, GoVersion: 1.0}

}

func (s *stats) GarbageCollect() bool {
	if s.Garbage > 0 && rand.Intn(s.Garbage) > s.GarbageFreq {
		return true
	}
	return false
}
