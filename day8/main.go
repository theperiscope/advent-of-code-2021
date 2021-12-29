package main

import (
	"AOC/pkg/utils"
	"fmt"
	"strings"
)

func part1(lines []string) {
	n := 0
	for _, line := range lines {
		x := strings.Split(line, " | ")
		y := strings.Split(x[1], " ")
		for _, z := range y {
			if len(z) == 2 || len(z) == 3 || len(z) == 4 || len(z) == 7 { // 1, 4, 7, 8 numbers have 2, 3, 4, 7 segments uniquely
				n++
			}
		}
	}

	fmt.Println("Part 1 Answer:", n)
}

func part2(lines []string) {
	sum := 0
	for _, line := range lines {
		x := strings.Split(line, " | ")
		y := strings.Split(x[0], " ")
		z := strings.Split(x[1], " ")

		_1 := utils.Filter(y, func(s string) bool { return len(s) == 2 })[0]
		_4 := utils.Filter(y, func(s string) bool { return len(s) == 4 })[0]
		_7 := utils.Filter(y, func(s string) bool { return len(s) == 3 })[0]
		_8 := utils.Filter(y, func(s string) bool { return len(s) == 7 })[0]

		// if 1 is ab, and 7 is dab, _7_extra will be "d"
		_7_extra := "" // the extra character we need to add to 1 to get 7
		for _, c := range _7 {
			if !strings.Contains(_1, string(c)) {
				_7_extra = string(c)
				break
			}
		}

		// 0, 6, and 9 always have length of 6 segments
		_0_6_9 := utils.Filter(y, func(s string) bool { return len(s) == 6 })

		// of the 6 parts, 4 are common/shared, we need only the remaining ones
		_0_6_9_map := utils.CountRepeatCharacters(_0_6_9)
		_0_6_9_remaining := []string{}
		for k, v := range _0_6_9_map {
			if v != 3 {
				_0_6_9_remaining = append(_0_6_9_remaining, k)
			}
		}

		_x := " "
		for _, c := range _0_6_9_remaining {
			if strings.Contains(_1, c) {
				_x = c
				break
			}
		}

		// we can identify the 6 in the 0-6-9 set from the 6-remaining parts (3)
		// because it does not contain elements of "1"
		_6 := utils.Filter(_0_6_9, func(s string) bool { return !strings.Contains(s, _x) })[0]

		// 2, 3, and 5 have 5 segments
		_2_3_5 := utils.Filter(y, func(s string) bool { return len(s) == 5 })

		// split the 2,5 and 5 segments in to shared/remaining; we only need the shared ones
		// the shared segments are the horizonal segments; we also know exactly first(top)
		// one is because of the "_7_extra"
		_2_3_5_map := utils.CountRepeatCharacters(_2_3_5)
		_2_3_5_shared := []string{}
		for k, v := range _2_3_5_map {
			if v == 3 {
				_2_3_5_shared = append(_2_3_5_shared, k)
			}
		}

		_2_3_5_shared_2 := utils.Filter(_2_3_5_shared, func(s string) bool { return s != _7_extra })

		// number 3 is represent by "7" + middle + bottom horizontal segments;
		// last two we calculate above because 2, 3 and 5 all share them
		_3_chars := []string{}
		_3_chars = append(_3_chars, strings.Split(_7, "")...)
		_3_chars = append(_3_chars, _2_3_5_shared_2...)

		_3 := utils.Filter(_2_3_5, func(s string) bool {
			return utils.SortString(s) == utils.SortString(strings.Join(_3_chars, ""))
		})[0]

		// 9 is one difference away from 3
		_0_9 := utils.Filter(y, func(s string) bool { return len(s) == 6 && s != _6 })
		_9 := ""
		for _, s := range _0_9 {
			n := s
			for _, c := range _3 {
				n = strings.Replace(n, string(c), "", -1)
			}
			if len(n) == 1 {
				_9 = s
				break
			}
		}

		// now it's easy to compute 0 too
		_0 := utils.Filter(y, func(s string) bool { return len(s) == 6 && s != _6 && s != _9 })[0]

		_2_5 := utils.Filter(_2_3_5, func(s string) bool { return s != _3 })
		_9map := utils.CountRepeatCharacters([]string{_9, _2_5[0]})
		_9map_ones := 0
		for _, v := range _9map {
			if v == 1 {
				_9map_ones++
			}
		}

		// 5 has only 1 difference, when comparing 2 and 5 to 9 ... we use that to distinguish between
		_5 := _2_5[0]
		_2 := _2_5[1]
		if _9map_ones > 1 {
			_5 = _2_5[1]
			_2 = _2_5[0]
		}

		// finally, build the decoding map
		decoded := map[string]int{
			utils.SortString(_0): 0,
			utils.SortString(_1): 1,
			utils.SortString(_2): 2,
			utils.SortString(_3): 3,
			utils.SortString(_4): 4,
			utils.SortString(_5): 5,
			utils.SortString(_6): 6,
			utils.SortString(_7): 7,
			utils.SortString(_8): 8,
			utils.SortString(_9): 9,
		}

		// compute answer
		answer := decoded[utils.SortString(z[0])]*1000 +
			decoded[utils.SortString(z[1])]*100 +
			decoded[utils.SortString(z[2])]*10 +
			decoded[utils.SortString(z[3])]

		sum += answer
	}
	fmt.Printf("Part 2 Answer: %d\n", sum)
}

func main() {
	utils.AssertArgs()
	lines := utils.AssertInput()
	part1(lines)
	part2(lines)
}
