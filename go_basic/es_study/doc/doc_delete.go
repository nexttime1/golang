package doc

import (
	"context"
	"fmt"
	"redis_1/es_study/global"
	"redis_1/es_study/models"
)

func DocDelete() {
	deleteResponse, err := global.ESClient.Delete().
		Index(models.UserModel{}.Index()).Id("y9-Y3pcBVEns8RWUvH0y").Refresh("true"). //会立刻刷新
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(deleteResponse)
}
