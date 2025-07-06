package indexs

import (
	"context"
	"fmt"
	"redis_1/es_study/global"
	"redis_1/es_study/models"
)

func CreateIndex() {
	index := "user_index"
	exist := ExistIndex(index)
	if exist {
		fmt.Println("索引已经存在 现在执行删除 重建")
		//存在了   先删除  再创建
		DeleteIndex(index)

	}

	createindex, err := global.ESClient.
		CreateIndex(index).
		BodyString(models.UserModel{}.Mapping()).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(createindex)
	fmt.Println("索引创建成功")
}
