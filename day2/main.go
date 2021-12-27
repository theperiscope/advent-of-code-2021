package main

import (
	"AOC/pkg/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(input []string) {

	forward := utils.Filter(input, func(s string) bool { return strings.Contains(s, "forward ") })
	upOrDown := utils.Filter(input, func(s string) bool { return strings.Contains(s, "up ") || strings.Contains(s, "down ") })

	hPos := 0
	for _, v := range forward {
		i, _ := strconv.Atoi(strings.Replace(v, "forward ", "", -1))
		hPos += i
	}

	vPos := 0
	for _, v := range upOrDown {
		s := v
		if strings.Contains(v, "up ") {
			s = strings.Replace(v, "up ", "-", -1)
		} else {
			s = strings.Replace(v, "down ", "", -1)
		}
		i, _ := strconv.Atoi(s)
		vPos += i
	}

	fmt.Printf("%d * %d = %d\n", hPos, vPos, hPos*vPos)

	return
}

func part2(input []string) {

	hPos := 0
	vPos := 0
	aim := 0

	for _, v := range input {
		s := v
		if strings.Contains(v, "up ") {
			s = strings.Replace(v, "up ", "", -1)
			i, _ := strconv.Atoi(s)

			aim -= i
		} else if strings.Contains(v, "down ") {
			s = strings.Replace(v, "down ", "", -1)
			i, _ := strconv.Atoi(s)

			aim += i
		} else {
			s = strings.Replace(v, "forward ", "", -1)
			i, _ := strconv.Atoi(s)

			hPos += i
			vPos += i * aim
		}
	}

	fmt.Printf("%d * %d = %d\n", hPos, vPos, hPos*vPos)

}

func main() {

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <inputfile>\n", utils.GetProgramName())
		return
	}

	input, _ := utils.ReadInput(argsWithoutProg[0])

	part1(input)
	part2(input)
	return

}
