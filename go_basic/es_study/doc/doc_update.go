package doc

import (
	"context"
	"fmt"
	"redis_1/es_study/global"
	"redis_1/es_study/models"
)

func DocUpdate() {
	res, err := global.ESClient.Update().Index(models.UserModel{}.Index()).Id("0N_D3pcBVEns8RWU233b").Doc(map[string]interface{}{
		"user_name": "宋可为",
	}).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}
