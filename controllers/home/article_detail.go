package home

import (
	"api_web/models"

	"github.com/astaxie/beego"
)

// ArticleDetailController 查询文章详情
type ArticleDetailController struct {
	beego.Controller
}

// Get 查询文章详情
func (c *ArticleDetailController) Get() {
	var (
		result map[string]interface{}
	)
	result = make(map[string]interface{})

	id, _ := c.GetInt("id")
	article, err := models.ArticleDetail(id)
	if err != nil {
		result["code"] = -1
		result["msg"] = "查询失败"
	} else {
		result["code"] = 1
		result["msg"] = "查询成功"
		result["data"] = article
	}
	c.Data["json"] = result
	c.ServeJSON()
}
