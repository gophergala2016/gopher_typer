package gopher_typer

import (
	"math/rand"
	"strings"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type gameLevel struct {
	tl.Level
	gt              *GopherTyper
	words           []*word
	currentWord     *word
	currentWordText *tl.Text
}

func (l *gameLevel) Activate() {
	l.gt.console.SetText("Game Level")

	numWords := 10
	w, h := l.gt.g.Screen().Size()

	for i := 0; i < numWords; i++ {
		w := NewWord(w/numWords*i, 0, l.gt.wordList[rand.Intn(len(l.gt.wordList))], tl.ColorGreen, tl.ColorDefault)
		l.AddEntity(w)
		l.words = append(l.words, w)
	}
	l.currentWordText = tl.NewText(0, h-1, "", tl.ColorRed, tl.ColorBlue)
	l.AddEntity(l.currentWordText)

	l.AddEntity(tl.NewText(0, h-2, strings.Repeat("*", w), tl.ColorBlack, tl.ColorDefault))
	l.gt.g.Screen().SetLevel(l)
}

func (l *gameLevel) Draw(screen *tl.Screen) {
	l.Level.Draw(screen)
	for _, w := range l.words {
		w.Update()
	}
	l.currentWord = nil
	for i, w := range l.words {
		if !w.complete {
			l.currentWord = l.words[i]
			break
		}
	}
	if l.currentWord != nil {
		l.currentWordText.SetText("Current Word: " + l.currentWord.str[l.currentWord.completedChars:])
	} else {
		l.gt.GoToEnd()
	}
}

func (l *gameLevel) Tick(e tl.Event) {
	if e.Type == tl.EventKey {
		if l.currentWord != nil {
			l.currentWord.KeyDown(e.Ch)
		}
	}
}
func (l *gameLevel) Update(dt time.Duration) {
}

func NewGameLevel(g *GopherTyper, fg, bg tl.Attr) gameLevel {
	l := tl.NewBaseLevel(tl.Cell{Bg: bg, Fg: fg})
	return gameLevel{l, g, []*word{}, nil, nil}
}
