package main

import (
	"AOC/pkg/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type vars map[string]int

type instruction struct {
	op string
	r1 string
	r2 interface{}
}

type alu struct {
	variables    vars
	instructions map[string]func(a *alu, op1 string, op2 interface{})
	inputs       []int
}

func (a *alu) reset() {
	for k := range a.variables {
		a.variables[k] = 0
	}
}

var instructions map[string]func(a *alu, op1 string, op2 interface{}) = map[string]func(a *alu, op1 string, op2 interface{}){
	"inp": func(a *alu, op1 string, op2 interface{}) {
		if len(a.inputs) == 0 {
			panic("STOP! no inputs available")
		}
		(*a).variables[op1] = a.inputs[0]
		a.inputs = a.inputs[1:]
		fmt.Println("INP:", a.variables)
	},
	"add": func(a *alu, op1 string, op2 interface{}) {
		switch op2 := op2.(type) {
		case int:
			(*a).variables[op1] = (*a).variables[op1] + op2
		case string:
			(*a).variables[op1] = (*a).variables[op1] + (*a).variables[op2]
		default:
			panic("add bad type")
		}
	},
	"mul": func(a *alu, op1 string, op2 interface{}) {
		switch op2 := op2.(type) {
		case int:
			(*a).variables[op1] = (*a).variables[op1] * op2
		case string:
			(*a).variables[op1] = (*a).variables[op1] * (*a).variables[op2]
		default:
			panic("mul bad type")
		}
	},
	"div": func(a *alu, op1 string, op2 interface{}) {
		switch op2 := op2.(type) {
		case int:
			(*a).variables[op1] = (*a).variables[op1] / op2
		case string:
			(*a).variables[op1] = (*a).variables[op1] / (*a).variables[op2]
		default:
			panic("div bad type")
		}
	},
	"mod": func(a *alu, op1 string, op2 interface{}) {
		switch op2 := op2.(type) {
		case int:
			(*a).variables[op1] = (*a).variables[op1] % op2
		case string:
			(*a).variables[op1] = (*a).variables[op1] % (*a).variables[op2]
		default:
			panic("mod bad type")
		}
	},
	"eql": func(a *alu, op1 string, op2 interface{}) {
		switch op2 := op2.(type) {
		case int:
			if (*a).variables[op1] == op2 {
				(*a).variables[op1] = 1
			} else {
				(*a).variables[op1] = 0
			}
		case string:
			if (*a).variables[op1] == (*a).variables[op2] {
				(*a).variables[op1] = 1
			} else {
				(*a).variables[op1] = 0
			}
		default:
			panic("eql bad type")
		}
	},
}

func (a *alu) readProgram(lines []string) []instruction {
	p := []instruction{}
	for _, line := range lines {
		if strings.Contains(line, "--") { // comment
			line = line[0:strings.Index(line, "--")]
		}
		l := strings.Trim(line, " \t")
		if len(l) == 0 {
			continue
		}
		ss := strings.Split(l, " ")
		if len(ss) == 2 {
			p = append(p, instruction{op: ss[0], r1: ss[1]})
		} else if len(ss) == 3 {
			n, err := strconv.Atoi(ss[2])
			if err == nil {
				p = append(p, instruction{op: ss[0], r1: ss[1], r2: n})
			} else {
				p = append(p, instruction{op: ss[0], r1: ss[1], r2: ss[2]})
			}
		} else {
			panic("STOP! bad instruction")
		}
	}
	return p
}

func (a *alu) execute(program []string) {
	a.reset()
	for _, instruction := range a.readProgram(program) {
		if _, ok := a.instructions[instruction.op]; !ok {
			panic("oh no!")
		}
		a.instructions[instruction.op](a, instruction.r1, instruction.r2)
	}
}

func part1() {
}

func part2() {
}

func testingAndPOCSample() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <inputfile>\n", utils.GetProgramName())
		return
	}

	lines, _ := utils.ReadInput(argsWithoutProg[0])

	// https://go.dev/play/p/S98GjeaGBX0
	// This WaitGroup is used to wait for all the
	// goroutines launched here to finish. Note: if a WaitGroup is
	// explicitly passed into functions, it should be done *by pointer*.
	var wg sync.WaitGroup

	for i := 94992992796199; i <= 94992992796199; i += 100000000000000 {
		wg.Add(1)
		// Avoid re-use of the same `i` value in each goroutine closure.
		// See [the FAQ](https://golang.org/doc/faq#closures_and_goroutines)
		// for more details.
		i := i
		start := i
		end := i

		// Wrap the worker call in a closure that makes sure to tell
		// the WaitGroup that this worker is done. This way the worker
		// itself does not have to be aware of the concurrency primitives
		// involved in its execution.
		go func() {
			defer wg.Done()
			worker(lines, i, start, end, instructions)
		}()
	}

	// Block until the WaitGroup counter goes back to 0;
	// all the workers notified they're done.
	wg.Wait()

	// 11931881141161
	// 94992992796199
	x, y, z := 0, 0, 0
	fmt.Println(x, y, z)
	x, y, z = inputProgramWithoutAlu(9, x, y, z, 0) // 1 7 7
	fmt.Println(x, y, z)
	x, y, z = inputProgramWithoutAlu(4, x, y, z, 1) // 1 12 194
	fmt.Println(x, y, z)
	x, y, z = inputProgramWithoutAlu(9, x, y, z, 2) // 1 6 5050
	fmt.Println(x, y, z)
	x, y, z = inputProgramWithoutAlu(9, x, y, z, 3) // 1 7 131307
	fmt.Println(x, y, z)
	x, y, z = inputProgramWithoutAlu(2, x, y, z, 4) // 1 9 3413991
	fmt.Println(x, y, z)
	x, y, z = inputProgramWithoutAlu(9, x, y, z, 5)
	fmt.Println(x, y, z)
	x, y, z = inputProgramWithoutAlu(9, x, y, z, 6)
	fmt.Println(x, y, z)
	x, y, z = inputProgramWithoutAlu(2, x, y, z, 7)
	fmt.Println(x, y, z)
	x, y, z = inputProgramWithoutAlu(7, x, y, z, 8)
	fmt.Println(x, y, z)
	x, y, z = inputProgramWithoutAlu(9, x, y, z, 9)
	fmt.Println(x, y, z)
	x, y, z = inputProgramWithoutAlu(6, x, y, z, 10)
	fmt.Println(x, y, z)
	x, y, z = inputProgramWithoutAlu(1, x, y, z, 11)
	fmt.Println(x, y, z)
	x, y, z = inputProgramWithoutAlu(9, x, y, z, 12)
	fmt.Println(x, y, z)
	x, y, z = inputProgramWithoutAlu(9, x, y, z, 13)
	fmt.Println(x, y, z)
	fmt.Println("=====")

	xx := 1
	for i := 13; i >= 1; i-- {
		fmt.Println(xx * B[i])
		xx *= B[i]
	}
}

