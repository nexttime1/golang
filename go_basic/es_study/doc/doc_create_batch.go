package doc

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"redis_1/es_study/global"
	"redis_1/es_study/models"
	"time"
)

func CreateBatch() {
	ModelList := []models.UserModel{
		{
			ID:        16,
			UserName:  "李白",
			Age:       20,
			NickName:  "夜空中最闪亮的李白",
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			ID:        12,
			UserName:  "丽萨",
			Age:       18,
			NickName:  "夜空中最闪亮的丽萨",
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	bulk := global.ESClient.Bulk().Index(models.UserModel{}.Index()).Refresh("true")
	for _, model := range ModelList {
		request := elastic.NewBulkCreateRequest().Doc(model)
		bulk = bulk.Add(request)
	}
	res, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", res)

}
