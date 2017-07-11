package admin

import (
	"api_web/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// AddArticleController 增加文章controller
type AddArticleController struct {
	beego.Controller
}

// Post post
func (c *AddArticleController) Post() {
	var (
		result map[string]interface{}
	)
	result = make(map[string]interface{})
	userInfo := c.GetSession("userinfo")
	result["data"] = nil
	if userInfo == nil {
		result["code"] = 2001
		result["msg"] = "登陆失效"
	} else {
		var (
			data struct {
				Title   string
				Content string
				Txt     string
				CSS     string
				Tags    []int
			}
		)
		userLogin := userInfo.(*models.UserLogin)
		user, err := models.GetUserByUID(userLogin.UID)
		if err == nil {
			json.Unmarshal(c.Ctx.Input.RequestBody, &data)
			article := models.Article{
				Title:   data.Title,
				Content: data.Content,
				Txt:     data.Txt,
				CSS:     data.CSS,
			}
			err = models.AddArticle(article, data.Tags, user)
		}
		result["code"] = 1
		result["msg"] = "添加成功"
		if err != nil {
			result["code"] = -1
			result["msg"] = "添加失败"
		}
	}
	c.Data["json"] = result
	c.ServeJSON()
}
