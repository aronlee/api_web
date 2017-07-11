package models

import "time"
import "github.com/astaxie/beego/orm"

// Comment 评论
type Comment struct {
	CID        int       `json:"c_i_d" orm:"pk;auto"`
	Floor      int       `json:"floor"`
	Content    string    `json:"content"`
	ArtID      int       `json:"art_i_d"`
	CreateTime time.Time `json:"create_time"`
	User       *User     `json:"user" orm:"rel(fk)"` //设置一对多关系
}

func init() {
	orm.RegisterModel(new(Comment))
}

// TableName 返回表名
func (*Comment) TableName() string {
	return "comment"
}
