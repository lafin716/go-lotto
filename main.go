package main

import (
	"github.com/lafin716/lotto/generate"
	"github.com/lafin716/lotto/view"
)

func main() {
	regame := true

	for regame {
		view.Greeting()
		lottoCount := view.AskLotto()
		lottoList := generate.GetLottoList(lottoCount)
		view.Print(lottoList)

		regame = view.ReGame()
	}

	view.GoodBye()
}
