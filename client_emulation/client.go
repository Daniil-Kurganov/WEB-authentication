package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"math/rand/v2"
	"os"
	"slices"
)

const numberTests = 3

var p, q, g, w, y, r uint32

func generateKeys() {
	for {
		p = getSimpleUint32()
		currentFactors := primaryFactorization(p - 1)
		slices.Reverse(currentFactors)
		for _, currentFactor := range currentFactors {
			if currentFactor != 0 {
				if ok, _ := millerRabinTest(currentFactor, numberTests); ok {
					q = currentFactor
					break
				}
			}
		}
		if q != 0 {
			break
		}
	}
	for {
		g = rand.Uint32N(math.MaxUint32)
		if extendedEuclideanAlgorithm(uint64(g), uint64(p)) != 1 {
			continue
		}
		g = moduloReduction(g, (p-1)/q, p)
		if g != 1 {
			break
		}
	}
	w = rand.Uint32N(q)
	y = moduloReduction(g, q-w, p)
}

func moduloReduction[T int | uint32 | uint64 | int64 | byte](numberInt T, power T, module T) (remainder T) {
	result := big.NewInt(0)
	result.Exp(big.NewInt(int64(numberInt)), big.NewInt(int64(power)), big.NewInt(int64(module)))
	remainder = T(result.Int64())
	return
}

func extendedEuclideanAlgorithm(a, b uint64) uint64 {
	remainders := []uint64{a, b}
	xs := []uint64{1, 0}
	ys := []uint64{0, 1}
	counterIterations := 1
	for {
		currentQuotines := remainders[counterIterations-1] / remainders[counterIterations]
		nextRemainder := remainders[counterIterations-1] % remainders[counterIterations]
		if nextRemainder == 0 {
			return remainders[counterIterations]
		} else {
			remainders = append(remainders, nextRemainder)
			xs = append(xs, xs[counterIterations-1]-currentQuotines*xs[counterIterations])
			ys = append(ys, ys[counterIterations-1]-currentQuotines*ys[counterIterations])
			counterIterations += 1
		}
	}
}

func millerRabinTest(checkedNumber, countOfWitness uint32) (numberIsPrimary bool, err error) {
	if checkedNumber < 5 {
		err = fmt.Errorf("invalid checked number: must be equal or bigger then 5, but got %d", checkedNumber)
		return
	}
	if checkedNumber%2 == 0 {
		err = fmt.Errorf("invalid checked number: must be odd, but got %d", checkedNumber)
		return
	}
	lessBy1 := uint64(checkedNumber - 1)
	var s int
	t := lessBy1
	for {
		if t%2 != 0 && uint64(math.Pow(2, float64(s)))*t == lessBy1 {
			// log.Printf("s = %d, t = %d", s, t)
			break
		}
		s += 1
		t /= 2
	}
	var witnesses []uint64
	for counterWitness := uint32(0); counterWitness < countOfWitness; counterWitness++ {
		witnesses = append(witnesses, rand.Uint64N(uint64(checkedNumber-4))+2)
	}
	for _, currentWitness := range witnesses {
		// log.Printf("Current witness: %d", currentWitness)
		currentDCM := extendedEuclideanAlgorithm(currentWitness, uint64(checkedNumber))
		// log.Printf(" DCM(%d, %d) = %d", currentWitness, checkedNumber, currentDCM)
		if currentDCM != 1 {
			// log.Printf(" DCM(a, n) = DCM(%d, %d) = %d ≠ 1 -> exit", currentWitness, checkedNumber, currentDCM)
			return
		}
		currentB := moduloReduction(currentWitness, t, uint64(checkedNumber))
		// log.Printf(" b = (%d ^ %d) mod(%d) = %d", currentWitness, t, checkedNumber, currentB)
		if currentB == 1 || currentB == lessBy1 {
			// log.Printf(" %d - probably simple", checkedNumber)
			continue
		}
		var counterK int
		for {
			if counterK >= s {
				// log.Printf(" %d (k) ≥ %d (s) -> exit", counterK, s)
				return
			}
			currentB = moduloReduction(currentB, 2, uint64(checkedNumber))
			// log.Printf(" b = (%d ^ 2) mod(%d) = %d", currentB, checkedNumber, currentB)
			counterK += 1
			if currentB == lessBy1 {
				// log.Printf(" %d - probably simple", checkedNumber)
				break
			}
		}
	}
	numberIsPrimary = true
	return
}

func getSimpleUint32(rightBorder ...uint32) (result uint32) {
	if len(rightBorder) == 0 {
		rightBorder = append(rightBorder, math.MaxInt32)
	}
	for {
		result = rand.Uint32N(rightBorder[0])
		if result == 0 {
			continue
		}
		if ok, _ := millerRabinTest(result, numberTests); ok {
			return
		}
	}
}

func primaryFactorization(number uint32) (factors []uint32) {
	for number%2 == 0 {
		factors = append(factors, 2)
		number = number / 2
	}
	for currentNumber := uint32(3); currentNumber*currentNumber <= number; currentNumber = currentNumber + 2 {
		for number%currentNumber == 0 {
			factors = append(factors, currentNumber)
			number = number / currentNumber
		}
	}
	if number > 2 {
		factors = append(factors, number)
	}
	return
}

func generateX() uint32 {
	r = rand.Uint32N(q)
	return moduloReduction(g, r, p)
}

func generateS(e uint32) uint64 {
	return moduloReduction(uint64(r)+moduloReduction(uint64(w)*uint64(e), 1, uint64(q)), 1, uint64(q))
}

func main() {
	log.SetFlags(0)
	generateKeys()
	log.Printf("Ключи:\n открытые:\n  P: %v\n  Q: %v\n  G: %v\n  Y: %v\n W (закрытый): %v\n\n", p, q, g, y, w)
	for {
		var choise int
		fmt.Print("\n\nВыберите действие:\n 0 - выход\n 1 - сгенерировать параметры раунда\n\nВведите цифру действия: ")
		fmt.Fscan(os.Stdin, &choise)
		switch choise {
		case 0:
			log.Print("\nЗавершение работы программы")
			return
		case 1:
			log.Printf("\nПараметр R: %d\nПараметр X: %d", r, generateX())
			var e uint32
			fmt.Print("\nВведите параметр E: ")
			fmt.Fscan(os.Stdin, &e)
			log.Printf("Параметр S: %d", generateS(e))
		}
	}
}
