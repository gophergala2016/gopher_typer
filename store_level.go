package gopher_typer

import (
	"fmt"
	"io/ioutil"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type storeLevel struct {
	tl.Level
	gt *GopherTyper
	bg tl.Attr
	fg tl.Attr

	items       []item
	currentItem int
}

func (l *storeLevel) refresh() {
	l.Level = tl.NewBaseLevel(tl.Cell{Bg: l.bg, Fg: l.fg})
	l.gt.store.AddEntity(&l.gt.console)
	l.gt.console.SetText("")

	w, h := l.gt.g.Screen().Size()
	rect := tl.NewRectangle(10, 2, w-20, h-4, tl.ColorGreen)
	l.AddEntity(rect)

	store, _ := ioutil.ReadFile("data/store.txt")
	c := tl.CanvasFromString(string(store))
	l.AddEntity(tl.NewEntityFromCanvas(w/2-len(c)/2, 4, c))

	msg := "Up/Down(j/k), Enter to purchase, N to return to the game"
	l.AddEntity(tl.NewText(w/2-len(msg)/2, 10, msg, tl.ColorBlack, tl.ColorDefault))

	msg = fmt.Sprintf("Cash: $%d", l.gt.stats.Dollars)
	l.AddEntity(tl.NewText(14, 11, msg, tl.ColorBlack, tl.ColorDefault))

	y := 12
	for idx, i := range l.items {
		i.Reset(l.gt)
		x := 14
		fg := tl.ColorBlack
		if i.Price() > l.gt.stats.Dollars {
			fg = tl.ColorRed
		}
		var price string
		if l.currentItem == idx {
			price = ">" + i.PriceDesc() + "<"
		} else {
			price = " " + i.PriceDesc()
		}
		l.AddEntity(tl.NewText(x, y, price, fg, tl.ColorDefault))
		x += len(i.PriceDesc()) + 4
		l.AddEntity(tl.NewText(x, y, i.Name(), tl.ColorBlue, tl.ColorDefault))
		y++
	}

	desc := l.items[l.currentItem].Desc()
	l.AddEntity(tl.NewText(14, y+1, desc, tl.ColorBlue, tl.ColorDefault))

	y = 12
	x := w - 30
	msg = fmt.Sprintf("Goroutines: %d", len(l.gt.items))
	l.AddEntity(tl.NewText(x, y, msg, tl.ColorBlue, tl.ColorDefault))
	y++
	msg = fmt.Sprintf("CPU Upgrades: %d", l.gt.stats.CpuUpgrades)
	l.AddEntity(tl.NewText(x, y, msg, tl.ColorBlue, tl.ColorDefault))
	y++
	msg = fmt.Sprintf("Go Version: %0.1f", l.gt.stats.GoVersion)
	l.AddEntity(tl.NewText(x, y, msg, tl.ColorBlue, tl.ColorDefault))
	y++

	l.gt.g.Screen().SetLevel(l)
}

func (l *storeLevel) Activate() {
	l.currentItem = 0
	l.refresh()
}

func (l *storeLevel) purchaseItem(id int) {
	itm := l.items[id]
	if itm.Price() <= l.gt.stats.Dollars {
		if itm.Purchase(l) {
			l.gt.items = append(l.gt.items, itm.Dupe())

			l.gt.items[len(l.gt.items)-1].SetId(len(l.gt.items))
		}
		l.gt.stats.Dollars -= itm.Price()
	}
}

func (l *storeLevel) Tick(e tl.Event) {
	if e.Type == tl.EventKey {
		if e.Key == tl.KeyArrowDown || e.Ch == 'j' {
			l.currentItem = (l.currentItem + 1) % len(l.items)
		} else if e.Key == tl.KeyArrowUp || e.Ch == 'k' {
			l.currentItem = (l.currentItem - 1)
			if l.currentItem < 0 {
				l.currentItem = len(l.items) - 1
			}
		} else if e.Key == tl.KeyEnter || e.Ch == 'e' {
			l.purchaseItem(l.currentItem)
		} else if e.Ch == 'N' || e.Ch == 'n' {
			l.gt.GoToGame()
			return
		}
		l.refresh()
	}
}

func NewBaseItems() []item {
	return []item{
		NewGoroutineItem(150*time.Millisecond, 500*time.Millisecond),
		&cpuUpgradeItem{},
		&goUpgradeItem{},
	}
}

func NewStoreLevel(g *GopherTyper, fg, bg tl.Attr) storeLevel {
	return storeLevel{gt: g, bg: bg, fg: fg, items: NewBaseItems()}
}
