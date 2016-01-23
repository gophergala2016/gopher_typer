package gopher_typer

import (
	"time"

	tl "github.com/JoelOtter/termloop"
)

type introLevel struct {
	tl.Level
	gt *GopherTyper
}

func (l *introLevel) Activate() {
	l.gt.console.SetText("Intro Level")
	w, h := l.gt.g.Screen().Size()
	quarterW := w / 4
	quarterH := h / 4
	rect := tl.NewRectangle(quarterW, quarterH, quarterW*2, quarterH*2, tl.ColorCyan)
	l.AddEntity(rect)

	msg := "Press any key to continue"
	text := tl.NewText(w/2-len(msg)/2, h/2, msg, tl.ColorBlue, tl.ColorDefault)
	l.AddEntity(text)

	l.gt.g.Screen().SetLevel(l)
}

func (l *introLevel) Update(dt time.Duration) {

}

func (l *introLevel) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		l.gt.GoToGame()
	}
}

func NewIntroLevel(g *GopherTyper, fg, bg tl.Attr) introLevel {
	l := tl.NewBaseLevel(tl.Cell{Bg: bg, Fg: fg})
	return introLevel{l, g}
}
