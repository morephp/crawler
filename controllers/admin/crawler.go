package admin

import (
	"github.com/astaxie/beego"
)

type CrawlerController struct {
	baseController
}

func (this *CrawlerController) Run() {

	go func() {
		beego.Info("test....")
	}()
	this.StopRun()
}
