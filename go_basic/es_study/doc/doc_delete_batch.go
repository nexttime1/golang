package doc

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"redis_1/es_study/global"
	"redis_1/es_study/models"
)

func DeleteBatch() {
	IdList := []string{
		"zN-n3pcBVEns8RWUJ326",
		"zd-n3pcBVEns8RWUoX3C",
	}
	//借助 Elasticsearch 客户端来初始化一个批量操作请求  桶  指定哪一个 index
	bulk := global.ESClient.Bulk().Index(models.UserModel{}.Index()).Refresh("true")
	//遍历 ID 列表并添加删除请求
	for _, id := range IdList {
		request := elastic.NewBulkDeleteRequest().Id(id) //删除单一文档的请求
		bulk.Add(request)

	}
	//执行前面构建的批量删除操作
	res, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Succeeded()) //删除的那些   是个切片

}
