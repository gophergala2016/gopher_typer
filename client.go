package gopherTyper

import (
	"os"

	tl "github.com/JoelOtter/termloop"
)

// GopherTyper handles the local state of the game.
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

// NewGopherTyper gets the game ready to run.
func NewGopherTyper() (*GopherTyper, error) {
	wReader, err := os.Open("data/words.txt")
	if err != nil {
		return nil, err
	}

	gt := GopherTyper{}
	gt.g = tl.NewGame()
	gt.g.Screen().SetFps(30)
	gt.wordList = newWordLoader(wReader)
	gt.intro = newIntroLevel(&gt, tl.ColorBlack, tl.ColorBlue)
	gt.game = newGameLevel(&gt, tl.ColorBlack, tl.ColorRed)
	gt.store = newStoreLevel(&gt, tl.ColorBlack, tl.ColorCyan)
	gt.end = newEndLevel(&gt, tl.ColorBlack, tl.ColorGreen)

	gt.stats = newStats()

	return &gt, nil
}

// Run starts the game, and blocks forever.
func (gt *GopherTyper) Run() {
	gt.goToIntro()
	gt.g.Start()
}

func (gt *GopherTyper) goToIntro() {
	gt.level = &gt.intro
	gt.intro.Activate()
}

func (gt *GopherTyper) goToGame() {
	if gt.stats.Lives == 0 {
		gt.stats = newStats()
		gt.items = []item{}
	}
	gt.level = &gt.game
	gt.game.Activate()
}

func (gt *GopherTyper) goToStore() {
	gt.level = &gt.store
	gt.store.Activate()
}

func (gt *GopherTyper) goToEndWin() {
	gt.level = &gt.end
	gt.end.ActivateWin()
}

func (gt *GopherTyper) goToEndFail() {
	gt.level = &gt.end
	gt.end.ActivateFail()
}
