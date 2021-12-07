package main

import (
	"github.com/lafin716/lotto/generate"
	"github.com/lafin716/lotto/view"
)


func main() {
	view.Greeting()
	lottoCount := view.AskLotto()
	lottoList := generate.GetLottoList(lottoCount)
	view.Print(lottoList)
}
