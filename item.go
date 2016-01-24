package gopher_typer

import "fmt"

type item interface {
	Name() string
	Desc() string
	Price() int
	PriceDesc() string

	Tick(gt *GopherTyper)
}

type goroutineItem struct {
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
func (i *goroutineItem) Tick(gt *GopherTyper) {
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
func (i *cpuUpgradeItem) Tick(gt *GopherTyper) {
}
