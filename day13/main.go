package main

import (
	"AOC/pkg/utils"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func part1(grid [][]int, folds []string, maxX int, maxY int) {
	i := 0
	s := strings.Split(strings.Replace(folds[i], "fold along ", "", -1), "=")
	foldType := s[0]
	foldParam, _ := strconv.Atoi(s[1])
	//fmt.Printf("fold along %s=%d\n", foldType, foldParam)

	if foldType == "x" {
		grid, maxX, maxY = foldX(grid, maxX, maxY, foldParam)
	} else if foldType == "y" {
		grid, maxX, maxY = foldY(grid, maxX, maxY, foldParam)
	}

	fmt.Printf("Part 1 Answer: %d\n", count(grid, maxX, maxY))
}

func part2(grid [][]int, folds []string, maxX int, maxY int) {
	for i := 0; i < len(folds); i++ {

		s := strings.Split(strings.Replace(folds[i], "fold along ", "", -1), "=")
		foldType := s[0]
		foldParam, _ := strconv.Atoi(s[1])
		//fmt.Printf("fold along %s=%d\n", foldType, foldParam)

		if foldType == "x" {
			grid, maxX, maxY = foldX(grid, maxX, maxY, foldParam)
		} else if foldType == "y" {
			grid, maxX, maxY = foldY(grid, maxX, maxY, foldParam)
		}
	}

	// make a scaled-up image from the grid so we can see the text using white color for 0s and cyan for 1s
	scale := 8
	width, height := maxX*scale, maxY*scale
	upLeft, lowRight := image.Point{0, 0}, image.Point{width, height}
	cyan := color.RGBA{100, 200, 200, 0xff}
	black := color.RGBA{0, 0, 0, 0xff}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for x := 0; x <= maxX; x++ {
		for y := 0; y < maxY; y++ {
			cc := black
			if grid[y][x] == 1 {
				cc = cyan
			}

			for i := x * scale; i < (x*scale)+scale; i++ {
				for j := y * scale; j < (y*scale)+scale; j++ {
					img.Set(i, j, cc)
				}
			}
		}
	}

	tmpFile, err := ioutil.TempFile(os.TempDir(), "day13-*.png")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}
	png.Encode(tmpFile, img)

	cmd := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", tmpFile.Name())
	cmd.Run()
	fmt.Println("Part 2 Answer: see " + tmpFile.Name())
}

func count(grid [][]int, maxX int, maxY int) int {
	n := 0
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if grid[i][j] == 1 {
				n++
			}
		}
	}

	return n
}

func foldY(grid [][]int, maxX, maxY, y int) (newGrid [][]int, newMaxX, newMaxY int) {
	newMaxY = y
	newMaxX = maxX

	newGrid = utils.CloneSliceIntInt(grid)
	for i, n := y+1, 1; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if grid[i][j] == 1 {
				newGrid[y-n][j] = grid[i][j]
			}
		}
		n++
	}
	return
}

func foldX(grid [][]int, maxX, maxY, x int) (newGrid [][]int, newMaxX, newMaxY int) {
	newMaxY = maxY
	newMaxX = x

	newGrid = utils.CloneSliceIntInt(grid)
	for i := 0; i <= maxY; i++ {
		for j, n := x+1, 1; j <= maxX; j++ {
			if grid[i][j] == 1 {
				newGrid[i][x-n] = grid[i][j]
			}
			n++
		}
	}
	return
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <inputfile>\n", utils.GetProgramName())
		return
	}

	lines, _ := utils.ReadInput(argsWithoutProg[0])
	data := [][]int{}
	folds := utils.Filter(lines, func(s string) bool { return strings.HasPrefix(s, "fold along") })
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		data = append(data, utils.StringToInt(strings.Split(line, ",")))
	}

	maxX, maxY := 0, 0

	// read input data
	for i := 0; i < len(data); i++ {
		if data[i][0] > maxX {
			maxX = data[i][0]
		}
		if data[i][1] > maxY {
			maxY = data[i][1]
		}
	}

	// map input data onto the grid
	grid := make([][]int, maxY+1)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, maxX+1)
	}

	for i := 0; i < len(data); i++ {
		grid[data[i][1]][data[i][0]] = 1
	}

	part1(grid, folds, maxX, maxY)
	part2(grid, folds, maxX, maxY)
}
