package main

import (
	_ "HelloBeego/routers"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//定义config变量 ,接收并且赋值为全局配置变量
	config:=beego.AppConfig
	//获取配置选择
	appName := config.String("appname")
	fmt.Println("项目应用名称",appName)
	port,err :=config.Int("httpport")
	if err!=nil {
		panic("项目配置信息解析错误,请查验后重试")
	}
	fmt.Println("应用监听端口:",port)
	beego.Run()

	drive:=config.String("db_driver")
	dbUser:=config.String("db_user")
	dbPassword:=config.String("db_password")
	dbIP:=config.String("db_ip")
	dbName:=config.String("db_name")
    db,err:=sql.Open()

}

