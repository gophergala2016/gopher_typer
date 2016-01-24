package gopher_typer

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type gameLevel struct {
	tl.Level
	gt                   *GopherTyper
	fg                   tl.Attr
	bg                   tl.Attr
	words                []*word
	currentWord          *word
	currentWordText      *tl.Text
	garbageText          *tl.Text
	garbageCollectEndsAt time.Time
}

func (l *gameLevel) Activate() {
	l.Level = tl.NewBaseLevel(tl.Cell{Bg: l.bg, Fg: l.fg})

	l.gt.stats.Garbage = 0

	l.gt.game.AddEntity(&l.gt.console)
	l.gt.console.SetText("")

	numWords := l.gt.stats.LevelsCompleted + 1
	w, h := l.gt.g.Screen().Size()
	l.words = []*word{}

	x := 0
	for i := 0; i < numWords; i++ {
		w := NewWord(x, 0, l.gt.wordList[rand.Intn(len(l.gt.wordList))], tl.ColorRed, tl.ColorGreen, tl.ColorBlue, tl.ColorCyan)
		l.AddEntity(w)
		l.words = append(l.words, w)
		x += len(w.str) + 2
	}
	l.currentWord = nil
	l.currentWordText = tl.NewText(0, h-1, "", tl.ColorRed, tl.ColorBlue)
	l.AddEntity(l.currentWordText)

	l.garbageText = tl.NewText(w, h-1, "", tl.ColorRed, tl.ColorBlue)
	l.AddEntity(l.garbageText)

	l.AddEntity(tl.NewText(0, h-2, strings.Repeat("*", w), tl.ColorBlack, tl.ColorDefault))
	for _, i := range l.gt.items {
		i.Reset(l.gt)
	}

	l.gt.g.Screen().SetLevel(l)
}

func (l *gameLevel) Draw(screen *tl.Screen) {
	l.Level.Draw(screen)

	if time.Now().After(l.garbageCollectEndsAt) {
		for _, i := range l.gt.items {
			i.Tick(l)
		}
	}

	sw, sh := screen.Size()
	gameLost := false
	gameWon := false
	totalComplete := 0
	for _, w := range l.words {
		w.Update()
		if !w.Complete() && w.y > sh-3 {
			gameLost = true
		}
		if w.Complete() {
			totalComplete++
		}
	}
	if totalComplete == len(l.words) {
		gameWon = true
	}
	if l.currentWord != nil && l.currentWord.Complete() {
		l.currentWord = nil
	}
	var possibleWords []int
	for i, w := range l.words {
		if !w.Complete() && w.startedBy == 0 {
			possibleWords = append(possibleWords, i)
		}
	}

	if l.currentWord == nil && len(possibleWords) > 0 {
		l.currentWord = l.words[possibleWords[rand.Intn(len(possibleWords))]]
		l.currentWord.startedBy = PC
	}

	if l.gt.stats.GarbageCollect() {
		l.gt.stats.Garbage = 0
		l.garbageCollectEndsAt = time.Now().Add(time.Second * 3)

	}
	var msg string
	if time.Now().Before(l.garbageCollectEndsAt) {
		msg = fmt.Sprintf("COLLECTING GARBAGE: PLEASE WAIT")
		l.garbageText.SetText(msg)
		bgColor := tl.ColorBlue
		if math.Remainder(float64(time.Now().Sub(l.garbageCollectEndsAt)), float64(time.Second)) > 0.5 {
			bgColor = tl.ColorBlack
		}
		l.garbageText.SetColor(tl.ColorRed, bgColor)
		l.garbageText.SetPosition(sw/2-len(msg)/2, 4)
	} else {
		msg = fmt.Sprintf("Garbage: %d ", l.gt.stats.Garbage)
		l.garbageText.SetText(msg)
		l.garbageText.SetColor(tl.ColorRed, tl.ColorBlue)
		l.garbageText.SetPosition(sw-len(msg), sh-1)
	}

	if l.currentWord != nil {
		l.currentWordText.SetText("Current Word: " + l.currentWord.str[l.currentWord.completedChars:])
	} else {
		l.currentWordText.SetText("")
	}
	// End conditions
	if gameWon {
		l.gt.GoToEndWin()
	}
	if gameLost {
		l.gt.GoToEndFail()
	}
}

func (l *gameLevel) Tick(e tl.Event) {
	if e.Type == tl.EventKey {
		if l.currentWord != nil {
			l.currentWord.KeyDown(e.Ch)
		}
	}
}

func NewGameLevel(g *GopherTyper, fg, bg tl.Attr) gameLevel {
	return gameLevel{gt: g, fg: fg, bg: bg}
}
