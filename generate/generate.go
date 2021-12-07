package generate

import (
	"math/rand"
	"time"
)

type Lotto []int

type LottoPack []Lotto

func GetLottoList(count int) []LottoPack {
	var lottoPack []LottoPack

	for i := 0; i < count; i++ {
		lottoPack = append(lottoPack, GetLotto())
	}

	return lottoPack
}

func GetLotto() []Lotto {
	var lotto []Lotto

	for i:= 1; i <= 5; i++ {
		lotto = append(lotto, generateLotto())
	}

	return lotto
}

func generateLotto() Lotto {
	var lotto Lotto

	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)

	for len(lotto) < 6 {
		current := random.Intn(44) + 1
		if !isExist(lotto, current) {
			lotto = append(lotto, current)
		}
	}

	return lotto
}

func isExist(stack []int, current int) bool {
	for i := 0; i < len(stack); i++ {
		if stack[i] == current {
			return true
		}
	}

	return false
}
