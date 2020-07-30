package main

import(
	"fmt"
	"strconv"
	"strings"
)

// 題目:
// 回文數在兩個方向上都相同。 由兩個兩位數的乘積組成的最大回文數為9009 = 91×99。
// 查找由兩個3位數字的乘積組成的最大回文。


// 如果數字為最多六位數的回文數P
// P = 100000x + 10000y + 1000z + 100z + 10y + x
//   = 100001x + 10010y + 1100z
//   = 11 (9091x + 910y + 100z)
// P 一定為11的倍數
func getLargest() int {
	largest := 0
	eleven := 0
	sub := 0
	for i := 999; i >= 100; i-- {
		if i % 11 == 0 {
			eleven = 999
			sub = 1
		} else {
			eleven = 990
			sub = 11
		}
		for j := eleven; j >= 100; j = j-sub {
			mul := i * j
			if FindPalindromicNumber(mul) {
				if mul > largest {
					largest = mul
				}
				break
			}
		}
	}
	return largest
}

// 字串判斷法(暴力且只能用特定長度)
func FindPalindromicNumber(num int) bool {
	str := strconv.FormatInt(int64(num), 10)
	count := strings.Count(str,"")-1
	if count == 6 {
		if str[0] == str[5] && str[1] == str[4] && str[2] == str[3] {
			return true
		}
	} else if count == 5 {
		if str[0] == str[4] && str[1] == str[3] {
			return true
		}
	}
	return false
}

// 用堆疊的方式判斷是否為回文數字
func isPalindromicNumber(num int) bool {
	reversed := 0
	n1 := num
	for {
		reversed = 10 * reversed + n1 % 10
		n1 = n1 / 10
		if n1 <= 0 {
			break
		}
	}
	return reversed == num
}

func main() {
	largest := getLargest()
	fmt.Println(largest)
	fmt.Println(isPalindromicNumber(6996))
}