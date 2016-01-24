package gopher_typer

import (
	"fmt"
	"io/ioutil"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type endLevel struct {
	tl.Level
	fg tl.Attr
	bg tl.Attr
	gt *GopherTyper

	win      bool
	tickWait time.Time

	endMessages     []*tl.Entity
	currentMessage  int
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

func (l *endLevel) PrintStats(amt, x, y int) {
	msg := fmt.Sprintf("Levels Complete: %d", l.gt.stats.LevelsCompleted)
	text := tl.NewText(x-len(msg)/2, y, msg, tl.ColorBlack, tl.ColorDefault)
	l.AddEntity(text)

	msg = fmt.Sprintf("Levels Attempted: %d", l.gt.stats.LevelsAttempted)
	text = tl.NewText(x-len(msg)/2, y, msg, tl.ColorBlack, tl.ColorDefault)
	l.AddEntity(text)

	msg = fmt.Sprintf("Cash Earned: $%d", amt)
	text = tl.NewText(x-len(msg)/2, y+2, msg, tl.ColorBlack, tl.ColorDefault)
	l.AddEntity(text)

	msg = fmt.Sprintf("Total Cash: $%d", l.gt.stats.Dollars)
	text = tl.NewText(x-len(msg)/2, y+3, msg, tl.ColorBlack, tl.ColorDefault)
	l.AddEntity(text)

	if l.win {
		msg = fmt.Sprintf("Press N for next level or S for store")
	} else {
		msg = fmt.Sprintf("Press R to retry level or S for store")
	}
	text = tl.NewText(x-len(msg)/2, y+5, msg, tl.ColorBlack, tl.ColorDefault)
	l.AddEntity(text)
}

func (l *endLevel) ActivateWin() {
	l.Level = tl.NewBaseLevel(tl.Cell{Bg: l.bg, Fg: l.fg})
	l.Level.AddEntity(&l.gt.console)

	l.win = true

	moneyEarned := 1000
	l.gt.stats.LevelsCompleted++
	l.gt.stats.LevelsAttempted++
	l.gt.stats.Dollars += moneyEarned
	l.gt.console.SetText("")
	w, h := l.gt.g.Screen().Size()
	quarterW := w / 4
	quarterH := h / 4
	rect := tl.NewRectangle(quarterW, quarterH, quarterW*2, quarterH*2, tl.ColorCyan)
	l.AddEntity(rect)

	l.endMessages = []*tl.Entity{}
	l.addEndMessage("data/you_win_a.txt", quarterW, quarterH)
	l.addEndMessage("data/you_win_b.txt", quarterW, quarterH)
	l.AddEntity(l.endMessages[l.currentMessage])

	l.PrintStats(moneyEarned, quarterW*2, quarterH*2)

	l.Activate()
}

func (l *endLevel) ActivateFail() {
	l.Level = tl.NewBaseLevel(tl.Cell{Bg: l.bg, Fg: l.fg})
	l.AddEntity(&l.gt.console)
	l.gt.console.SetText("")
	l.win = false
	l.gt.stats.LevelsAttempted++

	w, h := l.gt.g.Screen().Size()
	quarterW := w / 4
	quarterH := h / 4
	rect := tl.NewRectangle(quarterW/2, quarterH/2, quarterW*3, quarterH*3, tl.ColorCyan)
	l.AddEntity(rect)

	l.endMessages = []*tl.Entity{}
	l.addEndMessage("data/you_loose_a.txt", quarterW/2, quarterH)
	l.addEndMessage("data/you_loose_b.txt", quarterW/2, quarterH)
	l.AddEntity(l.endMessages[l.currentMessage])

	l.PrintStats(0, quarterW*2, quarterH*2)

	l.Activate()
}

func (l *endLevel) Activate() {
	l.gt.g.Screen().SetLevel(l)
	l.tickWait = time.Now().Add(time.Second)
}

func (l *endLevel) Draw(screen *tl.Screen) {
	l.Level.Draw(screen)

	if time.Now().After(l.swapMessageTime) {
		lastMessage := l.currentMessage
		l.swapMessageTime = time.Now().Add(250 * time.Millisecond)
		l.currentMessage = (l.currentMessage + 1) % len(l.endMessages)
		l.RemoveEntity(l.endMessages[lastMessage])
		l.AddEntity(l.endMessages[l.currentMessage])
	}

}

func (l *endLevel) Tick(e tl.Event) {
	if time.Now().After(l.tickWait) && e.Type == tl.EventKey {
		if e.Ch == 'N' || e.Ch == 'n' || e.Ch == 'R' || e.Ch == 'r' {
			l.gt.GoToGame()
		} else if e.Ch == 'S' || e.Ch == 's' {
			l.gt.GoToStore()
		}
	}
}

func NewEndLevel(g *GopherTyper, fg, bg tl.Attr) endLevel {
	return endLevel{gt: g, fg: fg, bg: bg}
}
