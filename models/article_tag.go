package models

import "github.com/astaxie/beego/orm"

// ArticleTag 文章标签信息
type ArticleTag struct {
	ID  int `json:"i_d" orm:"pk"`
	AID int `json:"a_i_d"`
	TID int `json:"t_i_d"`
}

func init() {
	orm.RegisterModel(new(ArticleTag))
}

// TableName 返回表名
func (*ArticleTag) TableName() string {
	return "article_tag"
}
