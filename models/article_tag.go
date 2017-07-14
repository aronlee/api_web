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

// GetTagsByArticle 根据articleId获取所有的article tag 列数据
func GetTagsByArticle(aid int) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	o := orm.NewOrm()
	qs := o.QueryTable(new(ArticleTag))
	_, err := qs.Filter("a_i_d", aid).All(&articleTags)
	if err != nil {
		return nil, err
	}
	return articleTags, nil
}
