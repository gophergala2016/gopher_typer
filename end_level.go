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
	c := tl.CanvasFromString(string(data))

	l.endMessages = append(l.endMessages, tl.NewEntityFromCanvas(x-len(c)/2, y, c))
}

func (l *endLevel) PrintStats(amt, x, y int) {
	msg := fmt.Sprintf("Levels Complete: %d", l.gt.stats.LevelsCompleted)
	text := tl.NewText(x-len(msg)/2, y, msg, tl.ColorBlack, tl.ColorDefault)
	l.AddEntity(text)
	y++

	msg = fmt.Sprintf("Levels Attempted: %d", l.gt.stats.LevelsAttempted)
	text = tl.NewText(x-len(msg)/2, y, msg, tl.ColorBlack, tl.ColorDefault)
	l.AddEntity(text)
	y++

	msg = fmt.Sprintf("Reward: $%d", amt)
	text = tl.NewText(x-len(msg)/2, y, msg, tl.ColorBlack, tl.ColorDefault)
	l.AddEntity(text)
	y++

	msg = fmt.Sprintf("Balance: $%d", l.gt.stats.Dollars)
	text = tl.NewText(x-len(msg)/2, y, msg, tl.ColorBlack, tl.ColorDefault)
	l.AddEntity(text)
	y++

	msg = fmt.Sprintf("Total Cash: $%d", l.gt.stats.TotalEarned)
	text = tl.NewText(x-len(msg)/2, y, msg, tl.ColorBlack, tl.ColorDefault)
	l.AddEntity(text)
	y++

	msg = fmt.Sprintf("Lives Remaining: %d", l.gt.stats.Lives)
	text = tl.NewText(x-len(msg)/2, y, msg, tl.ColorBlack, tl.ColorDefault)
	l.AddEntity(text)
	y++

	if l.win {
		msg = fmt.Sprintf("Press N for next level or S for store")
	} else if l.gt.stats.Lives > 0 {
		msg = fmt.Sprintf("Press R to retry level or S for store")
	} else {
		msg = fmt.Sprintf("Press Enter to quit or N for new game")
	}
	text = tl.NewText(x-len(msg)/2, y+1, msg, tl.ColorBlack, tl.ColorDefault)
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
	l.gt.stats.TotalEarned += moneyEarned
	l.gt.console.SetText("")
	w, h := l.gt.g.Screen().Size()
	rect := tl.NewRectangle(10, 2, w-20, h-4, tl.ColorCyan)
	l.AddEntity(rect)

	l.endMessages = []*tl.Entity{}
	l.addEndMessage("data/you_win_a.txt", w/2, 3)
	l.addEndMessage("data/you_win_b.txt", w/2, 3)
	l.AddEntity(l.endMessages[l.currentMessage])

	l.PrintStats(moneyEarned, w/2, 13)

	l.Activate()
}

func (l *endLevel) ActivateFail() {
	l.win = false
	l.gt.stats.LevelsAttempted++
	l.gt.stats.Lives--
	if l.gt.stats.Lives == 0 {
		l.ActivateGameOver()
		return
	}
	l.Level = tl.NewBaseLevel(tl.Cell{Bg: l.bg, Fg: l.fg})
	l.AddEntity(&l.gt.console)
	l.gt.console.SetText("")

	w, h := l.gt.g.Screen().Size()
	rect := tl.NewRectangle(10, 2, w-20, h-4, tl.ColorCyan)
	l.AddEntity(rect)

	l.endMessages = []*tl.Entity{}
	l.addEndMessage("data/you_loose_a.txt", w/2, 3)
	l.addEndMessage("data/you_loose_b.txt", w/2, 3)
	l.AddEntity(l.endMessages[l.currentMessage])

	l.PrintStats(0, w/2, 13)

	l.Activate()
}

func (l *endLevel) ActivateGameOver() {
	l.Level = tl.NewBaseLevel(tl.Cell{Bg: l.bg, Fg: l.fg})
	l.AddEntity(&l.gt.console)
	l.gt.console.SetText("")
	l.gt.g.SetEndKey(tl.KeyEnter)

	w, h := l.gt.g.Screen().Size()
	rect := tl.NewRectangle(10, 2, w-20, h-4, tl.ColorCyan)
	l.AddEntity(rect)

	l.endMessages = []*tl.Entity{}
	l.addEndMessage("data/game_over_a.txt", w/2, 3)
	l.addEndMessage("data/game_over_b.txt", w/2, 3)
	l.AddEntity(l.endMessages[l.currentMessage])

	l.PrintStats(0, w/2, 13)

	l.Activate()
}

func (l *endLevel) Activate() {
	l.gt.g.Screen().SetLevel(l)
	l.tickWait = time.Now().Add(500 * time.Millisecond)
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

func (l *endLevel) Tick(e tl.Event) {
	if time.Now().After(l.tickWait) && e.Type == tl.EventKey {
		if e.Ch == 'N' || e.Ch == 'n' || e.Ch == 'R' || e.Ch == 'r' {
			l.gt.g.SetEndKey(tl.KeyCtrlC)
			l.gt.GoToGame()
		} else if e.Ch == 'S' || e.Ch == 's' {
			l.gt.g.SetEndKey(tl.KeyCtrlC)
			l.gt.GoToStore()
		}
	}
}

func NewEndLevel(g *GopherTyper, fg, bg tl.Attr) endLevel {
	return endLevel{gt: g, fg: fg, bg: bg}
}
