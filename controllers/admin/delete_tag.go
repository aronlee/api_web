package admin

import (
	"api_web/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// DeleteTagController delete tag
type DeleteTagController struct {
	beego.Controller
}

// Post post
func (c *DeleteTagController) Post() {
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
				Id int
			}
		)
		json.Unmarshal(c.Ctx.Input.RequestBody, &data)
		if data.Id == 0 {
			result["code"] = -1
			result["msg"] = "参数错误"
		} else {
			err := models.DeleteTag(data.Id)
			if err != nil {
				result["code"] = -1
				result["msg"] = "删除失败"
			} else {
				result["code"] = 1
				result["msg"] = "删除成功"
			}
		}
	}
	c.Data["json"] = result
	c.ServeJSON()
}
