package main

import (
	"AOC/pkg/utils"
	"fmt"
	"sort"
	"strings"
)

func part1(characterLines [][]string) (incompleteLineStacks []utils.Stack) {
	incompleteLineStacks, score, m := []utils.Stack{}, 0, map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	for _, line := range characterLines {
		isCorrupted, stack := false, utils.Stack{}
		for _, char := range line {
			if char == "(" || char == "[" || char == "{" || char == "<" {
				stack.Push(char)
				continue
			}

			c, _ := stack.Peek()
			if (c == "(" && char == ")") || (c == "[" && char == "]") || (c == "{" && char == "}") || (c == "<" && char == ">") {
				stack.Pop()
				continue
			}

			if len(stack) > 1 {
				isCorrupted = true
				score += m[char]
				break
			}
		}
		if !isCorrupted {
			incompleteLineStacks = append(incompleteLineStacks, stack)
		}
	}

	fmt.Println("Part 1 Answer:", score)
	return
}

func part2(incompleteLineStacks []utils.Stack) {
	scores := []int{}
	for _, stack := range incompleteLineStacks {
		score, m := 0, map[string]int{"(": 1, "[": 2, "{": 3, "<": 4}
		for !stack.IsEmpty() {
			char, _ := stack.Pop()
			score = (score * 5) + m[char]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	middle := (len(scores) - 1) / 2 // expected to always have odd length of scores
	fmt.Println("Part 2 Answer:", scores[middle])
}

func main() {
	utils.AssertArgs()
	lines := utils.AssertInput()
	characterLines := [][]string{}
	for _, line := range lines {
		characterLines = append(characterLines, strings.Split(line, ""))
	}

	incompleteLines := part1(characterLines)
	part2(incompleteLines)
}
