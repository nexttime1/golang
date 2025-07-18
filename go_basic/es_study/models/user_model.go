package models

type UserModel struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Age       int    `json:"age"`
	NickName  string `json:"nick_name"`
	CreatedAt string `json:"created_at"`
}

func (UserModel) Index() string {
	return "user_index"
}

func (UserModel) Mapping() string {
	return `{
  "mappings": {
    "properties": {
      "nick_name": { 
        "type": "text" // 查询的时候是分词匹配
      },
      "user_name": { 
        "type": "keyword" // 完整匹配
      },
      "id": {
        "type": "integer"
      },
	  "age": {
        "type": "integer"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}`

}
