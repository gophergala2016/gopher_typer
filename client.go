package gopher_typer

import (
	"os"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type GopherTyper struct {
	g        *tl.Game
	wordList []string
	ticker   *time.Ticker
	intro    introLevel
	game     gameLevel
	store    storeLevel
	end      endLevel
	console  tl.Text
	level    Level
	stats    stats
}

func NewGopherTyper() (*GopherTyper, error) {
	wReader, err := os.Open("data/words.txt")
	if err != nil {
		return nil, err
	}

	gt := GopherTyper{}
	gt.g = tl.NewGame()
	gt.wordList = NewWordLoader(wReader)
	gt.intro = NewIntroLevel(&gt, tl.ColorBlack, tl.ColorBlue)
	gt.game = NewGameLevel(&gt, tl.ColorBlack, tl.ColorRed)
	gt.store = NewStoreLevel(&gt, tl.ColorBlack, tl.ColorCyan)
	gt.end = NewEndLevel(&gt, tl.ColorBlack, tl.ColorGreen)

	gt.intro.AddEntity(&gt.console)
	gt.game.AddEntity(&gt.console)
	gt.store.AddEntity(&gt.console)
	gt.end.AddEntity(&gt.console)

	return &gt, nil
}

func (gt *GopherTyper) Run() {
	gt.ticker = time.NewTicker(33 * time.Millisecond)
	go func() {
		prevTick := time.Now()
		for t := range gt.ticker.C {
			gt.Tick(t.Sub(prevTick))
			prevTick = t
		}
	}()
	gt.g.Start()
}

var count = 0

func (gt *GopherTyper) GoToIntro() {
	gt.level = &gt.intro
	gt.intro.Activate()
}

func (gt *GopherTyper) GoToGame() {
	gt.level = &gt.game
	gt.game.Activate()
}

func (gt *GopherTyper) GoToStore() {
	gt.level = &gt.store
	gt.store.Activate()
}

func (gt *GopherTyper) GoToEnd() {
	gt.level = &gt.end
	gt.end.Activate()
}

func (gt *GopherTyper) Tick(dt time.Duration) {
	if gt.level == nil {
		gt.GoToIntro()
	}
	gt.level.Update(dt)
	if count == 0 {
		gt.GoToIntro()
	}
	count++
}
