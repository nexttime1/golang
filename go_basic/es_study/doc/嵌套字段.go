package doc

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"redis_1/es_study/global"
	"redis_1/es_study/models"
)

/*
"title": {
    "type": "text",
    "fields": {
        "keyword": {
            "type": "keyword",
            "ignore_above": 256
        }
    }
},

*/

func Nested() {
	//可以模糊 可以精确
	limit := 4
	page := 1
	from := (page - 1) * limit
	//模糊
	query := elastic.NewMatchQuery("title", "夜空中最亮的星星")
	//精确
	termQuery := elastic.NewTermQuery("title.keyword", "夜空中")
	res, err := global.ESClient.Search(models.UserModel{}.Index()).Query(query).From(from).Size(limit).Do(context.Background())
	global.ESClient.Search(models.UserModel{}.Index()).Query(termQuery).From(from).Size(limit).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	count := res.Hits.TotalHits.Value
	fmt.Println(count)
	for _, item := range res.Hits.Hits {
		fmt.Println(string(item.Source))
	}
}
