package gopherTyper

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
	SetID(int)
	Reset(gt *GopherTyper)
	Purchase(g *storeLevel) bool
	Dupe() item

	Tick(g *gameLevel)
}

type goroutineItem struct {
	wakeAt      time.Time
	baseWait    time.Duration
	waitRange   time.Duration
	currentWord *word
	id          int
	cpuUpgrades int
	price       int
}

func (i *goroutineItem) Name() string {
	return "Goroutine"
}
func (i *goroutineItem) Desc() string {
	return "Add a goroutine to help type words for you"
}
func (i *goroutineItem) Price() int {
	return i.price
}
func (i *goroutineItem) PriceDesc() string {
	return fmt.Sprintf("$%d", i.Price())
}
func (i *goroutineItem) Tick(gl *gameLevel) {
	if time.Now().After(i.wakeAt) {
		if i.currentWord == nil {
			var possibleWords []int
			for i, w := range gl.words {
				if gl.currentWord != gl.words[i] && !w.Complete() && w.startedBy == 0 {
					possibleWords = append(possibleWords, i)
				}
			}
			if len(possibleWords) > 0 {
				i.currentWord = gl.words[possibleWords[rand.Intn(len(possibleWords))]]
				i.currentWord.completedChars++
				i.currentWord.startedBy = i.id
			}
		} else {
			i.currentWord.completedChars++
			gl.gt.stats.Garbage++
			if i.currentWord.Complete() {
				i.currentWord = nil
			}
		}

		i.sleep()
	}
}

func (i *goroutineItem) sleep() {
	i.wakeAt = time.Now().Add(i.baseWait/time.Duration(i.cpuUpgrades) + time.Duration(rand.Intn(int(i.waitRange))))
}

func (i *goroutineItem) SetID(id int) {
	i.id = id
}
func (i *goroutineItem) Reset(gt *GopherTyper) {
	i.currentWord = nil
	i.cpuUpgrades = gt.stats.CPUUpgrades
	i.price = 1000
	for _, itm := range gt.items {
		if itm.Name() == i.Name() {
			i.price *= 2
		}
	}
}
func (i *goroutineItem) Dupe() item {
	var dupe goroutineItem
	dupe = *i
	return &dupe
}
func (i *goroutineItem) Purchase(l *storeLevel) bool {
	return true
}

func newGoroutineItem(waitRange, baseWait time.Duration) *goroutineItem {
	item := goroutineItem{waitRange: waitRange, baseWait: baseWait, cpuUpgrades: 1}
	item.sleep()
	return &item
}

type cpuUpgradeItem struct {
	id    int
	price int
}

func (i *cpuUpgradeItem) Name() string {
	return "CPU Upgrade"
}
func (i *cpuUpgradeItem) Desc() string {
	return "Make your goroutines go faster"
}
func (i *cpuUpgradeItem) Price() int {
	return i.price
}
func (i *cpuUpgradeItem) PriceDesc() string {
	return fmt.Sprintf("$%d", i.Price())
}
func (i *cpuUpgradeItem) Tick(gl *gameLevel) {
}
func (i *cpuUpgradeItem) SetID(id int) {
	i.id = id
}
func (i *cpuUpgradeItem) Reset(gt *GopherTyper) {
	i.price = 2000 * gt.stats.CPUUpgrades
}
func (i *cpuUpgradeItem) Purchase(l *storeLevel) bool {
	l.gt.stats.CPUUpgrades++
	return false
}
func (i *cpuUpgradeItem) Dupe() item {
	var dupe cpuUpgradeItem
	dupe = *i
	return &dupe

}

type goUpgradeItem struct {
	id    int
	price int
}

func (i *goUpgradeItem) Name() string {
	return "Go Upgrade"
}
func (i *goUpgradeItem) Desc() string {
	return "Improves garbage collection performance"
}
func (i *goUpgradeItem) Price() int {
	return i.price
}
func (i *goUpgradeItem) PriceDesc() string {
	return fmt.Sprintf("$%d", i.Price())
}
func (i *goUpgradeItem) Tick(gl *gameLevel) {
}
func (i *goUpgradeItem) SetID(id int) {
	i.id = id
}
func (i *goUpgradeItem) Reset(gt *GopherTyper) {
	i.price = int(1000 * gt.stats.GoVersion)
}
func (i *goUpgradeItem) Purchase(l *storeLevel) bool {
	l.gt.stats.GoVersion += 0.1
	l.gt.stats.GarbageFreq += 3
	return false
}
func (i *goUpgradeItem) Dupe() item {
	var dupe goUpgradeItem
	dupe = *i
	return &dupe

}
