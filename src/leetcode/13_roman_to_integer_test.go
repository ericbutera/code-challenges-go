package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
https://leetcode.com/problems/roman-to-integer/

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, 2 is written as II in Roman numeral, just two ones added together.
12 is written as XII, which is simply X + II.
The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right.
However, the numeral for four is not IIII.
Instead, the number four is written as IV.
Because the one is before the five we subtract it making four.
The same principle applies to the number nine, which is written as IX.

There are six instances where subtraction is used:
- I can be placed before V (5) and X (10) to make 4 and 9.
- X can be placed before L (50) and C (100) to make 40 and 90.
- C can be placed before D (500) and M (1000) to make 400 and 900.

Given a roman numeral, convert it to an integer.

Symbol       Value
I             1
V             5
X             10
XL            40
L             50
XC            90
C             100
CD            400
D             500
CM            900
M             1000
*/

// type Symbol string
// const (
// 	M Symbol = "M"
// )
// var m = map[Symbol]int{
// 	M: 1_000,
// }

func romanToInt(input string) int {
	count := 0

	length := len(input)
	for x := 0; x < length; x++ {
		current := string(input[x])

		var current2 string
		hasTwo := x+2 <= length
		if hasTwo {
			current2 = string(input[x : x+2])
		}

		inc := 0

		if current == "M" {
			inc += 1_000
		} else if hasTwo && current2 == "CM" {
			inc += 900
			x++
		} else if current == "D" {
			inc += 500
		} else if hasTwo && current2 == "CD" {
			inc += 400
			x++
		} else if current == "C" {
			inc += 100
		} else if hasTwo && current2 == "XC" {
			inc += 90
			x++
		} else if current == "L" {
			inc += 50
		} else if hasTwo && current2 == "XL" {
			inc += 40
			x++
		} else if current == "X" {
			inc += 10
		} else if hasTwo && current2 == "IX" {
			inc += 9
			x++
		} else if current == "V" {
			inc += 5
		} else if hasTwo && current2 == "IV" {
			inc += 4
			x++
		} else if current == "I" {
			inc += 1
		}
		count += inc

		// fmt.Printf("finish x %v inc %v count %v\n", x, inc, count)
	}

	return count
}

func Test_13_MMMXLV(t *testing.T) {
	assert.Equal(t, 3045, romanToInt("MMMXLV"))
}

func Test_13_Example1(t *testing.T) {
	assert.Equal(t, 3, romanToInt("III"))
}

func Test_13_Example2(t *testing.T) {
	assert.Equal(t, 58, romanToInt("LVIII"))
}

func Test_13_Example3(t *testing.T) {
	assert.Equal(t, 1994, romanToInt("MCMXCIV"))
}
