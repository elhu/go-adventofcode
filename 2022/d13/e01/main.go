package main

import (
	"adventofcode/utils/files"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func atoi(str []byte) int {
	n, err := strconv.Atoi(string(str))
	if err != nil {
		panic(err)
	}
	return n
}

const (
	NUMBER = iota
	LIST   = iota
)

type Data struct {
	kind        int
	valueNumber int
	valueList   []*Data
}

func findNumEnd(raw []byte) int {
	for i := 0; i < len(raw); i++ {
		if raw[i] == ']' || raw[i] == ',' {
			return i
		}
	}
	return len(raw)
}

func findBracketEnd(raw []byte) int {
	stack := 0
	for i := 0; i < len(raw); i++ {
		if raw[i] == '[' {
			stack++
		}
		if raw[i] == ']' {
			stack--
			if stack == 0 {
				return i
			}
		}
	}
	panic(fmt.Sprintf("Couldn't find end of list %s", raw))
}

func parsePacket(raw []byte) (*Data, int) {
	if raw[0] == '[' {
		end := findBracketEnd(raw)
		packet := &Data{kind: LIST}
		if end == 1 {
			return packet, end + 2
		}
		for i := 0; i < end; {
			subPacket, j := parsePacket(raw[i+1 : end])
			packet.valueList = append(packet.valueList, subPacket)
			i += j
		}
		return packet, end + 2
	} else if raw[0] == ']' {
		panic("wtf")
	} else {
		end := findNumEnd(raw)
		return &Data{
			kind:        NUMBER,
			valueNumber: atoi(raw[:end]),
		}, end + 1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const (
	ORDERED     = iota
	NOT_ORDERED = iota
	CONTINUE    = iota
)

func ordered(left, right *Data) int {
	if left.kind == NUMBER && right.kind == NUMBER {
		if left.valueNumber < right.valueNumber {
			return ORDERED
		} else if left.valueNumber == right.valueNumber {
			return CONTINUE
		} else {
			return NOT_ORDERED
		}
	} else if left.kind == LIST && right.kind == LIST {
		for i := 0; i < min(len(left.valueList), len(right.valueList)); i++ {
			switch ordered(left.valueList[i], right.valueList[i]) {
			case ORDERED:
				return ORDERED
			case NOT_ORDERED:
				return NOT_ORDERED
			}
		}
		if len(left.valueList) < len(right.valueList) {
			return ORDERED
		} else if len(left.valueList) == len(right.valueList) {
			return CONTINUE
		} else {
			return NOT_ORDERED
		}
	} else {
		if left.kind == NUMBER {
			return ordered(&Data{kind: LIST, valueList: []*Data{{kind: NUMBER, valueNumber: left.valueNumber}}}, right)
		} else {
			return ordered(left, &Data{kind: LIST, valueList: []*Data{{kind: NUMBER, valueNumber: right.valueNumber}}})
		}
	}
}

func packetToStr(packet *Data) string {
	if packet.kind == LIST {
		data := make([]string, len(packet.valueList))
		for i, d := range packet.valueList {
			data[i] = packetToStr(d)
		}
		return fmt.Sprintf("[%s]", strings.Join(data, ","))
	} else {
		return fmt.Sprintf("%d", packet.valueNumber)
	}
}

func main() {
	rawPairs := bytes.Split(files.ReadFile(os.Args[1]), []byte("\n\n"))
	res := 0
	for i, pair := range rawPairs {
		parts := bytes.Split(pair, []byte("\n"))
		left, _ := parsePacket(parts[0])
		right, _ := parsePacket(parts[1])
		if ordered(left, right) == ORDERED {
			res += i + 1
		}
	}
	fmt.Println(res)
}
