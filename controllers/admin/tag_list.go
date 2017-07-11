package admin

import (
	"api_web/models"

	"github.com/astaxie/beego"
)

// TagListController tags list controller
type TagListController struct {
	beego.Controller
}

// Get Get
func (c *TagListController) Get() {
	var (
		result   map[string]interface{}
		pageNo   int
		pageSize int
		err      error
		tags     []*models.Tag
		total    int64
	)
	result = make(map[string]interface{})
	userinfo := c.GetSession("userinfo")
	result["data"] = nil
	if userinfo == nil {
		result["code"] = 2001
		result["msg"] = "登陆失效"
	} else {
		pageNo, _ = c.GetInt("pageNo")
		pageSize, _ = c.GetInt("pageSize")
		if pageNo == 0 {
			tags, total, err = models.TagsAll()
		} else {
			tags, total, err = models.TagList(pageNo, pageSize)
		}
		if err != nil {
			result["code"] = -1
			result["msg"] = "查询失败"
		} else {
			page := models.PageUtil(int(total), pageNo, pageSize, tags)
			result["code"] = 1
			result["msg"] = "查询成功"
			result["data"] = page
		}
	}
	c.Data["json"] = result
	c.ServeJSON()
}
