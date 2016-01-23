package gopher_typer

import (
	"os"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type GopherTyper struct {
	g          *tl.Game
	wordList   []string
	ticker     *time.Ticker
	introLevel tl.Level
	gameLevel  tl.Level
	storeLevel tl.Level
	endLevel   tl.Level
}

func NewGopherTyper() (*GopherTyper, error) {
	wReader, err := os.Open("data/words.txt")
	if err != nil {
		return nil, err
	}

	gt := GopherTyper{}
	gt.g = tl.NewGame()
	gt.wordList = NewWordLoader(wReader)
	gt.introLevel = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlue, Fg: tl.ColorGreen})
	gt.gameLevel = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorRed, Fg: tl.ColorGreen})
	gt.storeLevel = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorYellow, Fg: tl.ColorGreen})
	gt.endLevel = tl.NewBaseLevel(tl.Cell{Bg: tl.ColorCyan, Fg: tl.ColorGreen})

	gt.GoToIntro()

	gt.ticker = time.NewTicker(33 * time.Millisecond)
	go func() {
		prevTick := time.Now()
		for t := range gt.ticker.C {
			gt.Tick(t.Sub(prevTick))
			prevTick = t
		}
	}()
	return &gt, nil
}

func (gt *GopherTyper) Run() {
	gt.g.Start()
}

var count = 0

func (gt *GopherTyper) GoToIntro() {
	gt.g.Screen().SetLevel(gt.introLevel)
}

func (gt *GopherTyper) GoToGame() {
	gt.g.Screen().SetLevel(gt.gameLevel)
}

func (gt *GopherTyper) GoToStore() {
	gt.g.Screen().SetLevel(gt.storeLevel)
}

func (gt *GopherTyper) GoToEnd() {
	gt.g.Screen().SetLevel(gt.endLevel)
}

func (gt *GopherTyper) Tick(dt time.Duration) {
	count++
	if count == 100 {
		gt.GoToGame()
	}
	if count == 200 {
		gt.GoToStore()
	}
	if count == 300 {
		gt.GoToEnd()
	}

}
