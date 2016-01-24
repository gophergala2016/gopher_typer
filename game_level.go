package gopher_typer

import (
	"log"
	"math/rand"
	"strings"

	tl "github.com/JoelOtter/termloop"
)

type gameLevel struct {
	tl.Level
	gt              *GopherTyper
	fg              tl.Attr
	bg              tl.Attr
	words           []*word
	currentWord     *word
	currentWordText *tl.Text
}

func (l *gameLevel) Activate() {
	l.Level = tl.NewBaseLevel(tl.Cell{Bg: l.bg, Fg: l.fg})

	l.gt.game.AddEntity(&l.gt.console)
	l.gt.console.SetText("")

	numWords := l.gt.stats.LevelsCompleted + 1
	w, h := l.gt.g.Screen().Size()
	l.words = []*word{}

	for i := 0; i < numWords; i++ {
		w := NewWord(w/numWords*i, 0, l.gt.wordList[rand.Intn(len(l.gt.wordList))], tl.ColorRed, tl.ColorGreen, tl.ColorBlue, tl.ColorCyan)
		l.AddEntity(w)
		l.words = append(l.words, w)
	}
	l.currentWordText = tl.NewText(0, h-1, "", tl.ColorRed, tl.ColorBlue)
	l.AddEntity(l.currentWordText)

	l.AddEntity(tl.NewText(0, h-2, strings.Repeat("*", w), tl.ColorBlack, tl.ColorDefault))
	for _, i := range l.gt.items {
		i.Reset(l)
	}

	l.gt.g.Screen().SetLevel(l)
}

func (l *gameLevel) Draw(screen *tl.Screen) {
	l.Level.Draw(screen)

	for idx, i := range l.gt.items {
		i.Tick(l)
	}

	_, sh := screen.Size()
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
	var possibleWords []int
	for i, w := range l.words {
		if w.Complete() && l.words[i] == l.currentWord {
			l.currentWord = nil
		} else if !w.Complete() && w.startedBy == 0 {
			possibleWords = append(possibleWords, i)
		}
	}

	if l.currentWord == nil && len(possibleWords) > 0 {
		l.currentWord = l.words[possibleWords[rand.Intn(len(possibleWords))]]
		l.currentWord.startedBy = PC
	}

	// End conditions
	if l.currentWord != nil {
		l.currentWordText.SetText("Current Word: " + l.currentWord.str[l.currentWord.completedChars:])
	}
	if gameWon {
		for _, w := range l.words {
			log.Printf("Words: %+v", w)
		}
		for _, i := range l.gt.items {
			log.Printf("Item: %+v", i)
		}

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
