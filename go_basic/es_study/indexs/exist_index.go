package indexs

import (
	"context"
	"redis_1/es_study/global"
)

func ExistIndex(indexName string) bool {
	exist, _ := global.ESClient.IndexExists(indexName).Do(context.Background())
	return exist
}
