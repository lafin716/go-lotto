package view

import (
	"fmt"
	"github.com/lafin716/lotto/generate"
	"strconv"
	"strings"
)

func Greeting() {
	showMessage("@@@@ 대 박 기 원 @@@@ \n환영합니다!\n")
}

func notNumber() {
	showMessage("0 이상의 숫자를 입력하세요.\n")
}

func notChoice() {
	showMessage("예, 아니오로만 입력하세요.\n")
}

func getCount(lottoCount *int) bool {
	showMessage("로또를 몇개 만드시겠습니까? : ")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		notNumber()
		return false
	}

	num, err := strconv.Atoi(input)
	if err != nil || num < 1 {
		notNumber()
		return false
	}

	*lottoCount = num
	return true
}

func AskLotto() int {
	var lottoCount int
	for !getCount(&lottoCount) {
	}
	showMessage(strconv.Itoa(lottoCount) + "개 생성 중...\n")
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
