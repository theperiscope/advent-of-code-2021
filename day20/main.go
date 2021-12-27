package main

import (
	"AOC/pkg/utils"
	"fmt"
	"os"
)

const (
	unlit bool = false
	lit   bool = true
)

var currentOutsideEmptyBg = unlit

func count(image [][]bool) int {
	n, width, height := 0, len(image[0]), len(image)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if image[y][x] {
				n++
			}
		}
	}
	return n
}

func step(image [][]bool, corrections []bool) [][]bool {
	expandStep, w, h := 1, len(image[0]), len(image)

	// make image larger; fills outside with currentOutsideEmptyBg
	result := make([][]bool, w+(expandStep*2))
	for y := -expandStep; y <= h+expandStep-1; y++ {
		result[y+expandStep] = make([]bool, w+(expandStep*2))
		for x := -expandStep; x <= w+expandStep-1; x++ {
			result[y+expandStep][x+expandStep] = getNewPixelValue(image, y, x, corrections)
		}
	}
	if (currentOutsideEmptyBg == unlit && corrections[0] == lit) || (currentOutsideEmptyBg == lit && corrections[len(corrections)-1] == unlit) {
		currentOutsideEmptyBg = !currentOutsideEmptyBg
	}

	return result
}

func getNewPixelValue(image [][]bool, y, x int, corrections []bool) bool {
	result, w, h := 0, len(image[0]), len(image)

	// order matters
	for _, yx := range [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 0}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
		result <<= 1
		ny, nx := y+yx[0], x+yx[1]
		if ny >= 0 && ny <= h-1 && nx >= 0 && nx <= w-1 { // inside unexpanded image
			if image[ny][nx] {
				result |= 1
			}
		} else {
			if currentOutsideEmptyBg == lit {
				result |= 1
			}
		}
	}
	return corrections[result]
}

func part1(img [][]bool, corrections []bool) {
	for i := 0; i < 2; i++ {
		img = step(img, corrections)
	}
	fmt.Println("Part 1 Answer:", count(img))
}

func part2(img [][]bool, corrections []bool) {
	for i := 0; i < 50; i++ {
		img = step(img, corrections)
	}
	fmt.Println("Part 2 Answer:", count(img))
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <inputfile>\n", utils.GetProgramName())
		return
	}

	lines, _ := utils.ReadInput(argsWithoutProg[0])

	corrections := make([]bool, 0, len(lines[0]))
	for _, c := range lines[0] {
		corrections = append(corrections, c == '#')
	}

	image, width, height := [][]bool{}, len(lines[2]), len(lines[2:])
	for i := 0; i < width; i++ {
		image = append(image, []bool{})
		for j := 0; j < height; j++ {
			image[i] = append(image[i], lines[2+i][j] == '#')
		}
	}

	part1(image, corrections)
	part2(image, corrections)
}
