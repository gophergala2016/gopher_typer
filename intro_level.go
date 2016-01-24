package gopher_typer

import (
	"io/ioutil"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type introLevel struct {
	tl.Level
	gt              *GopherTyper
	pressAKeyText   *tl.Text
	needsRefresh    bool
	swapMessageTime time.Time
	reverseText     bool
}

func (l *introLevel) Activate() {
	l.needsRefresh = true
	l.gt.g.Screen().SetLevel(l)
}
func (l *introLevel) refresh() {
	l.gt.intro.AddEntity(&l.gt.console)
	l.gt.console.SetText("")
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

	instructions, _ := ioutil.ReadFile("data/instructions.txt")
	l.AddEntity(tl.NewEntityFromCanvas(quarterW, h/2+2, tl.CanvasFromString(string(instructions))))

	l.needsRefresh = false
}

func (l *introLevel) Draw(screen *tl.Screen) {
	if l.needsRefresh {
		l.refresh()
	}
	if time.Now().After(l.swapMessageTime) {
		if l.reverseText {
			l.pressAKeyText.SetColor(tl.ColorBlue, tl.ColorDefault)
		} else {
			l.pressAKeyText.SetColor(tl.ColorBlue|tl.AttrReverse, tl.ColorDefault)
		}
		l.reverseText = !l.reverseText
		l.swapMessageTime = time.Now().Add(500 * time.Millisecond)
	}
	l.Level.Draw(screen)
}

func (l *introLevel) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		l.gt.GoToGame()
	}
}

func NewIntroLevel(g *GopherTyper, fg, bg tl.Attr) introLevel {
	l := tl.NewBaseLevel(tl.Cell{Bg: bg, Fg: fg})
	return introLevel{Level: l, gt: g}
}
