package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	// 生成数据
	ch := make(chan struct{}, 50)
	endCh := make(chan struct{})
	defer close(ch)
	defer close(endCh)
	for i := 0; i < 1000; i++ {
		ch <- struct{}{}
		filePath := fmt.Sprintf("./demo/data/data%d.json", i)
		go CreateData(filePath, ch, endCh)
	}
	for i := 0; i < 1000; i++ {
		<-endCh
	}
	fmt.Println("write data end")

	// 读取数据
}

func GetData() {

}

func CreateData(filePath string, ch, endCh chan struct{}) {
	<-ch
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("create file err%v", err)
	}
	defer file.Close()
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
	file.Write(data)
	endCh <- struct{}{}
}
