package main

import (
	"AOC/pkg/utils"
	"fmt"
	"regexp"
	"strconv"
)

func explode(input string) string {
	level := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '[' {
			level++
			continue
		}
		if input[i] == ']' {
			level--
			continue
		}
		if input[i-1] == '[' && level == 5 { //explode
			initialPos := i - 1
			left, right := "", ""
			for input[i] >= '0' && input[i] <= '9' {
				left += string(input[i])
				i++
			}
			i++ // the ,
			for input[i] >= '0' && input[i] <= '9' {
				right += string(input[i])
				i++
			}
			endPos := i // the ]
			l, _ := strconv.Atoi(left)
			r, _ := strconv.Atoi(right)

			prevCommaPos, nextCommaPos := -1, -1
			for i := initialPos; i >= 0; i-- {
				if input[i] == ',' {
					prevCommaPos = i
					break
				}
			}
			for i := endPos; i < len(input); i++ {
				if input[i] == ',' {
					nextCommaPos = i
					break
				}
			}
			prevNumber := -1
			prevNumberBegin, prevNumberEnd := prevCommaPos-1, prevCommaPos-1
			if prevCommaPos >= 0 {
				for input[prevNumberBegin] < '0' || input[prevNumberBegin] > '9' { // skip [ before comma for example
					prevNumberBegin--
					prevNumberEnd--
				}
				for input[prevNumberBegin] >= '0' && input[prevNumberBegin] <= '9' {
					prevNumberBegin--
				}
				prevNumberBegin++
				prevNumber, _ = strconv.Atoi(input[prevNumberBegin : prevNumberEnd+1])
			}
			nextNumber := -1
			nextNumberBegin, nextNumberEnd := nextCommaPos+1, nextCommaPos+1
			if nextCommaPos >= 0 {
				for input[nextNumberBegin] < '0' || input[nextNumberBegin] > '9' { // skip [ after comma for example
					nextNumberBegin++
					nextNumberEnd++
				}
				for input[nextNumberEnd] >= '0' && input[nextNumberEnd] <= '9' {
					nextNumberEnd++
				}
				nextNumberEnd--
				nextNumber, _ = strconv.Atoi(input[nextNumberBegin : nextNumberEnd+1])
			}

			// work, order is important so we don't lose our positions
			s := input
			if nextCommaPos >= 0 {
				after := ""
				s, after = s[0:nextNumberBegin], s[nextNumberEnd+1:]
				s += fmt.Sprintf("%d", r+nextNumber)
				s += after
			}

			s = s[0:initialPos] + "0" + s[endPos+1:]

			if prevCommaPos >= 0 {
				after := ""
				s, after = s[0:prevNumberBegin], s[prevNumberEnd+1:]
				s += fmt.Sprintf("%d", l+prevNumber)
				s += after
			}

			return s
		}
	}

	return input
}

func split(input string) string {
	r := regexp.MustCompile("\\d{2,}")
	index := r.FindStringIndex(input)
	if len(index) > 0 {
		begin, end := index[0], index[0]
		for input[end] >= '0' && input[end] <= '9' {
			end++
		}
		end--
		n, _ := strconv.Atoi(input[begin : end+1])
		l := n / 2
		r := n - (n / 2)

		return input[0:begin] + fmt.Sprintf("[%d,%d]", l, r) + input[end+1:]
	}

	return input
}

func plus(n1, n2 string) string {
	return fmt.Sprintf("[%s,%s]", n1, n2)
}

func magnitude(input string) int {
	s := utils.Stack{}
	sums := utils.Stack{}
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c == '[' {
			s.Push("PAIR")
			sums.Push("PAIR")
		} else if c >= '0' && c <= '9' {
			n := ""
			for input[i] >= '0' && input[i] <= '9' {
				n += string(input[i])
				i++
			}
			i--
			nn, _ := strconv.Atoi(n)
			sums.Push(fmt.Sprintf("%d", nn))
			s.Push(fmt.Sprintf("%d", c-'0'))
		} else if c == ',' {
			s.Push("COMMA")
			sums.Push("COMMA")
		} else if c == ']' {
			n1, _ := sums.Pop()
			sums.Pop() // COMMA
			n3, _ := sums.Pop()
			sums.Pop() // PAIR
			l, _ := strconv.Atoi(n3)
			r, _ := strconv.Atoi(n1)
			n := 3*l + 2*r
			sums.Push(fmt.Sprintf("%d", n))

			p1, _ := s.Pop()
			s.Pop() // COMMA
			p3, _ := s.Pop()
			s.Pop() // PAIR
			s.Push(fmt.Sprintf("[%s,%s]", p3, p1))
		} else {
			panic("oh no")
		}
	}

	ss, _ := sums.Pop()
	v, _ := strconv.Atoi(ss)
	return v
}

func part1(lines []string) {
	input1, input2 := lines[0], lines[1]
	for i := 1; i < len(lines); i++ {
		s := plus(input1, input2)
		for {
			ss := explode(s)
			if ss == s {
				ss = split(s)
				if ss == s {
					//fmt.Println("Done.")
					break
				}
				//fmt.Println("After split,", ss)
			} else {
				//fmt.Println("After explode,", ss)
			}
			s = ss
		}
		if i == len(lines)-1 {
			input1 = s
			break
		}
		input1, input2 = s, lines[i+1]
	}

	fmt.Println(input1)
	fmt.Println(magnitude(input1))
}

func part2(lines []string) {
	max := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if lines[i] == lines[j] {
				continue
			}

			input1, input2 := lines[i], lines[j]
			s := plus(input1, input2)
			for {
				ss := explode(s)
				if ss == s {
					ss = split(s)
					if ss == s {
						//fmt.Println("Done.")
						break
					}
					//fmt.Println("After split,", ss)
				} else {
					//fmt.Println("After explode,", ss)
				}
				s = ss
			}
			if i == len(lines)-1 {
				input1 = s
				break
			}
			input1, input2 = s, lines[i+1]

			m := magnitude(input1)
			if m > max {
				max = m
			}
		}
	}

	fmt.Println(max)
}

func main() {
	utils.AssertArgs()
	lines := utils.AssertInput()
	part1(lines)
	part2(lines)
}
