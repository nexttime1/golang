package doc

import (
	"context"
	"fmt"
	"redis_1/es_study/global"
	"redis_1/es_study/models"
	"time"
)

func DocCreate() {
	user := models.UserModel{
		ID:        15,
		UserName:  "薛提猛",
		Age:       22,
		NickName:  "夜空中最闪亮的小星星",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	//第一个Index 是选择要做关于index的命令  第二个是操作哪一个index
	indexResponse, err := global.ESClient.Index().Index(user.Index()).BodyJson(&user).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v \n", indexResponse)

}
