package core

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"redis_1/es_study/global"
)

func EsConnect() {
	//创建客户端实例   把嗅探功能关闭
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetBasicAuth("elastic", "123456"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client)
	global.ESClient = client
}
