package generate

import (
	"crypto/rand"
	"log"
	"math/big"
)

type Lotto []int64

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

	for i := 1; i <= 5; i++ {
		lotto = append(lotto, generateLotto())
	}

	return lotto
}

func generateLotto() Lotto {
	var lotto Lotto

	for len(lotto) < 6 {
		current, err := rand.Int(rand.Reader, big.NewInt(44))
		if err != nil {
			log.Fatal(err)
		}

		if !isExist(lotto, current.Int64()+1) {
			lotto = append(lotto, current.Int64()+1)
		}
	}

	return lotto
}

func isExist(stack []int64, current int64) bool {
	for i := 0; i < len(stack); i++ {
		if stack[i] == current {
			return true
		}
	}

	return false
}
