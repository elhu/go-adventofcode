package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	LITERAL = iota
	PAIR    = iota
)

type Number struct {
	kind   int
	parent *Number
	left   *Number
	right  *Number
	value  int
}

func parseNumber(parent *Number, line string) int {
	lft := true
	for i := 0; i < len(line); i++ {
		c := line[i]
		if c == ',' {
			lft = false
		} else if c >= '0' && c <= '9' {
			stop := strings.IndexAny(line[i:], "[],")
			val, err := strconv.Atoi(line[i : i+stop])
			if err != nil {
				panic(fmt.Errorf("Non number found: %s\n", line[i:stop]))
			}
			number := &Number{kind: LITERAL, parent: parent, value: val}
			i += stop - 1
			if lft {
				parent.left = number
			} else {
				parent.right = number
			}
		} else if c == ']' {
			return i + 1
		} else if c == '[' {
			if lft {
				parent.left = &Number{kind: PAIR, parent: parent}
				i += parseNumber(parent.left, line[i+1:])
			} else {
				parent.right = &Number{kind: PAIR, parent: parent}
				i += parseNumber(parent.right, line[i+1:])
			}
		} else {
			panic("WTF")
		}
	}
	return len(line)
}

func split(number *Number) bool {
	if number.kind == LITERAL {
		lft := number.value / 2
		rgt := number.value - lft

		if number.value > 9 {
			number.kind = PAIR
			number.left = &Number{kind: LITERAL, parent: number, value: lft}
			number.right = &Number{kind: LITERAL, parent: number, value: rgt}
			return true
		}
	} else {
		return split(number.left) || split(number.right)
	}
	return false
}

func inRightSubTree(root *Number, number *Number) bool {
	if root == nil || root == number {
		return true
	}
	if root.kind == PAIR {
		if found := inRightSubTree(root.right, number); found {
			return true
		}
	}
	return false
}

func inLeftSubTree(root *Number, number *Number) bool {
	if root == nil || root == number {
		return true
	}
	if root.kind == PAIR {
		if found := inLeftSubTree(root.left, number); found {
			return true
		}
	}
	return false
}

func addRightLeft(number *Number, val int) {
	// find rightmost number
	if number.kind == LITERAL {
		number.value += val
	} else {
		addRightLeft(number.right, val)
	}
}

func addLeftRight(number *Number, val int) {
	// find left number
	if number.kind == LITERAL {
		number.value += val
	} else {
		addRightLeft(number.left, val)
	}
}

func addRightRight(root *Number, number *Number, val int) {
	// find leftmost number that doesn't share the subtree
	if inRightSubTree(root, number) {
		if root.parent != nil {
			addRightRight(root.parent, number, val)
		}
	} else {
		root = root.right
		for root.kind != LITERAL {
			root = root.left
		}
		root.value += val
	}
}

func addLeftLeft(root *Number, number *Number, val int) {
	// find rightmost number that doesn't share the subtree
	if inLeftSubTree(root, number) {
		if root.parent != nil {
			addLeftLeft(root.parent, number, val)
		}
	} else {
		root = root.left
		for root.kind != LITERAL {
			root = root.right
		}
		root.value += val
	}
}

func explode(number *Number, depth int) bool {
	if number.kind == LITERAL {
		return false
	}
	if depth == 4 {
		if number.left.kind != LITERAL || number.right.kind != LITERAL {
			panic(fmt.Errorf("Only supposed to explode pairs of literals, got: %s\n", number.toString()))
		}
		left, right := number.left.value, number.right.value
		number.kind = LITERAL
		number.value = 0
		number.left = nil
		number.right = nil
		// if number to explode is right
		if number.parent.right == number {
			// add left to the rightmost number in sibling
			addRightLeft(number.parent.left, left)
			// add right to the leftmost number in different subtree
			addRightRight(number.parent, number, right)
		} else { // if number to explode is left
			// add right to leftmost number in sibling
			addLeftRight(number.parent.right, right)
			// add left to rightmost number in different subtree
			addLeftLeft(number.parent, number, left)
		}
		return true
	}
	return explode(number.left, depth+1) || explode(number.right, depth+1)
}

func reduce(number *Number) {
	changed := true
	for changed {
		changed = explode(number, 0)
		if !changed {
			changed = split(number)
		}
	}
}

func magnitude(number *Number) int {
	if number.kind == LITERAL {
		return number.value
	}
	return magnitude(number.left)*3 + magnitude(number.right)*2
}

func add(a, b *Number) *Number {
	parent := &Number{kind: PAIR, left: a, right: b}
	a.parent = parent
	b.parent = parent
	return parent
}

func (number *Number) toString() string {
	res := ""
	if number.kind == LITERAL {
		res += strconv.Itoa(number.value)
		return res
	}
	res += "["
	if number.left.kind == LITERAL {
		res += strconv.Itoa(number.left.value)
	} else {
		res += number.left.toString()
	}
	res += (",")
	if number.right.kind == LITERAL {
		res += strconv.Itoa(number.right.value)
	} else {
		res += number.right.toString()
	}
	res += "]"
	return res
}

func parse(data []string) []*Number {
	var res []*Number
	for _, l := range data {
		number := &Number{kind: PAIR}
		parseNumber(number, l[1:])
		res = append(res, number)
	}
	return res
}

func main() {
	data := files.ReadLines(os.Args[1])
	numbers := parse(data)
	sum := numbers[0]
	for i := 1; i < len(numbers); i++ {
		sum = add(sum, numbers[i])
		reduce(sum)
	}
	fmt.Println(magnitude(sum))
}
