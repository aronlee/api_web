package home

import (
	"api_web/models"

	"github.com/astaxie/beego"
)

// ArticleListController 文章列表
type ArticleListController struct {
	beego.Controller
}

// Get get 获取文章列表
func (c *ArticleListController) Get() {
	var (
		result map[string]interface{}
	)
	result = make(map[string]interface{})
	userinfo := c.GetSession("userinfo")
	if userinfo == nil {
		result["code"] = 2001
		result["msg"] = "登陆失效"
	} else {
		pageNo, _ := c.GetInt("pageNo")
		pageSize, _ := c.GetInt("pageSize")
		if pageNo < 1 || pageSize < 1 {
			result["code"] = -1
			result["msg"] = "[pageNo] or [pageSize] 不能小于1"
		} else {
			articles, total, err := models.HomeArticleList(pageNo, pageSize)
			if err != nil {
				beego.Error(err)
				result["code"] = -1
				result["msg"] = "查询失败"
			} else {
				page := models.PageUtil(int(total), pageNo, pageSize, articles)
				result["code"] = 1
				result["msg"] = "查询成功"
				result["data"] = page
			}
		}
	}

	c.Data["json"] = result
	c.ServeJSON()
}
