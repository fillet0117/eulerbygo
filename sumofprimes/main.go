package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num == 1 {
		return false
	} else if num != 2 && num % 2 == 0 {
		return false
	} else if num == 2 || num == 3 {
		return true
	} else if num != 3 && num % 3 == 0 {
		return false
	} else {
		nSqrt :=  int(math.Floor(math.Sqrt(float64(num))))
		for i := 3; i <= nSqrt; i++ {
			if num % i == 0 {
				return false
			}
		}
		return true
	}
}

func sumOfPrime(num int) int {
	sum := 0
	for i := 1; i < num; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

func main() {
	sum := sumOfPrime(2000000)
	fmt.Println(sum)
}