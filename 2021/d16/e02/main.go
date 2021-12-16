package main

import (
	"adventofcode/utils/files"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertToBinary(data []byte) []byte {
	data = bytes.ToLower(data)
	var res []byte
	for _, c := range data {
		n, err := strconv.ParseInt(string(c), 16, 8)
		check(err)
		b := fmt.Sprintf("%04s", strconv.FormatInt(n, 2))
		res = append(res, []byte(b)...)
	}
	return res
}

const LITERAL = 4
const OPERATOR_TOTAL_LENGTH = '0'
const OPERATOR_PACKET_COUNT = '1'
const HEADER_SIZE = 6

func parseLiteral(data []byte) (int, int) {
	var bits []byte
	i := 0
	for ; i < len(data)-4; i += 5 {
		bits = append(bits, data[i+1:i+5]...)
		if data[i] == '0' {
			break
		}
	}
	i += 5
	n, err := strconv.ParseInt(string(bits), 2, 64)
	check(err)
	return int(n), i
}

var operators = map[int64]func([]int) int{
	0: func(vals []int) int {
		sum := 0
		for _, v := range vals {
			sum += v
		}
		return sum
	},
	1: func(vals []int) int {
		prod := 1
		for _, v := range vals {
			prod *= v
		}
		return prod
	},
	2: func(vals []int) int {
		min := vals[0]
		for _, v := range vals {
			if v < min {
				min = v
			}
		}
		return min
	},
	3: func(vals []int) int {
		max := vals[0]
		for _, v := range vals {
			if v > max {
				max = v
			}
		}
		return max
	},
	5: func(vals []int) int {
		if vals[0] > vals[1] {
			return 1
		}
		return 0
	},
	6: func(vals []int) int {
		if vals[0] < vals[1] {
			return 1
		}
		return 0
	},
	7: func(vals []int) int {
		if vals[0] == vals[1] {
			return 1
		}
		return 0
	},
}

func parseOperator(packetType int64, data []byte) (int, int) {
	lengthTypeId := data[0]
	processedLength := 0
	if lengthTypeId == OPERATOR_TOTAL_LENGTH {
		packetLength, err := strconv.ParseInt(string(data[1:16]), 2, 64)
		check(err)
		var packetVals []int
		for processedLength < int(packetLength) {
			val, pl := parsePacket(data[16+processedLength:])
			packetVals = append(packetVals, val)
			processedLength += pl
		}
		return operators[packetType](packetVals), int(packetLength) + 16
	} else if lengthTypeId == OPERATOR_PACKET_COUNT {
		pc, err := strconv.ParseInt(string(data[1:12]), 2, 64)
		check(err)
		var packetVals []int
		for i := int64(0); i < pc; i++ {
			val, pl := parsePacket(data[12+processedLength:])
			packetVals = append(packetVals, val)
			processedLength += pl
		}
		return operators[packetType](packetVals), processedLength + 12
	} else {
		panic(fmt.Errorf("Unknown length type ID: %c\n", lengthTypeId))
	}
}

func parsePacket(data []byte) (int, int) {
	_, err := strconv.ParseInt(string(data[0:3]), 2, 8)
	check(err)
	packetType, err := strconv.ParseInt(string(data[3:6]), 2, 8)
	var processedLength int
	var val int

	switch packetType {
	case LITERAL:
		val, processedLength = parseLiteral(data[6:])
	default:
		val, processedLength = parseOperator(packetType, data[6:])
	}
	return val, processedLength + HEADER_SIZE
}

func main() {
	data := bytes.TrimRight(files.ReadFile(os.Args[1]), "\n")
	binData := convertToBinary(data)
	res, _ := parsePacket(binData)
	fmt.Println(res)
}
