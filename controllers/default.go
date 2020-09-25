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
	//name1 :=c.GetString("name")
	//age1,err:= c.GetInt("age")

	name:=c.Ctx.Input.Query("name")
	age:=c.Ctx.Input.Query("age")
	sex:=c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)
}
func (c *MainController) Post() {
	fmt.Println("Post类型的请求")

	user:=c.Ctx.Request.FormValue("user")
	fmt.Println("用户名为:",user)
	password:=c.Ctx.Request.FormValue("password")
	fmt.Println("用户名为:",password)

	dataBytes,err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err!=nil {
		c.Ctx.WriteString("数据接收失败，请重试")
		return
	}
	var beego models.Beego
	err = json.Unmarshal(dataBytes,&beego)
	if err!=nil {
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("用户名:",beego.Name,"年龄:",beego.Age)
	c.Ctx.WriteString("用户名是:"+beego.Name)
}