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

func notVisited(path []*Device, device *Device) bool {
	for _, d := range path {
		if d == device {
			return false
		}
	}
	return true
}

func printPath(path []*Device) {
	var names []string
	for _, d := range path {
		names = append(names, d.name)
	}
	fmt.Println(strings.Join(names, " -> "))
}

func solve(startDevice *Device, devices map[string]*Device) int {
	target := "out"
	path := []*Device{startDevice}
	queue := [][]*Device{path}
	validPaths := [][]*Device{}

	for len(queue) > 0 {
		currentPath := queue[0]
		queue = queue[1:]
		currentDevice := currentPath[len(currentPath)-1]
		if currentDevice.name == target {
			validPaths = append(validPaths, currentPath)
			continue
		}
		for _, outputDevice := range currentDevice.outputs {
			if notVisited(currentPath, outputDevice) {
				newPath := make([]*Device, len(currentPath))
				copy(newPath, currentPath)
				newPath = append(newPath, outputDevice)
				queue = append(queue, newPath)
			}
		}
	}
	// for _, path := range validPaths {
	// 	printPath(path)
	// }
	return len(validPaths)
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
	fmt.Println(solve(devices["you"], devices))
}
