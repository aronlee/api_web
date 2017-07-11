// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api_web/controllers/admin"
	"api_web/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/user",
		beego.NSRouter("/login", &user.LoginController{}),
		beego.NSRouter("/regist", &user.RegistController{}),
	)
	admin := beego.NewNamespace("/admin",
		beego.NSRouter("/tagList", &admin.TagListController{}),
		beego.NSRouter("/addTag", &admin.AddTagController{}),
		beego.NSRouter("/deleteTag", &admin.DeleteTagController{}),
		beego.NSRouter("/addArticle", &admin.AddArticleController{}),
	)

	beego.AddNamespace(ns)
	beego.AddNamespace(admin)
}
