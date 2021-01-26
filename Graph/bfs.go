package graph

import (
	
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

//结点定义
type NodeQueue struct {
    nodes []Node
    
}

// 实现 BFS 遍历
func (g *Graph) BFS(f func(node *Node)) {

    // 初始化队列
    q := NewNodeQueue()
    // 取图的第一个节点入队列
    head := g.nodes[0]
    q.Enqueue(*head)
    // 标识节点是否已经被访问过
    visited := make(map[*Node]bool)
    visited[head] = true
    // 遍历所有节点直到队列为空
    for {
        if q.IsEmpty() {  //队列为空直接退出
            break
        }
        node := q.Dequeue()   //结点出队列
        visited[node] = true  //将该出队列结点标记为已访问true
        nexts := g.edges[*node]  
        // 将所有未访问过的邻接节点入队列
        for _, next := range nexts {
            // 如果节点已被访问过
            if visited[next] {
                continue
            }
            q.Enqueue(*next)
            visited[next] = true
        }
        // 对每个正在遍历的节点执行回调
        if f != nil {
            f(node)
        }
    }
}

// 生成节点队列
func NewNodeQueue() *NodeQueue {
    q := NodeQueue{}   
    q.nodes = []Node{}
    return &q
}

// 入队列
func (q *NodeQueue) Enqueue(n Node) {
    q.nodes = append(q.nodes, n)
}

// 出队列
func (q *NodeQueue) Dequeue() *Node {  
    node := q.nodes[0]
    q.nodes = q.nodes[1:]
    return &node
}

// 判空
func (q *NodeQueue) IsEmpty() bool {   
    return len(q.nodes) == 0
}