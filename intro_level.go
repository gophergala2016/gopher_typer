package gopher_typer

import tl "github.com/JoelOtter/termloop"

type introLevel struct {
	tl.Level
	gt *GopherTyper
}

func (l *introLevel) Activate() {
	l.gt.console.SetText("Intro Level")
	l.gt.g.Screen().SetLevel(l)
}

func NewIntroLevel(g *GopherTyper, fg, bg tl.Attr) introLevel {
	l := tl.NewBaseLevel(tl.Cell{Bg: bg, Fg: fg})
	return introLevel{l, g}
}
