package routers

import (
	"HelloBeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.MainController{})
    //
    
}
