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
	"api_web/controllers/home"
	"api_web/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	userNS := beego.NewNamespace("/user",
		beego.NSRouter("/login", &user.LoginController{}),
		beego.NSRouter("/regist", &user.RegistController{}),
	)
	adminNS := beego.NewNamespace("/admin",
		beego.NSRouter("/tagList", &admin.TagListController{}),
		beego.NSRouter("/addTag", &admin.AddTagController{}),
		beego.NSRouter("/deleteTag", &admin.DeleteTagController{}),
		beego.NSRouter("/addArticle", &admin.AddArticleController{}),
		beego.NSRouter("/articleList", &admin.ArticleListController{}),
	)
	homeNS := beego.NewNamespace("/home",
		beego.NSRouter("/articleList", &home.ArticleListController{}),
		beego.NSRouter("/articleDetail", &home.ArticleDetailController{}),
	)

	beego.AddNamespace(userNS)
	beego.AddNamespace(adminNS)
	beego.AddNamespace(homeNS)
}
