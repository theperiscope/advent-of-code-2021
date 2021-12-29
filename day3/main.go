package main

import (
	"fmt"
	"math"

	"AOC/pkg/utils"
)

func part1(input []string) {

	S := len(input)
	N := len(input[0])
	ones := make([]int, N)

	for _, v := range input {
		for i := 0; i < N; i++ {
			var x int = int(v[i] - '0')
			ones[i] += x
		}
	}

	fmt.Printf("S=%d, N=%d, ones=%v\n", S, N, ones)

	for i, v := range ones {
		if v >= int(math.Ceil(float64(S/2.0))) {
			ones[i] = 1
		} else {
			ones[i] = 0
		}
	}

	fmt.Printf("S=%d, N=%d, ones=%v\n", S, N, ones)

	GAMMA := 0
	EPS := 0
	for i := 0; i < len(ones); i++ {
		GAMMA += ones[len(ones)-i-1] * int(math.Pow(2, float64(i)))
		EPS += (1 - ones[len(ones)-i-1]) * int(math.Pow(2, float64(i)))
	}

	fmt.Printf("GAMMA=%d, EPS=%d, GAMMA*EPS=%d\n", GAMMA, EPS, GAMMA*EPS)

}

func part2(input []string) {
	oxy(input, 0, len(input[0]))
	co2(input, 0, len(input[0]))
}

func oxy(input []string, i int, stop int) {
	if i > stop {
		return
	}

	S := len(input)
	ones := 0

	for _, v := range input {
		var x int = int(v[i] - '0')
		ones += x
	}

	mostCommon := 0
	if ones >= int(math.Ceil(float64(S)/float64(2))) {
		mostCommon = 1
	}

	//fmt.Printf("S=%d, ones=%d, mostCommon=%d\n", S, ones, mostCommon)

	keep := utils.Filter(input, func(s string) bool { return s[i]-'0' == byte(mostCommon) })

	if len(keep) == 1 {
		//fmt.Println(keep[0])

		OXY := 0
		n := keep[0]
		for i := 0; i < len(n); i++ {
			digit := int(n[len(n)-i-1] - '0')
			OXY += digit * int(math.Pow(2, float64(i)))
		}

		fmt.Printf("%s = %d\n", keep[0], OXY)
		return
	}

	oxy(keep, i+1, stop)
}

func co2(input []string, i int, stop int) {
	if i > stop {
		return
	}

	S := len(input)
	ones := 0

	for _, v := range input {
		var x int = int(v[i] - '0')
		ones += x
	}

	leastCommon := 1
	if ones >= int(math.Ceil(float64(S)/float64(2))) {
		leastCommon = 0
	}

	keep := utils.Filter(input, func(s string) bool { return s[i]-'0' == byte(leastCommon) })

	if len(keep) == 1 {
		//fmt.Println(keep[0])

		CO2 := 0
		n := keep[0]
		for i := 0; i < len(n); i++ {
			digit := int(n[len(n)-i-1] - '0')
			CO2 += digit * int(math.Pow(2, float64(i)))
		}

		fmt.Printf("%s = %d\n", keep[0], CO2)
		return
	}

	co2(keep, i+1, stop)

}

func main() {
	utils.AssertArgs()
	lines := utils.AssertInput()
	part1(lines)
	part2(lines)
}
