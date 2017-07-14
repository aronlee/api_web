package models

import "time"
import "github.com/astaxie/beego/orm"
import "github.com/astaxie/beego"

func init() {
	orm.RegisterModel(new(Article))
}

// Article 文章
type Article struct {
	AID        int       `json:"a_i_d" orm:"pk;auto"`
	UID        int       `json:"uid"`
	Title      string    `json:"title" orm:"size(127)"`
	Content    string    `json:"content"`
	Txt        string    `json:"txt"`
	CSS        string    `json:"c_s_s"`
	ViewNum    int       `json:"view_num"`
	CommentNum int       `json:"comment_num"`
	LikeNum    int       `json:"like_num"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	// Tags       []*Tag     `json:"tags" orm:"rel(m2m)"`          // 设置多对多的关系
	// Comments []*Comment `json:"comments" orm:"reverse(many)"` // 设置一对多的反向关系
}

// TableName 返回表名
func (*Article) TableName() string {
	return "article"
}

// AddArticle 增加文章
func AddArticle(article Article, tagIds []int, user *User) error {
	var (
		artTags []ArticleTag
		aid     int64
	)

	article.UID = user.UID
	article.CreateTime = time.Now()
	article.UpdateTime = time.Now()

	o := orm.NewOrm()
	err := o.Begin()
	aid, err = o.Insert(&article)

	for _, v := range tagIds {
		artTags = append(artTags, ArticleTag{
			AID: int(aid),
			TID: v,
		})
	}

	if err != nil {
		beego.Error(err)
	}

	if len(tagIds) > 0 {
		_, err = o.InsertMulti(len(tagIds), artTags)
	}

	if err != nil {
		beego.Error(err)
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

// ArticleList 分页查询 article
func ArticleList(pageNo, pageSize int) ([]*Article, int64, error) {
	var (
		articles []*Article
	)
	offset := (pageNo - 1) * pageSize
	o := orm.NewOrm()
	qs := o.QueryTable(new(Article))
	count, err := qs.Count()
	if err != nil {
		return nil, 0, err
	}
	_, err = qs.Limit(pageSize, offset).All(&articles)
	if err != nil {
		return nil, 0, err
	}
	return articles, count, nil
}
