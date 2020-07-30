package main

import (
	"fmt"
)
// 題目:
// 求出前一百個自然數的平方和與和的平方之差
// 1平方 + 2平方 +......+ n平方 (平方和) = n(n+1)(2n+1)/6 
// (1+2+...+n)的平方 (合的平方) = (1+n)*n/2 
func GetDifference(num int) int {
	sumofsquares := num * (num + 1) * (2 * num + 1) / 6
	squaresofsum := ((1+num) * num / 2) * ((1+num) * num / 2)
	return squaresofsum - sumofsquares
}

func main() {
	sub := GetDifference(100)
	fmt.Println(sub)
}