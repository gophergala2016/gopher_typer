package gopher_typer

import (
	"time"

	tl "github.com/JoelOtter/termloop"
)

type endLevel struct {
	tl.Level
	gt *GopherTyper
}

func (l *endLevel) Activate() {
	l.gt.stats.LevelCompleted++
	l.gt.stats.Dollars += 1000
	l.gt.console.SetText("End Level")
	l.gt.g.Screen().SetLevel(l)
}

func (l *endLevel) Update(dt time.Duration) {

}

func NewEndLevel(g *GopherTyper, fg, bg tl.Attr) endLevel {
	l := tl.NewBaseLevel(tl.Cell{Bg: bg, Fg: fg})
	return endLevel{l, g}
}
