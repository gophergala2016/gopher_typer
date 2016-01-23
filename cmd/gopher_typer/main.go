package main

import (
	"log"

	"github.com/gophergala2016/gopher_typer"
)

func main() {
	gt, err := gopher_typer.NewGopherTyper()
	if err != nil {
		log.Fatal(err)
	}
	gt.Run()

}
