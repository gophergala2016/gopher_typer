package gopher_typer

import "time"

type Level interface {
	Update(time.Duration)
}
