package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "regexp"
  "strconv"
  "strings"
)

var nodeExp = regexp.MustCompile(`(?P<name>\d+) <-> (?P<neighbours>(\d+)(, \d+)*)`)

type Node struct {
  name int
  neighbours []*Node
}

func NewNode(name int) *Node {
  n := Node{name: name}
  n.neighbours = make([]*Node, 0)
  return &n
}

func (n *Node) AddEdge(o *Node) {
  n.neighbours = append(n.neighbours, o)
  o.neighbours = append(o.neighbours, n)
}

func (n *Node) Visit(visited map[int]struct{}) {
  _, seen := visited[n.name]
  if seen {
    return
  }

  visited[n.name] = struct{}{}
  for _, neighbour := range n.neighbours {
    neighbour.Visit(visited)
  }
}

func FindOrCreateNode(nodes map[int]*Node, name int) *Node {
  n := nodes[name]
  if n == nil {
    n = NewNode(name)
    nodes[name] = n
  }
  return n
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func readLines(fh *bufio.Reader, c chan string) {
  for {
    line, err := fh.ReadString('\n')
    c <- strings.TrimSuffix(line, "\n")
    if err == io.EOF {
      break
    }
  }
  close(c)
}

func parseLine(line string) (int, []int) {
  match := nodeExp.FindStringSubmatch(line)
  name, _ := strconv.Atoi(match[1])
  neighbours := make([]int, 0)
  for _, n := range strings.Split(match[2], ", ") {
    name, _ := strconv.Atoi(n)
    neighbours = append(neighbours, name)
  }
  return name, neighbours
}

func buildGraph(c chan string) map[int]*Node {
  nodes := make(map[int]*Node)
  for line := range c {
    name, neighbours := parseLine(line)
    node := FindOrCreateNode(nodes, name)
    for _, n := range neighbours {
      neighbour := FindOrCreateNode(nodes, n)
      node.AddEdge(neighbour)
    }
  }
  return nodes
}

func countGroups(nodes map[int]*Node) int {
  group := 0
  visited := make(map[int]struct{})
  for _, node := range nodes {
    _, seen := visited[node.name]
    if !seen {
      node.Visit(visited)
      group += 1
    }
  }
  return group
}

func main() {
  fh, err := os.Open(os.Args[1])
  check(err)
  defer fh.Close()

  reader := bufio.NewReader(fh)

  c := make(chan string, 100)

  go readLines(reader, c)
  nodes := buildGraph(c)
  fmt.Println(countGroups(nodes))
}
