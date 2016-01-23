package gopher_typer

import tl "github.com/JoelOtter/termloop"

type gameLevel struct {
	tl.Level
	gt *GopherTyper
}

func (l *gameLevel) Activate() {
	l.gt.console.SetText("Game Level")
	l.gt.g.Screen().SetLevel(l)

}

func NewGameLevel(g *GopherTyper, fg, bg tl.Attr) gameLevel {
	l := tl.NewBaseLevel(tl.Cell{Bg: bg, Fg: fg})
	return gameLevel{l, g}
}
