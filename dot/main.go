package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"gonum.org/v1/gonum/graph/formats/dot"
	"gonum.org/v1/gonum/graph/formats/dot/ast"
)

func unquote(s string) string {
	return strings.Trim(s, `"`)
}

func label(attrs []*ast.Attr) string {
	for _, a := range attrs {
		if a.Key == "label" {
			return unquote(a.Val)
		}
	}
	return ""
}

func nodeID(n *ast.Node) string {
	return unquote(n.ID)
}

func dumpMap(msg string, m map[string]string) {
	fmt.Println(msg)
	for k, v := range m {
		fmt.Println(k, "=", v)
	}
}

func main() {

	file := flag.String("file", "", "input dot file to parse")
	target := flag.String("target", "", "central target (the one we're trying to find the dependencies")

	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)

	a, err := dot.Parse(r)
	if err != nil {
		panic(err)
	}
	if len(a.Graphs) != 1 {
		panic(err)
	}
	g := a.Graphs[0]
	fmt.Println(g.Directed, g.Strict, g.ID)

	// get node maps (node name -> node label and reverse)
	node2label, label2node := getNodeMaps(*g)

	dumpMap("node2label", node2label)
	dumpMap("label2node=", label2node)

	// now find all edges starting at node=*target
	for _, s := range g.Stmts {
		switch s.(type) {
		case *(ast.EdgeStmt):
			e := s.(*ast.EdgeStmt)
			f := e.From.(*ast.Node)
			if nodeID(f) == label2node[*target] {
				fmt.Println(*target, "->", node2label[e.To.Vertex.String()])
			}
		}
	}
}

func getNodeMaps(g ast.Graph) (map[string]string, map[string]string) {
	node2label := make(map[string]string)
	label2node := make(map[string]string)
	for _, s := range g.Stmts {
		switch s.(type) {
		case *(ast.NodeStmt):
			n := s.(*ast.NodeStmt)
			nid := nodeID(n.Node)
			label2node[label(n.Attrs)] = nid
			node2label[nid] = label(n.Attrs)
		}
	}

	return node2label, label2node
}
