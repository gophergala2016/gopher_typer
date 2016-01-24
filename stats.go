package gopherTyper

import "math/rand"

type stats struct {
	LevelsCompleted int
	LevelsAttempted int
	Dollars         int
	TotalEarned     int
	CPUUpgrades     int
	GoVersion       float32
	Lives           int
	Garbage         int
	GarbageFreq     int
}

func newStats() stats {
	return stats{Lives: 3, CPUUpgrades: 1, GarbageFreq: 10, GoVersion: 1.0}
}

func (s *stats) GarbageCollect() bool {
	if s.Garbage > 0 && rand.Intn(s.Garbage) > s.GarbageFreq {
		return true
	}
	return false
}
