package gopher_typer

import (
	"fmt"
	"math/rand"
	"time"
)

type item interface {
	Name() string
	Desc() string
	Price() int
	PriceDesc() string

	Tick(g *gameLevel)
}

type goroutineItem struct {
	wakeAt    time.Time
	baseWait  time.Duration
	waitRange time.Duration
}

func (i *goroutineItem) Name() string {
	return "Goroutine"
}
func (i *goroutineItem) Desc() string {
	return "Add a goroutine to help type words for you"
}
func (i *goroutineItem) Price() int {
	return 1000
}
func (i *goroutineItem) PriceDesc() string {
	return fmt.Sprintf("$%d", i.Price())
}
func (i *goroutineItem) Tick(gl *gameLevel) {
	if time.Now().After(i.wakeAt) {
		//eat a word
		var possibleWords []int
		for i, w := range gl.words {
			if gl.currentWord != gl.words[i] && !w.complete {
				possibleWords = append(possibleWords, i)
			}
		}
		if len(possibleWords) > 0 {
			gl.words[possibleWords[rand.Intn(len(possibleWords))]].complete = true
		}

		i.sleep()
	}
}

func (i *goroutineItem) sleep() {
	i.wakeAt = time.Now().Add(i.baseWait + time.Duration(rand.Intn(int(i.waitRange))))
}

func NewGoroutineItem(waitRange, baseWait time.Duration) *goroutineItem {
	item := goroutineItem{waitRange: waitRange, baseWait: baseWait}
	item.sleep()
	return &item
}

type cpuUpgradeItem struct {
}

func (i *cpuUpgradeItem) Name() string {
	return "CPU Upgrade"
}
func (i *cpuUpgradeItem) Desc() string {
	return "Make your goroutines go faster"
}
func (i *cpuUpgradeItem) Price() int {
	return 2000
}
func (i *cpuUpgradeItem) PriceDesc() string {
	return fmt.Sprintf("$%d", i.Price())
}
func (i *cpuUpgradeItem) Tick(gl *gameLevel) {
}
