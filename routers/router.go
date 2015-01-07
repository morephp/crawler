package routers

import (
	"github.com/astaxie/beego"
	"github.com/morephp/crawler/controllers"
	"github.com/morephp/crawler/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")

	beego.Router("/admin", &admin.LoginController{}, "*:Index")

	beego.Router("/admin/main", &admin.IndexController{}, "*:Main")
	beego.Router("/admin/info", &admin.IndexController{}, "*:Info")

	beego.Router("/admin/user", &admin.UserController{}, "*:Index")
	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
	beego.Router("/admin/user/update", &admin.UserController{}, "*:Update")
	beego.Router("/admin/user/del", &admin.UserController{}, "*:Del")

	beego.Router("/admin/config", &admin.ConfigController{}, "*:Index")
	beego.Router("/admin/help", &admin.HelpController{}, "*:Index")

	beego.Router("/admin/login", &admin.LoginController{}, "post:Login")
	beego.Router("/admin/logout", &admin.LoginController{}, "*:Logout")

}
