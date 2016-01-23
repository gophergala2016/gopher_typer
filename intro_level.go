package gopher_typer

import (
	"io/ioutil"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type introLevel struct {
	tl.Level
	gt            *GopherTyper
	pressAKeyText *tl.Text
	totalTime     float64
}

func (l *introLevel) Activate() {
	l.gt.console.SetText("Intro Level")
	w, h := l.gt.g.Screen().Size()
	quarterW := w / 4
	quarterH := h / 4
	rect := tl.NewRectangle(quarterW, quarterH, quarterW*2, quarterH*2, tl.ColorCyan)
	l.AddEntity(rect)

	logo, _ := ioutil.ReadFile("data/logo.txt")
	logoEntity := tl.NewEntityFromCanvas(quarterW, quarterH, tl.CanvasFromString(string(logo)))
	l.AddEntity(logoEntity)

	msg := "Press any key to continue"
	l.pressAKeyText = tl.NewText(w/2-len(msg)/2, h/2, msg, tl.ColorBlue|tl.AttrReverse, tl.ColorDefault)
	l.AddEntity(l.pressAKeyText)

	l.gt.g.Screen().SetLevel(l)
}

func (l *introLevel) Update(dt time.Duration) {
	l.totalTime += dt.Seconds()
	if int(l.totalTime*2)%2 == 1 {
		l.pressAKeyText.SetColor(tl.ColorBlue, tl.ColorDefault)
	} else {
		l.pressAKeyText.SetColor(tl.ColorBlue|tl.AttrReverse, tl.ColorDefault)
	}
}

func (l *introLevel) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		l.gt.GoToGame()
	}
}

func NewIntroLevel(g *GopherTyper, fg, bg tl.Attr) introLevel {
	l := tl.NewBaseLevel(tl.Cell{Bg: bg, Fg: fg})
	return introLevel{l, g, nil, 0}
}
