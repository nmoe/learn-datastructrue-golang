package main

import (
	"fmt"
	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

//Node struct  define
//结点定义
type Node struct {
	Value Item
}

//Graph define
//图的定义
type Graph struct {
	nodes []*Node          //结点集合,是一个切片类型,为指向Node结构体的指针
	edges map[Node][]*Node //邻接表表示的无向图
}


//
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

//AddNode method
//添加结点
//append用来将元素添加到切片末尾并返回结果,append(切片类型,其他的元素)
func (g *Graph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

//AddEdge method
//增加边
func (g *Graph) AddEdge(u, v *Node) {
	//不存在图时,创建一个图
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*u] = append(g.edges[*u], v) //创建u->v的边
	g.edges[*v] = append(g.edges[*v], u) //无向图,也存在v->u的边
}

//格式化 输出图
func (g *Graph) String() {
	str := ""
	for _, node := range g.nodes {
		str += node.String() + "->"
		nexts := g.edges[*node]
		for _, next := range nexts {
			str += next.String() + " "
		}
		str += "\n"
	}
	fmt.Println(str)
}

func main() {
	g := Graph{}
	n1, n2, n3, n4, n5 := Node{Value:1}, Node{Value:2}, Node{Value:3}, Node{Value:4}, Node{Value:5}

	g.AddNode(&n1)
	g.AddNode(&n2)
	g.AddNode(&n3)
	g.AddNode(&n4)
	g.AddNode(&n5)

	g.AddEdge(&n1, &n2)
	g.AddEdge(&n1, &n5)
	g.AddEdge(&n2, &n3)
	g.AddEdge(&n2, &n4)
	g.AddEdge(&n2, &n5)
	g.AddEdge(&n3, &n4)
	g.AddEdge(&n4, &n5)

	g.String()
}
