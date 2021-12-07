package view

import (
	"fmt"
	"time"
)

type delayString struct {
	text string
}

func execute(s rune, duration time.Duration) {
	fmt.Printf("%c", s)
	time.Sleep(duration)
}

func executeLine(s rune, duration time.Duration) {
	fmt.Printf("%c\n", s)
	time.Sleep(duration)
}

func (ds delayString) show(duration int64) {
	for _, s := range ds.text {
		execute(s, time.Duration(duration) * time.Millisecond)
	}
}

func (ds delayString) showLine(duration int64) {
	for _, s := range ds.text {
		executeLine(s, time.Duration(duration) * time.Millisecond)
	}
}
