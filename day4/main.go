package main

import (
	"AOC/day4/bingo"
	"AOC/pkg/utils"
	"fmt"
	"strings"
)

func main() {
	utils.AssertArgs()
	lines := utils.AssertInput()
	numbers := utils.StringToInt(strings.Split(lines[0], ","))

	p1 := []*bingo.BingoCard{}
	p2 := []*bingo.BingoCard{}

	for i := 2; i < len(lines); i += 6 {
		card := bingo.NewBingoCard(lines[i : i+5])
		p1 = append(p1, card)
		p2 = append(p2, card)
	}

	//fmt.Printf("%v\n", numbers)
	fmt.Printf("%d bingo cards\n", len(p1))

	part1(numbers, p1)
	part2(numbers, p2)
}

func part1(numbers []int, cards []*bingo.BingoCard) {
	fmt.Println("PART 1\n======")
	fmt.Println()

	bingo := false
	for _, n := range numbers {
		//fmt.Printf("N = %d\n", n)
		for _, card := range cards {
			marked := card.Mark(n)
			//card.Print()
			if !marked {
				continue
			}

			if card.Check() {
				card.Print()
				fmt.Printf("BINGO! card.Sum = %d * %d = %d\n", card.Sum, n, card.Sum*n)
				bingo = true
				break
			}
		}

		if bingo {
			break
		}
	}
	fmt.Println()
}

func part2(numbers []int, cards []*bingo.BingoCard) {
	fmt.Println("PART 2\n======")
	fmt.Println()

	var winners []*bingo.BingoCard
	lastNumberCalled := -1
	for _, n := range numbers {
		//fmt.Printf("N = %d\n", n)
		for _, card := range cards {
			if card.Won {
				continue
			}

			lastNumberCalled = n
			marked := card.Mark(n)
			//card.Print()
			if !marked {
				continue
			}

			if card.Check() {
				card.Won = true
				//card.Print()
				//fmt.Printf("BINGO! card.Sum = %d * %d = %d\n\n", card.Sum, n, card.Sum*n)
				winners = append(winners, card)
			}
		}
	}

	lastWinner := winners[len(winners)-1]
	lastWinner.Print()
	fmt.Printf("BINGO! card.Sum = %d * %d = %d\n", lastWinner.Sum, lastNumberCalled, lastWinner.Sum*lastNumberCalled)

}
