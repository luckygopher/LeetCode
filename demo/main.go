package main

import (
	"bufio"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

// 用户信息数据结构
type User struct {
	// 用户id
	Id int64 `json:"id"`
	// 用户姓名
	Name string `json:"name"`
	// 父母是否参与火星研究
	ParentMarsResearch bool `json:"parent_mars_research"`
	// 总年纳税(单位:w)
	AnnualTax int64 `json:"annual_tax"`
	// 是否有犯罪记录
	CriminalRecord bool `json:"criminal_record"`
	// 是否有亲兄弟姐妹
	BrothersAndSisters bool `json:"brothers_and_sisters"`
	// 是否工程师
	Engineer bool `json:"engineer"`
	// 报名时间
	Time int64 `json:"time"`
}

// 思路
// 1、构造用户数据，10亿数据拆分为1000个小文件存储，一个文件100万，每行一条数据
// 2、读取文件，计算积分并构建1000个节点的小顶堆
// 3、弹出最低位，得出每个文件的前1000
// 4、将每个文件的1000进行排序

func main() {
	// 生成数据
	// ch := make(chan struct{}, 100)
	// endCh := make(chan struct{}, 100)
	// defer close(ch)
	// defer close(endCh)
	// for i := 0; i < 1000; i++ {
	// 	ch <- struct{}{}
	// 	filePath := fmt.Sprintf("./demo/data/data%d.json", i)
	// 	go WriteData(filePath, ch, endCh)
	// }
	// for i := 0; i < 1000; i++ {
	// 	<-endCh
	// }
	// fmt.Println("write data end")

	// 读取数据

}

func WriteData(filePath string, ch, endCh chan struct{}) {
	<-ch
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("create file err%v", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	randData := []bool{true, false}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	user := make([]*User, 0, 1000000)
	for j := 0; j < 1000000; j++ {
		user = append(user, &User{
			Id:                 r.Int63(),
			Name:               "aaa",
			ParentMarsResearch: randData[r.Intn(2)],
			AnnualTax:          r.Int63n(2000000),
			CriminalRecord:     randData[r.Intn(2)],
			BrothersAndSisters: randData[r.Intn(2)],
			Engineer:           randData[r.Intn(2)],
			Time:               time.Now().Unix(),
		})
	}
	data, _ := json.Marshal(user)
	writer.Write(data)
	writer.Flush()
	endCh <- struct{}{}
}

// 小顶堆
type Node struct {
	Value int
}

// 用于构建结构体切片为最小堆，需要调用down函数
func Init(nodes []Node) {
	for i := len(nodes)/2 - 1; i >= 0; i-- {
		down(nodes, i, len(nodes))
	}
}

// 需要down(下沉)的元素在切片中的索引为i，n为heap的长度，将该元素下沉到该元素
// 对应的子树合适的位置，从而满足该子树为最小堆的要求
func down(nodes []Node, i, n int) {
	parent := i
	child := 2*parent + 1
	temp := nodes[parent].Value
	for {
		if child < n {
			if child+1 < n && nodes[child].Value > nodes[child+1].Value {
				child++
			}
			if temp <= nodes[child].Value {
				break
			}
			nodes[parent].Value = nodes[child].Value
			parent = child
			child = child*2 + 1
		} else {
			break
		}
	}
	nodes[parent].Value = temp
}

// 用于保证插入新元素(j为元素的索引，切片末尾插入，堆低插入)的结构体切片之后仍然是一个最小堆
func up(nodes []Node, j int) {
	child := j
	parent := (j - 1) / 2
	for {
		if child == 0 {
			break
		}
		if nodes[parent].Value < nodes[child].Value {
			break
		}
		temp := nodes[child].Value
		nodes[child].Value = nodes[parent].Value
		nodes[parent].Value = temp
		child = parent
		parent = (parent - 1) / 2
	}
}

// 弹出最小元素，并保证弹出后的结构体切片仍然是一个最小堆，第一个返回值是弹出的节点的信息，
// 第二个参数是Pop操作后得到的新的结构体切片
func Pop(nodes []Node) (Node, []Node) {
	min := nodes[0]
	nodes[0].Value = nodes[len(nodes)-1].Value
	nodes = nodes[:len(nodes)-1]
	down(nodes, 0, len(nodes)-1)
	return min, nodes
}

// 保证插入新元素时，结构体切片仍然是一个最小堆，需要调用up函数
func Push(node Node, nodes []Node) []Node {
	nodes = append(nodes, node)
	up(nodes, len(nodes)-1)
	return nodes
}

// 移除切片中指定索引的元素，保证移除后结构体切片仍然是一个最小堆
func Remove(nodes []Node, node Node) []Node {
	for i := 0; i < len(nodes); i++ {
		if node.Value == nodes[i].Value {
			nodes[i].Value = nodes[len(nodes)-1].Value
			nodes = nodes[0 : len(nodes)-1]
			down(nodes, 0, len(nodes)-1)
			break
		}
	}
	return nodes
}
