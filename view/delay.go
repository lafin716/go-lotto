package view

import (
	"fmt"
	"time"
)

type delayString struct {
	text string
}

const DEFAULT_DELAY = 35

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
		execute(s, time.Duration(duration)*time.Millisecond)
	}
}

func (ds delayString) showLine(duration int64) {
	for _, s := range ds.text {
		executeLine(s, time.Duration(duration)*time.Millisecond)
	}
}

func showMessage(text string) {
	d := delayString{text: text}
	d.show(DEFAULT_DELAY)
}

func showMessageWithTime(text string, duration int64) {
	d := delayString{text: text}
	d.show(duration)
}
