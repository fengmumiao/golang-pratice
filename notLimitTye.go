package main

import (
	"fmt"
	"encoding/json"
)

type ArticleClass struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ParentId int `json:"pid"`
	List []*ArticleClass `json:"list,omitempty"`
}



func buildData(list []*ArticleClass) map[int]map[int]*ArticleClass  {
	var data map[int]map[int]*ArticleClass = make(map[int]map[int]*ArticleClass)
	for _,v:=range list {
		id := v.Id
		pid := v.ParentId
		if _,ok:=data[pid];!ok{
			data[pid] = make(map[int]*ArticleClass)
		}
		data[pid][id] = v
	}
	return data
}

func makeTreeCore(index int, data map[int]map[int]*ArticleClass) []*ArticleClass  {
	tmp := make([]*ArticleClass, 0)
	for id,item:= range data[index]{
		if data[id] != nil{
			item.List=makeTreeCore(id, data)
		}
		tmp=append(tmp,item)
	}
	return tmp
}

func main()  {
	var sByte =[]byte(`[{"id":1,"name":"\u7535\u8111","pid":0},{"id":2,"name":"\u624b\u673a","pid":0},{"id":3,"name":"\u7b14\u8bb0\u672c","pid":1},{"id":4,"name":"\u53f0\u5f0f\u673a","pid":1},{"id":5,"name":"\u667a\u80fd\u673a","pid":2},{"id":6,"name":"\u529f\u80fd\u673a","pid":2},{"id":7,"name":"\u8d85\u7ea7\u672c","pid":3},{"id":8,"name":"\u6e38\u620f\u672c","pid":3}]`)
	//var sByte =[]byte(`{"id":1,"name":"\u7535\u8111","pid":0}`)
	var dianqi = new([]*ArticleClass)
	err := json.Unmarshal(sByte, &dianqi)
	if err != nil {
		fmt.Errorf("Can not decode data: %v\n", err)
	}
	fmt.Printf("%v\n", dianqi)
	var list []*ArticleClass
	data := buildData(list)
	result := makeTreeCore(0, data)
	fmt.Println(result)
}