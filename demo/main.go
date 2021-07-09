package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
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
	// 积分
	Score int64 `json:"score"`
	// 父母是否参与火星研究
	ParentMarsResearch bool `json:"parent_mars_research"`
	// 总年纳税
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
	// ch := make(chan struct{}, 10)
	// endCh := make(chan struct{}, 10)
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
	sortData := make([]*User, 0, 1000)
	for i := 0; i < 1; i++ {
		filePath := fmt.Sprintf("./demo/data/data%d.json", i)
		ReadData(filePath, sortData)
	}
	fmt.Println(sortData)
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
	for j := 0; j < 1000000; j++ {
		data, _ := json.Marshal(User{
			Id:                 r.Int63(),
			Name:               "aaa",
			Score:              0,
			ParentMarsResearch: randData[r.Intn(2)],
			AnnualTax:          r.Int63n(2000000),
			CriminalRecord:     randData[r.Intn(2)],
			BrothersAndSisters: randData[r.Intn(2)],
			Engineer:           randData[r.Intn(2)],
			Time:               time.Now().Unix(),
		})
		dataStr := fmt.Sprintf("%s\n", string(data))
		writer.WriteString(dataStr)
	}
	writer.Flush()
	endCh <- struct{}{}
}

func ReadData(filePath string, sortData []*User) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("read file err%v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		userInfo := new(User)
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		json.Unmarshal([]byte(str), userInfo)
		// 计算积分
		var score int64
		if userInfo.ParentMarsResearch {
			score += 500
		}
		if userInfo.Engineer {
			score += 200
		}
		if userInfo.AnnualTax > 1000000 && userInfo.CriminalRecord && userInfo.BrothersAndSisters {
			score += 300
		}
		userInfo.Score = score
		if len(sortData) < 1000 {
			sortData = append(sortData, userInfo)
		}
		InsertSort(sortData, userInfo)
	}
}

//插入排序
func InsertSort(data []*User, user *User) {
	for i := 1; i <= len(data); i++ {
		insert := user
		j := i
		for j > 0 && (data[j-1].Score < insert.Score || (data[j-1].Score == insert.Score && data[j-1].Time > insert.Time)) {
			if j >= 1000 {
				j--
				continue
			}

			data[j] = data[j-1]
			j--
		}
		if j < 1000 {
			data[j] = insert
		}
	}
}
