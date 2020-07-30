package main

import (
	"fmt"
	"math"
)

// 題目:
// 找第10001個質數是多少

func isPrime(num int) bool {
	if num == 1 {
		return false
	} else if num < 4 {
		return true
	} else if num % 2 == 0 {
		return false
	} else if num < 9 {
		return true
	} else if num % 3 == 0 {
		return false
	} else{
		nSqrt := int(math.Floor(math.Sqrt(float64(num))))
		for i := 3; i <= nSqrt; i++ {
			if num % i == 0 {
				return false
			}
		}
		return true
	}
}

func GetPrime(num int) int {
	count := 0
	i := 1
	for {
		if isPrime(i) {
			count++
		}
		if count == num {
			return i
		}
		i++
	}
}

func main() {
	prime := GetPrime(10001)
	fmt.Println(prime)
}