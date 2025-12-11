package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Device struct {
	name    string
	outputs []*Device
}

var re = regexp.MustCompile(`(\w+)`)

func getDevice(devices map[string]*Device, name string) *Device {
	if device, exists := devices[name]; exists {
		return device
	}
	device := &Device{name: name}
	devices[name] = device
	return device
}

func solve(devices map[string]*Device, start, end string) int {
	cache := make(map[string]int)
	var doSearch func(map[string]*Device, string, string, map[string]int) int
	doSearch = func(devices map[string]*Device, start, end string, cache map[string]int) int {
		if val, found := cache[start]; found {
			return val
		}
		if start == end {
			return 1
		}
		i := 0
		for _, n := range devices[start].outputs {
			i += doSearch(devices, n.name, end, cache)
		}
		cache[start] = i
		return i
	}
	return doSearch(devices, start, end, cache)
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	devices := make(map[string]*Device)
	for _, line := range lines {
		parts := re.FindAllString(line, -1)
		device := getDevice(devices, parts[0])
		for _, outputName := range parts[1:] {
			outputDevice := getDevice(devices, outputName)
			device.outputs = append(device.outputs, outputDevice)
		}
	}
	res := 0
	res += solve(devices, "svr", "dac") * solve(devices, "dac", "fft") * solve(devices, "fft", "out")
	res += solve(devices, "svr", "fft") * solve(devices, "fft", "dac") * solve(devices, "dac", "out")
	fmt.Println(res)
}
