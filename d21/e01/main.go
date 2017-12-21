package main

import (
  "os"
  "io/ioutil"
  "strings"
  "fmt"
  "bytes"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func getImage(filename string) []string {
  data, err := ioutil.ReadFile(filename)
  check(err)

  return strings.Split(string(data), "\n")
}

func rotate(lines []string) []string {
  n := len(lines)
  ret := make([][]byte, n)
  for i := 0; i < n; i++ {
    ret[i] = make([]byte, n)
    for j := 0; j < n; j++ {
      ret[i][j] = lines[n - j - 1][i]
    }
  }
  stringifiedRet := make([]string, n)
  for i, line := range ret {
    stringifiedRet[i] = string(line)
  }
  return stringifiedRet
}

func flip(s string) string {
  l := bytes.Split([]byte(s), []byte("/"))
  for i := 0; i < len(l); i++ {
    l[i][0], l[i][len(l[i]) - 1] = l[i][len(l[i]) - 1], l[i][0]
  }
  return string(bytes.Join(l, []byte("/")))
}

func permutations(s string) [8]string {
  res := [8]string{}
  // Initial pos + flip
  res[0] = s
  res[1] = flip(s)

  // 90 degrees + flip
  res[2] = strings.Join(rotate(strings.Split(s, "/")), "/")
  res[3] = flip(res[2])

  // 180 degrees + flip
  res[4] = strings.Join(rotate(strings.Split(res[2], "/")), "/")
  res[5] = flip(res[4])

  // 270 degrees + flip
  res[6] = strings.Join(rotate(strings.Split(res[4], "/")), "/")
  res[7] = flip(res[6])
  return res
}

func getPatterns(filename string) map[string]string {
  data, err := ioutil.ReadFile(filename)
  check(err)

  patterns := make(map[string]string)
  for _, line := range strings.Split(string(data), "\n") {
    parts := strings.Split(line, " => ")
    for _, p := range permutations(parts[0]) {
      patterns[p] = parts[1]
    }
  }
  return patterns
}

func countPixels(img []string) int {
  c := 0
  for _, l := range img {
    c += strings.Count(l, "#")
  }
  return c
}

func cutZones(img []string) [][]string {
  size := 3
  if len(img) % 2 == 0 {
    size = 2
  }
  zones := make([][]string, len(img) / size)
  for i := 0; i < len(img) / size; i++ {
    zones[i] = make([]string, len(img) / size)
    for j := 0; j < len(img) / size; j++ {
      zone := make([]string, size)
      for k := 0; k < size; k++ {
        zone[k] = img[i * size + k][j * size:(j + 1) * size]
      }
      zones[i][j] = strings.Join(zone, "/")
    }
  }
  return zones
}

func rebuildImage(zones [][]string) []string {
  res := make([]string, 0)
  size := len(strings.Split(zones[0][0], "/"))
  for i := 0; i < len(zones); i++ {
    for j := 0; j < size; j++ {
      str := ""
      for k := 0; k < len(zones[i]); k++ {
        parts := strings.Split(zones[i][k], "/")
        str += parts[j]
      }
      res = append(res, str)
    }
  }
  return res
}

func solve(img []string, patterns map[string]string) int {
  for a := 0; a < 5; a++ {
    zones := cutZones(img)

    newZones := make([][]string, len(zones))
    for i := 0; i < len(zones); i++ {
      newZones[i] = make([]string, len(zones[i]))
      for j := 0; j < len(zones[i]); j++ {
        z, found := patterns[zones[i][j]]
        newZones[i][j] = z
        if !found {
          panic("Couldn't find pattern " + zones[i][j])
        }
      }
    }
    img = rebuildImage(newZones)
    fmt.Printf("Finished round %d\n", a)
  }
  return countPixels(img)
}

func main() {
  start := getImage(os.Args[1])
  patterns := getPatterns(os.Args[2])

  fmt.Println(permutations(".#./..#/###"))

  fmt.Println(solve(start, patterns))
}
