package doc

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"redis_1/es_study/global"
	"redis_1/es_study/models"
)

//主要针对  text 字段   也可以 keyword字段

func MatchSearch() {
	limit := 3
	page := 1
	from := (page - 1) * limit

	query := elastic.NewMatchQuery("nick_name", "夜空中")
	res, err := global.ESClient.Search(models.UserModel{}.Index()).Query(query).From(from).Size(limit).Do(context.Background())
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
