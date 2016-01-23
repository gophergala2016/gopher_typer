package gopher_typer

import (
	"fmt"
	"io/ioutil"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type endLevel struct {
	tl.Level
	gt *GopherTyper

	endMessages     []*tl.Entity
	currentMessage  int
	totalTime       float64
	swapMessageTime time.Time
}

func (l *endLevel) addEndMessage(path string, x, y int) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		l.gt.console.SetText(fmt.Sprintf("Err: %+v", err))
		return
	}
	l.endMessages = append(l.endMessages, tl.NewEntityFromCanvas(x, y, tl.CanvasFromString(string(data))))

}

func (l *endLevel) ActivateWin() {
	l.gt.stats.LevelCompleted++
	l.gt.stats.Dollars += 1000
	l.gt.console.SetText("")
	w, h := l.gt.g.Screen().Size()
	quarterW := w / 4
	quarterH := h / 4
	rect := tl.NewRectangle(quarterW, quarterH, quarterW*2, quarterH*2, tl.ColorCyan)
	l.AddEntity(rect)

	l.addEndMessage("data/you_win_a.txt", quarterW, quarterH)
	l.addEndMessage("data/you_win_b.txt", quarterW, quarterH)
	l.AddEntity(l.endMessages[l.currentMessage])

	l.Activate()
}

func (l *endLevel) ActivateFail() {
	l.gt.console.SetText("")
	w, h := l.gt.g.Screen().Size()
	quarterW := w / 4
	quarterH := h / 4
	rect := tl.NewRectangle(quarterW/2, quarterH/2, quarterW*3, quarterH*3, tl.ColorCyan)
	l.AddEntity(rect)

	l.addEndMessage("data/you_loose_a.txt", quarterW/2, quarterH)
	l.addEndMessage("data/you_loose_b.txt", quarterW/2, quarterH)
	l.AddEntity(l.endMessages[l.currentMessage])

	l.Activate()
}

func (l *endLevel) Activate() {
	l.gt.g.Screen().SetLevel(l)
}

func (l *endLevel) Draw(screen *tl.Screen) {
	l.Level.Draw(screen)

	if time.Now().After(l.swapMessageTime) {
		lastMessage := l.currentMessage
		l.swapMessageTime = time.Now().Add(500 * time.Millisecond)
		l.currentMessage = (l.currentMessage + 1) % len(l.endMessages)
		l.RemoveEntity(l.endMessages[lastMessage])
		l.AddEntity(l.endMessages[l.currentMessage])
	}

}

func (l *endLevel) Update(dt time.Duration) {
	l.totalTime += dt.Seconds()
}

func NewEndLevel(g *GopherTyper, fg, bg tl.Attr) endLevel {
	l := tl.NewBaseLevel(tl.Cell{Bg: bg, Fg: fg})
	return endLevel{Level: l, gt: g}
}
