package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Tag))
}

// Tag 标签
type Tag struct {
	TID        int       `json:"t_i_d" orm:"pk;auto"`
	Name       string    `json:"name" orm:"size(20)"`
	UID        int       `json:"u_i_d"`
	CreateTime time.Time `json:"create_time"`
}

// TableName 返回表名
func (*Tag) TableName() string {
	return "tag"
}

var (
	// ErrorTagNull tag名称不能为空
	ErrorTagNull = errors.New("tag名称不能为空")
	// ErrorTagExist tag名称已存在
	ErrorTagExist = errors.New("tag名称已存在")
)

// AddTag 增加tag
func AddTag(t Tag, uid int) (*Tag, error) {
	if len(t.Name) == 0 {
		return nil, ErrorTagNull
	}
	exist := existTag(t)
	if exist {
		return nil, ErrorTagExist
	}
	o := orm.NewOrm()
	t.UID = uid
	t.CreateTime = time.Now()
	_, err := o.Insert(&t)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return &t, nil
}

// DeleteTag 删除tag
func DeleteTag(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Tag{TID: id})
	return err
}

// TagList 查询list列表
func TagList(pageNo, pageSize int) ([]*Tag, int64, error) {
	var tags []*Tag
	o := orm.NewOrm()
	tag := new(Tag)
	qs := o.QueryTable(tag)
	count, err := qs.Count()
	if err != nil {
		return nil, 0, err
	}
	offset := (pageNo - 1) * pageSize
	_, err = qs.Limit(pageSize, offset).All(&tags)
	if err != nil {
		return nil, 0, err
	}
	return tags, count, nil
}

// TagsAll 查询所有tags
func TagsAll() ([]*Tag, int64, error) {
	var tags []*Tag
	o := orm.NewOrm()
	tag := new(Tag)
	count, err := o.QueryTable(tag).All(&tags)
	if err != nil {
		return nil, 0, err
	}
	return tags, count, nil
}

func existTag(t Tag) bool {
	o := orm.NewOrm()
	tag := new(Tag)
	qs := o.QueryTable(tag)
	exist := qs.Filter("name", t.Name).Exist()
	return exist
}
