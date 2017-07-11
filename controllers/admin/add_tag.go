package admin

import (
	"api_web/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// AddTagController add tag
type AddTagController struct {
	beego.Controller
}

// Post post
func (c *AddTagController) Post() {
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
			tag models.Tag
		)
		userLogin := userInfo.(*models.UserLogin)
		json.Unmarshal(c.Ctx.Input.RequestBody, &tag)
		addtag, err := models.AddTag(tag, userLogin.UID)
		if err == models.ErrorTagExist {
			result["code"] = 1001
			result["msg"] = "该tag已存在"

		} else if err == models.ErrorTagNull {
			result["code"] = 1002
			result["msg"] = "标签名称不能为空"
		} else if err != nil {
			result["code"] = -1
			result["msg"] = "标签添加失败"
		} else {
			result["code"] = 1
			result["msg"] = "标签添加成功"
		}
		result["data"] = addtag
	}
	c.Data["json"] = result
	c.ServeJSON()
}
