package main

import (
	"beeblog/controllers"
	"beeblog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	//注册数据库
	models.RegisterDB()
}
func main() {
	//开启ORM调试模式
	orm.Debug = true
	//自动建表
	orm.RunSyncdb("default", false, true)
	//注册beego路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/topic", &controllers.TopicController{})

	//运行beego
	beego.Run()
}
