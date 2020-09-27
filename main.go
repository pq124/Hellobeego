package main

import (
	"HelloBeego/db_myseq"
	_ "HelloBeego/routers"
	"github.com/astaxie/beego"
)

func main() {

//连接数据库
	db_myseq.Connect()
	beego.Run()

}

