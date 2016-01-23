package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gophergala2016/gopher_typer"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	gt, err := gopher_typer.NewGopherTyper()
	if err != nil {
		log.Fatal(err)
	}
	gt.Run()

}
