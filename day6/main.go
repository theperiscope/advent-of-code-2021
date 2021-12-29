package main

import (
	"AOC/pkg/utils"
	"fmt"
	"log"
	"os"
	"strconv"
)

func gen(targetDay int, initialState int) int {
	remaining, s := targetDay-(initialState+1), 1
	for remaining >= 0 {
		s += gen(remaining, 8)
		remaining -= 7
	}
	return s
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 2 {
		fmt.Printf("Usage: %s <targetDay> <inputfile>\n", utils.GetProgramName())
		os.Exit(1)
	}

	targetDay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lanternFish, err := utils.ReadInputIntCsv(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	s, cache := 0, map[int]int{}
	for _, fish := range lanternFish {
		g, ok := cache[fish]
		if !ok {
			g = gen(targetDay, fish)
			cache[fish] = g
		}
		s += g
	}

	fmt.Println("Sum:", s)
}
