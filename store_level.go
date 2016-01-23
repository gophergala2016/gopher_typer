package gopher_typer

import (
	"time"

	tl "github.com/JoelOtter/termloop"
)

type storeLevel struct {
	tl.Level
	gt *GopherTyper
	bg tl.Attr
	fg tl.Attr
}

func (l *storeLevel) Activate() {
	l.Level = tl.NewBaseLevel(tl.Cell{Bg: l.bg, Fg: l.fg})
	l.gt.store.AddEntity(&l.gt.console)
	l.gt.console.SetText("Store Level")
	l.gt.g.Screen().SetLevel(l)
}

func (l *storeLevel) Update(dt time.Duration) {

}

func NewStoreLevel(g *GopherTyper, fg, bg tl.Attr) storeLevel {
	return storeLevel{gt: g, bg: bg, fg: fg}
}
