package gopher_typer

import (
	"os"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type GopherTyper struct {
	g        *tl.Game
	wordList []string
	intro    introLevel
	game     gameLevel
	store    storeLevel
	end      endLevel
	console  tl.Text
	level    tl.Level
	stats    stats
	items    []item
}

func NewGopherTyper() (*GopherTyper, error) {
	wReader, err := os.Open("data/words.txt")
	if err != nil {
		return nil, err
	}

	gt := GopherTyper{}
	gt.g = tl.NewGame()
	gt.g.Screen().SetFps(30)
	gt.wordList = NewWordLoader(wReader)
	gt.intro = NewIntroLevel(&gt, tl.ColorBlack, tl.ColorBlue)
	gt.game = NewGameLevel(&gt, tl.ColorBlack, tl.ColorRed)
	gt.store = NewStoreLevel(&gt, tl.ColorBlack, tl.ColorCyan)
	gt.end = NewEndLevel(&gt, tl.ColorBlack, tl.ColorGreen)

	gt.stats = NewStats()

	return &gt, nil
}

func (gt *GopherTyper) Run() {
	gt.GoToIntro()
	gt.g.Start()
}

func (gt *GopherTyper) GoToIntro() {
	gt.level = &gt.intro
	gt.intro.Activate()
}

func (gt *GopherTyper) GoToGame() {
	if gt.stats.Lives == 0 {
		gt.stats = NewStats()
		gt.items = []item{}
	}
	gt.level = &gt.game
	gt.game.Activate()
}

func (gt *GopherTyper) GoToStore() {
	gt.level = &gt.store
	gt.store.Activate()
}

func (gt *GopherTyper) GoToEndWin() {
	gt.level = &gt.end
	gt.end.ActivateWin()
}

func (gt *GopherTyper) GoToEndFail() {
	gt.level = &gt.end
	gt.end.ActivateFail()
}

func (gt *GopherTyper) Tick(dt time.Duration) {
	if gt.level == nil {
		gt.GoToIntro()
	}
}
