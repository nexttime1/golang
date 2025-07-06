package doc

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"redis_1/es_study/global"
	"redis_1/es_study/models"
)

/*

{
  "hits": {
    "total": {          // 总匹配数的详细信息
      "value": 100,     // 总匹配数（你需要的 count）
      "relation": "eq"  // 计数方式（eq=精确计数，gte=估计值）
    },
    "hits": [           // 当前页的文档列表
      {
        "_source": "{\"id\":\"1\",\"name\":\"张三\"}"  // 文档原始数据
      },
      {
        "_source": "{\"id\":\"2\",\"name\":\"李四\"}"
      }
    ]
  }
}

*/

func DocSearch() {
	limit := 2
	page := 2
	from := (page - 1) * limit
	query := elastic.NewBoolQuery() //查全部
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
