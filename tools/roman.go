package main

import "fmt"

var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

func rome(x, y, operatorIndex string) {
	int1 := roman[x]
	int2 := roman[y]
	res := calc(int1, int2, operatorIndex)

	if (int1 > 0 && int1 < 11) && (int2 > 0 && int2 < 11) && res > 0 {

		fmt.Println(x, operatorIndex, y, "=", getKeyRoman(roman, res))
	} else {
		panic("Недопустимая операция")
	}

}

func getKeyRoman(roman map[string]int, res int) string {
	for key, val := range roman {
		if val == res {
			return key
		}
	}
	return ""
}
