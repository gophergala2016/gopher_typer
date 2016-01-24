package gopher_typer

import (
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
	_, sh := screen.Size()
	gameOver := false
	for _, w := range l.words {
		w.Update()
		if !w.complete && w.y > sh-3 {
			gameOver = true
		}
	}
	var possibleWords []int
	for i, w := range l.words {
		if !w.complete {
			possibleWords = append(possibleWords, i)
		} else if l.words[i] == l.currentWord {
			l.currentWord = nil
		}
	}

	if l.currentWord == nil && len(possibleWords) > 0 {
		l.currentWord = l.words[possibleWords[rand.Intn(len(possibleWords))]]
	}

	// End conditions
	if l.currentWord != nil {
		l.currentWordText.SetText("Current Word: " + l.currentWord.str[l.currentWord.completedChars:])
		l.currentWord.SetColor(tl.ColorGreen, tl.ColorBlue)
	} else {
		l.gt.GoToEndWin()
	}
	if gameOver {
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
