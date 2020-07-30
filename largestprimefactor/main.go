package main

import (
	"fmt"
	"math"
)
// 題目:
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

// math.Sqrt 開根號
// math.Floor 取floor
func LargestPrime(num int) int {
	n := int(math.Floor(math.Sqrt(float64(num))))
	largest := 0
	for i := 3; i <= n; i++ {
		check := getPrime(i)
		if check && num % i == 0 {
			largest = i
		}
	}
	if largest == 0 {
		return num
	}
	return largest
}

func getPrime(num int) bool {
	if num % 2 == 0 {
		return false
	} else if num > 2 {
		nSqrt := int(math.Floor(math.Sqrt(float64(num))))
		for i := 3; i <= nSqrt; i++ {
			if num % i == 0 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func main() {
	largest := LargestPrime(600851475143)
	fmt.Println(largest)
}