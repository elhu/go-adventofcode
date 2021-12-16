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

func parseLiteral(data []byte, vsAdd func(int64)) (int64, int) {
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
	return n, i
}

func parseOperator(data []byte, vsAdd func(int64)) int {
	lengthTypeId := data[0]
	processedLength := 0
	if lengthTypeId == OPERATOR_TOTAL_LENGTH {
		packetLength, err := strconv.ParseInt(string(data[1:16]), 2, 64)
		check(err)
		for processedLength < int(packetLength) {
			_, pl := parsePacket(data[16+processedLength:], vsAdd)
			processedLength += pl
		}
		return int(packetLength) + 16
	} else if lengthTypeId == OPERATOR_PACKET_COUNT {
		pc, err := strconv.ParseInt(string(data[1:12]), 2, 64)
		check(err)
		for i := int64(0); i < pc; i++ {
			_, pl := parsePacket(data[12+processedLength:], vsAdd)
			processedLength += pl
		}
		return processedLength + 12
	} else {
		panic(fmt.Errorf("Unknown length type ID: %c\n", lengthTypeId))
	}
}

func versionSum() (func() int64, func(int64)) {
	var vs = int64(0)
	get := func() int64 {
		return vs
	}
	adder := func(n int64) {
		vs += n
	}
	return get, adder
}

func parsePacket(data []byte, vsAdd func(int64)) (int, int) {
	packetVersion, err := strconv.ParseInt(string(data[0:3]), 2, 8)
	check(err)
	vsAdd(packetVersion)
	packetType, err := strconv.ParseInt(string(data[3:6]), 2, 8)
	var processedLength int
	switch packetType {
	case LITERAL:
		_, processedLength = parseLiteral(data[6:], vsAdd)
	default:
		processedLength = parseOperator(data[6:], vsAdd)
	}
	return 0, processedLength + HEADER_SIZE
}

func main() {
	data := bytes.TrimRight(files.ReadFile(os.Args[1]), "\n")
	binData := convertToBinary(data)
	vsGet, vsAdd := versionSum()
	parsePacket(binData, vsAdd)
	fmt.Println(vsGet())
}
