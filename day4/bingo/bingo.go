package bingo

import "fmt"

// negative values in data signify checked
type BingoCard struct {
	Data   [5][5]int
	Marked [5][5]bool
	Sum    int
	Won    bool
}

func NewBingoCard(rows []string) *BingoCard {
	card := &BingoCard{}

	if len(rows) != 5 {
		panic("exactly 5 rows expected")
	}

	for i := 0; i < len(rows); i++ {
		cols := [5]int{}
		n, err := fmt.Sscanf(rows[i], "%d %d %d %d %d", &cols[0], &cols[1], &cols[2], &cols[3], &cols[4])
		if n != 5 {
			panic("supposed to parse 5 items")
		}
		if err != nil {
			panic(err)
		}
		for j := 0; j < len(cols); j++ {
			card.Data[i][j] = cols[j]
			card.Sum += cols[j]
		}
	}

	return card
}

func (card *BingoCard) Mark(n int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if card.Data[i][j] == n {
				if card.Marked[i][j] {
					return true
				}
				card.Marked[i][j] = true
				card.Sum -= card.Data[i][j]
				return true
			}
		}
	}

	return false
}

func (card *BingoCard) Check() bool {
	for i := 0; i < 5; i++ {
		row := card.Marked[i][0] &&
			card.Marked[i][1] &&
			card.Marked[i][2] &&
			card.Marked[i][3] &&
			card.Marked[i][4]

		col := card.Marked[0][i] &&
			card.Marked[1][i] &&
			card.Marked[2][i] &&
			card.Marked[3][i] &&
			card.Marked[4][i]

		if row || col {
			return true
		}
	}

	return false
}

func (card *BingoCard) Print() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if card.Marked[i][j] {
				// https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences
				fmt.Printf(" \033[104m%2d\033[0m ", card.Data[i][j])
			} else {
				fmt.Printf(" \033[90m%2d\033[0m ", card.Data[i][j])
			}
		}
		fmt.Println()
	}
}
