package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(simpleInterest(798, 6.28, 1))
	fmt.Println(simpleInterest(234, 9, 3))
	fmt.Println(simpleInterest(8000, 3, 5))

}

func simpleInterest(principle, rate, time float64) float64 {
	interest := (principle * rate * time) / 100
	return math.Round(interest*100) / 100
}
