package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.baudu.com"
	c.Data["Email"] = "1084146423@qq.com"
	c.TplName = "index.tpl"
}
