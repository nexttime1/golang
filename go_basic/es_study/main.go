package main

import (
	"redis_1/es_study/core"
	"redis_1/es_study/doc"
)

func main() {
	core.EsConnect()
	//indexs.CreateIndex()

	//doc.DocCreate()
	//单个删除  根据id删除   没有会报错
	//doc.DocDelete()

	//删除多个   没有不会报错
	//doc.DeleteBatch()
	//添加多个
	//doc.CreateBatch()
	//查询
	//doc.DocSearch()

	//精确匹配
	//doc.ExactSearch()

	//模糊匹配
	//doc.MatchSearch()

	//修改
	doc.DocUpdate()
}
