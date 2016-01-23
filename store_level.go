package gopher_typer

import tl "github.com/JoelOtter/termloop"

type storeLevel struct {
	tl.Level
	gt *GopherTyper
}

func (l *storeLevel) Activate() {
	l.gt.console.SetText("Store Level")
	l.gt.g.Screen().SetLevel(l)
}

func NewStoreLevel(g *GopherTyper, fg, bg tl.Attr) storeLevel {
	l := tl.NewBaseLevel(tl.Cell{Bg: bg, Fg: fg})
	return storeLevel{l, g}
}
