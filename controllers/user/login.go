package user

import (
	"api_web/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

// LoginController about login
type LoginController struct {
	beego.Controller
}

// Post for intercepting post request
// @Title user login
// @Description user login
// @Param username formData string true "name for user"
// @Param password formData string true "password for user"
// @Success 200 {object} models.User
// @Failure 403 {object}
// @router /login [post]
func (c *LoginController) Post() {
	var (
		result map[string]interface{}
		user   struct {
			Username string
			Password string
		}
	)
	result = make(map[string]interface{})
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	userLogin, err := models.Login(user.Username, user.Password)
	if err == models.ErrorName {
		result["code"] = 1001
		result["msg"] = "用户名不存在"
	} else if err == models.ErrorPassword {
		result["code"] = 1002
		result["msg"] = "密码错误"
	} else if err != nil {
		result["code"] = -1
		result["msg"] = "登陆失败"
	} else {
		c.SetSession("userinfo", userLogin)
		fmt.Println(userLogin)
		result["code"] = 1
		result["msg"] = "登陆成功"
	}
	c.Data["json"] = result
	c.ServeJSON()
}
