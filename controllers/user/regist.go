package user

import (
	"api_web/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// RegistController about regist
type RegistController struct {
	beego.Controller
}

// Post for intercepting post request
// @Title user regist
// @Description user regist
// @Param username formData string true "name for user"
// @Param password formData string true "password for user"
// @Success 200 {object} models.User
// @Failure 403 {object}
// @router /regist [post]
func (c *RegistController) Post() {
	var (
		u      models.User
		result map[string]interface{}
		uu     struct {
			Username string
			Password string
		}
	)
	result = make(map[string]interface{})
	json.Unmarshal(c.Ctx.Input.RequestBody, &uu)
	u.Username = uu.Username
	isExist := models.UserExists("username", u.Username)

	if isExist {
		result["code"] = 1001
		result["msg"] = "用户已存在"
	} else {
		userLogin, err := models.CreateUser(u, uu.Password)
		if err != nil {
			result["code"] = 1002
			result["msg"] = "用户注册失败"
		} else {
			c.SetSession("userinfo", userLogin)
			result["code"] = 1
			result["msg"] = "用户注册成功"
		}
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func isExitUser(userName string) bool {
	return false
}
