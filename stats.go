package gopher_typer

type stats struct {
	LevelsCompleted int
	LevelsAttempted int
	Dollars         int
	CpuUpgrades     int
	GoVersion       int
	Lives           int
}

func NewStats() stats {
	return stats{Lives: 3, CpuUpgrades: 1}

}
