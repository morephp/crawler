package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Index() {
	this.Redirect("/admin", 302)
	this.StopRun()
}
