package indexs

import (
	"context"
	"fmt"
	"redis_1/es_study/global"
)

func DeleteIndex(indexName string) {
	_, err := global.ESClient.DeleteIndex(indexName).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(indexName, "删除索引成功")
}
