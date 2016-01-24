package gopher_typer

import (
	"time"

	tl "github.com/JoelOtter/termloop"
)

type word struct {
	str                    string
	createdAt              time.Time
	v                      float64
	complete               bool
	completedChars         int
	deleteAt               time.Time
	x, y                   int
	fgComplete, fgTodo, bg tl.Attr
}

func NewWord(x, y int, val string, fg, bg tl.Attr) *word {
	return &word{str: val, createdAt: time.Now(), v: 2, x: x, y: y}
}

func (w *word) SetColor(fgComplete, fgTodo, bg tl.Attr) {
	w.fgComplete = fgComplete
	w.fgTodo = fgTodo
	w.bg = bg
}

func (w *word) Draw(s *tl.Screen) {
	for i, ch := range w.str {
		if i < w.completedChars {
			s.RenderCell(w.x+i, w.y, &tl.Cell{Fg: w.fgComplete, Bg: w.bg, Ch: ch})
		} else {
			s.RenderCell(w.x+i, w.y, &tl.Cell{Fg: w.fgTodo, Bg: w.bg, Ch: ch})
		}
	}
}

func (w *word) Tick(e tl.Event) {
}

func (w *word) Update() {
	w.y = int((time.Now().Sub(w.createdAt)).Seconds() * w.v)
	if w.complete {
		w.bg = tl.AttrUnderline
	}
}

func (w *word) KeyDown(ch rune) {
	found := false
	for i, r := range w.str {
		if i == w.completedChars && r == ch {
			w.completedChars++
			found = true
			break
		}
	}
	if !found {
		w.createdAt = w.createdAt.Add(-1 * time.Second)
	}
	if w.completedChars == len(w.str) {
		w.complete = true
	}
}
