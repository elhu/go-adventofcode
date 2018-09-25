package main

import "os"
import "bufio"
import "strings"
import "io"
import "fmt"
import "regexp"
import "strconv"

var nodeExp = regexp.MustCompile(`(?P<name>\w+) \((?P<weight>\d+)\)( -> (?P<children>.+))?`)

type Node struct {
  weight int
  fullWeight int
  name string
  children []*Node
  parent *Node
}

func NewNode(weight int, name string) *Node {
  node := Node{}
  node.weight = weight
  node.fullWeight = 0
  node.name = name
  node.children = make([]*Node, 0)

  return &node
}

func (n *Node) Print() {
  fmt.Printf("----\n%s (%d) -> %d:\n", n.name, n.weight, n.fullWeight)
  for _ ,child := range n.children {
    fmt.Printf("%s (%d) -> %d:\n", child.name, child.weight, child.fullWeight)
    // child.Print()
  }
}

func (n *Node) ComputeWeights() int {
  n.fullWeight = n.weight
  for _ ,child := range n.children {
    n.fullWeight += child.ComputeWeights()
  }
  return n.fullWeight
}

func (n *Node) Balanced() bool {
  for i := 1; i < len(n.children); i++ {
    if n.children[i - 1].weight != n.children[i].weight {
      return false
    }
  }
  return true
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

func parseLine(line string) (string, int, []string) {
  match := nodeExp.FindStringSubmatch(line)
  name := match[1]
  weight, _ := strconv.Atoi(match[2])
  children := strings.Split(match[4], ", ")
  return name, weight, children
}

func findRoot(node *Node) *Node {
  if node.parent != nil {
    return findRoot(node.parent)
  }
  return node
}

func buildTree(c chan string) *Node {
  seen := make(map[string]*Node)
  var lastNode *Node
  for line := range c {
    name, weight, childrenNames := parseLine(line)
    node, found := seen[name]
    if !found {
      node = NewNode(weight, name)
      seen[name] = node
    }
    node.weight = weight
    for _, childName := range childrenNames {
      if childName != "" {
        child, found := seen[childName]
        if !found {
          child = NewNode(0, childName)
        }
        child.parent = node
        seen[childName] = child
        node.children = append(node.children, child)
      }
    }
    lastNode = node
  }
  root := findRoot(lastNode)
  root.ComputeWeights()
  return root
}

func findImbalance(root *Node) *Node {
  for _, child := range root.children {
    if !child.Balanced() {
      return findImbalance(child).parent
    }
  }
  return root
}

func solve(root *Node) int {
  diff := 0
  for i := 1; i < len(root.children); i++ {
    diff = root.children[i].fullWeight - root.children[i - 1].fullWeight
    if diff != 0 {
      return root.children[i].fullWeight + diff
    }
  }
  return diff
}

func main() {
  fh, err := os.Open(os.Args[1])
  check(err)
  defer fh.Close()

  reader := bufio.NewReader(fh)

  c := make(chan string, 100)

  go readLines(reader, c)
  root := buildTree(c)
  fmt.Println(solve(findImbalance(root)))
}