func main() {
	solutions := []int{}
	solutionsSearch(0, 0, "", func(s string) {
		n, _ := strconv.Atoi(s)
		solutions = append(solutions, n)
	})

	min, max := utils.MinMax(solutions)
	fmt.Println("Part 1 Answer:", max)
	fmt.Println("Part 2 Answer:", min)
}

// representation of input data as coefficients from analysis of instruction in Excel
var A []int = []int{13, 11, 12, 10, 14, -1, 14, -16, -8, 12, -16, -13, -6, -6}
var B []int = []int{1, 1, 1, 1, 1, 26, 1, 26, 26, 1, 26, 26, 26, 26}
var C []int = []int{6, 11, 5, 6, 8, 14, 9, 4, 7, 13, 11, 11, 6, 1}

// generated from B in reverse order 26, 26*26, 26*26*26, 26*26*26*26, 26*26*26*26*1, 26*26*26*26*1*26, etc.
var maxZ []int = []int{8031810176, 8031810176, 8031810176, 8031810176, 8031810176, 8031810176, 308915776, 308915776, 11881376, 456976, 456976, 17576, 676, 26}

// converted from excel rules. A, B, C are the parameters different for each digit position in instructions 5, 6 and 16
func inputProgramWithoutAlu(w, x, y, z, stepIndex int) (newX int, newY int, newZ int) {
	a, b, c := A[stepIndex], B[stepIndex], C[stepIndex]
	newX = (z % 26) + a // ref1
	newZ = z / b        // ref2, ref3
	if newX == w {
		newX = 1
	} else {
		newX = 0
	}
	if newX == 0 {
		newX = 1
		newY = 26 // ref3
	} else {
		newX = 0
		newY = 1
	}
	newZ *= newY // ref3
	newY = (w + c) * newX
	newZ += ((w + c) * newX) // ref3
	return
}

// from inputProgramWithoutAlu focused on "z" only
func inputProgramWithoutAlu_Z(w, z, stepIndex int) int {
	if z%26+A[stepIndex] == w { // ref1
		return z / B[stepIndex] // ref2
	} else {
		return 26*(z/B[stepIndex]) + w + C[stepIndex] // ref3
	}
}

func solutionsSearch(depth int, z int, answer string, answerAction func(string)) {
	if depth == 14 {
		if z == 0 {
			answerAction(answer)
		}
		return
	} else if z >= maxZ[depth] { // this really speeds things up
		return
	}
	// try 1..9 for each position with Z value based on func
	for i := 1; i <= 9; i++ {
		solutionsSearch(depth+1, inputProgramWithoutAlu_Z(i, z, depth), answer+string(byte('0'+i)), answerAction)
	}
}

func worker(input []string, i, start, end int, f map[string]func(a *alu, op1 string, op2 interface{})) {
	a := &alu{
		variables:    map[string]int{"x": 0, "w": 0, "y": 0, "z": 0},
		instructions: f,
		inputs:       []int{},
	}

	for n := end; n >= start; n -= 1 {
		a.inputs = splitInt(n)
		a.execute(input)
	}
}

func splitInt(n int) []int {
	if n < 0 {
		n = -n
	}
	totalDigits := int(math.Log10(float64(n)) + 1)
	digits := make([]int, totalDigits)
	i := totalDigits - 1
	for {
		lastDigit := n % 10
		digits[i] = lastDigit
		i, n = i-1, n/10
		if n == 0 {
			break
		}
	}
	return digits
}
