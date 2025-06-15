package src

import (
	"math/big"
	"math/rand/v2"
)

const (
	numberAuthenticationRounds = 3
	successThreshold           = 2. / 3.
)

var (
	p, g, y, x, e, s uint64
	successRounds    int
)

func generateE() {
	e = uint64(rand.Uint32())
}

func checkX() bool {
	// log.Printf("g = %d, s = %d, p = %d\nFirst part: %d", g, s, p, moduloReduction(g, s, p))
	// log.Printf("y = %d, e = %d, p = %d\nSecond part: %d", y, e, p, moduloReduction(y, e, p))
	checkedX := moduloReduction(moduloReduction(g, s, p)*moduloReduction(y, e, p), 1, p)
	// log.Printf("X = %d, checked x = %d", x, checkedX)
	return x == checkedX
}

func sessionResult() bool {
	sessionThreshold := float64(successRounds) / float64(numberAuthenticationRounds)
	return sessionThreshold >= successThreshold
}

func moduloReduction[T int | uint32 | uint64 | int64 | byte](numberInt T, power T, module T) (remainder T) {
	result := big.NewInt(0)
	result.Exp(big.NewInt(int64(numberInt)), big.NewInt(int64(power)), big.NewInt(int64(module)))
	remainder = T(result.Int64())
	return
}
