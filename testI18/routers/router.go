package routers

import (
	"github.com/astaxie/beego"
	"testI18/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
