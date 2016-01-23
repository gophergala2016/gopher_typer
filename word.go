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
	x, y           int
}

func NewWord(x, y int, val string, fg, bg tl.Attr) *word {
	return &word{tl.NewText(x, y, val, fg, bg), val, time.Now(), 2, false, 0, time.Time{}, x, y}
}

func (w *word) Update() {
	w.y = int((time.Now().Sub(w.createdAt)).Seconds() * w.v)
	w.SetPosition(w.x, int(w.y))
	if w.complete {
		fg, _ := w.Color()
		w.SetColor(fg, tl.AttrUnderline)
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
