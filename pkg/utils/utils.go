package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func GetProgramName() string {
	return filepath.Base(os.Args[0])
}

func AssertArgs() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <inputfile>\n", GetProgramName())
		os.Exit(1)
	}
}

func AssertInput() []string {
	lines, err := ReadInput(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	return lines
}

func AssertInputInt() []int {
	ints, err := ReadInputInt(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	return ints
}

func AssertInputIntCsv() []int {
	ints, err := ReadInputIntCsv(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	return ints
}

func ReadString(s string) (input []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input, nil
}

func ReadInput(fileName string) (input []string, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input, nil
}

func ReadInputInt(fileName string) (input []int, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		input = append(input, i)
	}

	return input, nil
}

func ReadInputIntCsv(fileName string) (input []int, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ",")
		for _, n := range numbers {
			i, _ := strconv.Atoi(strings.Trim(n, " "))
			input = append(input, i)
		}
	}

	return input, nil
}

func Filter(input []string, f func(string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range input {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func StringToInt(input []string) []int {
	numbers := []int{}
	for _, s := range input {
		i, _ := strconv.Atoi(s)
		numbers = append(numbers, i)
	}

	return numbers
}

func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
}

func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func Sum(array []int) int {
	s := 0
	for _, value := range array {
		s += value
	}
	return s
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func CountRepeatCharacters(input []string) (result map[string]int) {
	result = make(map[string]int)
	for _, s := range input {
		for i := 0; i < len(s); i++ {
			result[string(s[i])]++
		}
	}

	return
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func CloneSliceBoolBool(slice [][]bool) (result [][]bool) {
	result = make([][]bool, len(slice))
	for i := range slice {
		result[i] = make([]bool, len(slice[i]))
		copy(result[i], slice[i])
	}

	return
}

func CloneSliceIntInt(slice [][]int) (result [][]int) {
	result = make([][]int, len(slice))
	for i := range slice {
		result[i] = make([]int, len(slice[i]))
		copy(result[i], slice[i])
	}

	return
}

func IsLower(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsLower(r) {
			return false
		}
	}
	return true
}
