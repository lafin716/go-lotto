package view

import (
	"fmt"
	"github.com/lafin716/lotto/generate"
	"strconv"
	"strings"
)

const DelayTime = 50

func Greeting() {
	greet := delayString{text: "@@@@ 대 박 기 원 @@@@ \n환영합니다!\n"}
	greet.show(DelayTime)
}

func printNotNumber() {
	retry := delayString{text: "0 이상의 숫자를 입력하세요.\n"}
	retry.show(DelayTime)
}

func getCount(lottoCount *int) bool {
	ask := delayString{"로또를 몇개 만드시겠습니까? : "}
	ask.show(DelayTime)

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		printNotNumber()
		return false
	}

	num, err := strconv.Atoi(input)
	if err != nil || num < 1 {
		printNotNumber()
		return false
	}

	*lottoCount = num
	return true
}

func AskLotto() int {
	var lottoCount int
	for !getCount(&lottoCount) {
	}
	loading := delayString{text: strconv.Itoa(lottoCount) + "개 생성 중...\n"}
	loading.show(DelayTime)

	count(3)

	return lottoCount
}

func Print(lottoList []generate.LottoPack) {
	for i, lotto := range lottoList {
		fmt.Printf("====== %d 회차 ======\n", i+1)
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
