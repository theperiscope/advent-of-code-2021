package main

import (
	"AOC/pkg/utils"
	"fmt"
)

func roll(sequence int, diceSides int) (d1, d2, d3 int) {
	return (0+(sequence-1)*3)%diceSides + 1, (1+(sequence-1)*3)%diceSides + 1, (2+(sequence-1)*3)%diceSides + 1
}

type state struct {
	p1p int
	p1s int
	p2p int
	p2s int
}

func (s state) checkGameOver(targetPoints int) bool {
	return s.p1s >= targetPoints || s.p2s >= targetPoints
}

func (s state) moveAndScoreP1(moveBy int) state {
	ppos, step := []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9}, moveBy%10
	newPos := ppos[(s.p1p+step)%10]
	return state{p1p: newPos, p1s: s.p1s + newPos, p2p: s.p2p, p2s: s.p2s}
}

func (s state) moveAndScoreP2(moveBy int) state {
	ppos, step := []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9}, moveBy%10
	newPos := ppos[(s.p2p+step)%10]
	return state{p1p: s.p1p, p1s: s.p1s, p2p: newPos, p2s: s.p2s + newPos}
}

func game(s state, targetPoints, diceSides int) (diceRolls, winningPlayer, winningPlayerScore, losingPlayer, losingPlayerScore int) {
	i := -1
	for s.p1s < targetPoints && s.p2s < targetPoints {
		i += 2
		d1, d2, d3 := roll(i, diceSides)
		s = s.moveAndScoreP1(d1 + d2 + d3)
		if s.checkGameOver(targetPoints) {
			diceRolls, winningPlayer, winningPlayerScore, losingPlayer, losingPlayerScore = i*3, 1, s.p1s, 2, s.p2s
			continue
		}

		d1, d2, d3 = roll(i+1, diceSides)
		s = s.moveAndScoreP2(d1 + d2 + d3)
		if s.checkGameOver(targetPoints) {
			diceRolls, winningPlayer, winningPlayerScore, losingPlayer, losingPlayerScore = i*3, 2, s.p2s, 1, s.p1s
			continue
		}
	}
	return
}

func part1(initialState state) {
	diceRolls, winningPlayer, _, losingPlayer, losingPlayerScore := game(initialState, 1000, 100)
	fmt.Printf("Player %d wins after %d rolls.\n", winningPlayer, diceRolls)
	fmt.Printf("Part 1 Answer: Player %d loses, %d * %d = %d\n", losingPlayer, losingPlayerScore, diceRolls*3, losingPlayerScore*diceRolls)
}

func part2(initialState state) {
	// key:sum, value: number of times it occurs within all permutations
	possibleRollSumCounts := map[int]int{}
	for r1 := 1; r1 <= 3; r1++ {
		for r2 := 1; r2 <= 3; r2++ {
			for r3 := 1; r3 <= 3; r3++ {
				possibleRollSumCounts[r1+r2+r3] += 1
			}
		}
	}

	currentStates := map[state]int{initialState: 1}

	p1w, p2w := 0, 0

	for len(currentStates) > 0 {
		newStates := map[state]int{}

		// P1 first plays all rolls in all possible current states and we calculate new states
		for roll, rollCount := range possibleRollSumCounts {
			for state, stateCount := range currentStates {
				s := state.moveAndScoreP1(roll)
				if s.p1s >= 21 {
					p1w += rollCount * stateCount
				} else {
					newStates[s] += rollCount * stateCount
				}
			}
		}

		currentStates, newStates = newStates, map[state]int{}
		if len(currentStates) == 0 {
			break
		}

		// P2 then plays all rolls in updated possible current states and we calculate again new states
		for roll, rollCount := range possibleRollSumCounts {
			for state, stateCount := range currentStates {
				s := state.moveAndScoreP2(roll)
				if s.p2s >= 21 {
					p2w += rollCount * stateCount
				} else {
					newStates[s] += rollCount * stateCount
				}
			}
		}

		currentStates = newStates
		if len(currentStates) == 0 {
			break
		}
	}

	fmt.Println("Part 2 Answer:", utils.MaxInt(p1w, p2w))
}

func main() {
	initialState := state{p1p: 4, p2p: 7, p1s: 0, p2s: 0}
	part1(initialState)
	part2(initialState)
}
