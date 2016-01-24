package gopher_typer

import (
	"time"

	tl "github.com/JoelOtter/termloop"
)

type word struct {
	str                   string
	createdAt             time.Time
	v                     float64
	startedBy             int
	completedChars        int
	deleteAt              time.Time
	x, y, baseY           int
	fgComplete, fgTodo    tl.Attr
	bgPlayer, bgGoroutine tl.Attr
}

const PC = -1

func NewWord(x, y int, val string, fgComplete, fgTodo, bgPlayer, bgGoroutine tl.Attr) *word {
	return &word{str: val, createdAt: time.Now(), v: 2, x: x, y: y, baseY: y, fgComplete: fgComplete, fgTodo: fgTodo, bgPlayer: bgPlayer, bgGoroutine: bgGoroutine}
}

func (w *word) Draw(s *tl.Screen) {
	for i, ch := range w.str {
		if w.startedBy == 0 {
			s.RenderCell(w.x+i, w.y, &tl.Cell{Fg: w.fgTodo, Bg: tl.ColorDefault, Ch: ch})
		} else {
			var bg tl.Attr
			if w.startedBy == PC {
				bg = w.bgPlayer
			} else {
				bg = w.bgGoroutine
			}
			if i < w.completedChars {
				s.RenderCell(w.x+i, w.y, &tl.Cell{Fg: w.fgComplete, Bg: bg, Ch: ch})
			} else {
				s.RenderCell(w.x+i, w.y, &tl.Cell{Fg: w.fgTodo, Bg: bg, Ch: ch})
			}
		}
	}
}

func (w *word) Tick(e tl.Event) {
}

func (w *word) Complete() bool {
	return w.completedChars == len(w.str)
}

func (w *word) Update() {
	w.y = w.baseY + int((time.Now().Sub(w.createdAt)).Seconds()*w.v)

	if w.Complete() {
		w.bgPlayer = tl.AttrUnderline
		w.bgGoroutine = tl.AttrUnderline
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
}
