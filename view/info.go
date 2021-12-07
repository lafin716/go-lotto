package view

import (
	"fmt"
	"github.com/lafin716/lotto/generate"
	"strconv"
	"strings"
)

func Greeting() {
	greet := delayString{text: "GM : 환영합니다... 로또를 몇개 만드시겠습니까? : "}
	greet.show(50)
}

func AskLotto() int {
	var lottoCount int
	fmt.Scanln(&lottoCount)
	loading := delayString{text: strconv.Itoa(lottoCount) + "개 생성 중...\n"}
	loading.show(50)

	count(3)

	return lottoCount
}

func Print(lottoList []generate.LottoPack) {
	for i, lotto := range lottoList {
		fmt.Printf("====== %d 회차 ======\n", i + 1)
		for _, num := range lotto {
			fmt.Printf("%d\n", num)
		}
		fmt.Printf("===================\n")
	}
}

func count(c int) {
	var s []string
	for i := c; i > 0; i-- {
		s = append(s, strconv.Itoa(i))
	}

	count := delayString{text: strings.Join(s, "")}
	count.showLine(1000)
}