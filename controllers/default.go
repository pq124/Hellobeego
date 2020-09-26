package controllers

import (
	"HelloBeego/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.baudu.com"
	c.Data["Email"] = "1084146423@qq.com"
	c.TplName = "index.tpl"

}
func (c *MainController) Post(){
	//body := c.Ctx.Request.Body
	dataByes, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收失败,请重试")
		return
	}
	//json包解析
	var person models.Person
	err  = json.Unmarshal(dataByes,&person)
	if err != nil {
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("用户名：",person.User,"，地址：",person.Address)
	c.Ctx.WriteString("用户名是:"+person.User)

}