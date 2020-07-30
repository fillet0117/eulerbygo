package main

import (
	"fmt"
)

// 題目:
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

func multiples(num int) int {
	sum := 0
	for i := 0; i < num; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			sum += i
		} else if i % 3 == 0 {
			sum += i
		} else if i % 5 == 0 {
			sum += i
		}
	}
	return sum
}
// 另一種解
// 先算3的倍數的總和:
// 3+6+9+.....+999 = 3*(1+2+3+...+333)
// 5的倍數總和:
// 5+10+15+...+995 = 5*(1+2+3+...+199)
// 上面兩個相加在減掉15的倍數:
// 15+30+45+...+(999/15上限) = 15*(1+2+3+...+66)

func SumDivisibleBy(num int, below int) int {
	p := below / num
	return num * (p * (1 + p) / 2)
}

func main() {
	num := 1000
	sum := multiples(num)
	fmt.Println(sum)

	sum = SumDivisibleBy(3, num - 1) + SumDivisibleBy(5, num - 1) - SumDivisibleBy(15, num - 1)
	fmt.Println(sum)
}