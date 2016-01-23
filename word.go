package gopher_typer

import (
	"time"

	tl "github.com/JoelOtter/termloop"
)

type word struct {
	*tl.Text
	str            string
	createdAt      time.Time
	v              float64
	complete       bool
	completedChars int
	deleteAt       time.Time
}

func NewWord(x, y int, val string, fg, bg tl.Attr) *word {
	return &word{tl.NewText(x, y, val, fg, bg), val, time.Now(), 2, false, 0, time.Time{}}
}

func (w *word) Update(dt time.Duration) {
	x, _ := w.Position()
	y := (time.Now().Sub(w.createdAt)).Seconds() * w.v
	w.SetPosition(x, int(y))
	if w.complete {
		fg, _ := w.Color()
		w.SetColor(fg, tl.AttrUnderline)
	}
}

func (w *word) KeyDown(ch rune) {
	for i, r := range w.str {
		if i == w.completedChars && r == ch {
			w.completedChars++
			break
		}
	}
	if w.completedChars == len(w.str) {
		w.complete = true
	}
}
