package main

import "fmt"

type ArticleClass struct {
	Id int `json:"id"`
	ParentId int `json:"parent_id"`
	Name int `json:"name"`
	List []*ArticleClass `json:"list,omitempty"`
}

func buildData(list []*ArticleClass) map[int]map[int]*ArticleClass  {
	var data map[int]map[int]*ArticleClass = make(map[int]map[int]*ArticleClass)
	for _,v:=range list {
		id := v.Id
		fid := v.ParentId
		if _,ok:=data[fid];!ok{
			data[fid] = make(map[int]*ArticleClass)
		}
		data[fid][id] = v
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
	var list []*ArticleClass
	data := buildData(list)
	result := makeTreeCore(0, data)
	fmt.Println(result)
}