package user

import (
	"github.com/astaxie/beego"
)

// OauthController c
type OauthController struct {
	beego.Controller
}

// Get get
func (c *OauthController) Get() {
	var (
		result map[string]interface{}
	)
	result = make(map[string]interface{})
	c.Data["json"] = result
	c.ServeJSON()
}
