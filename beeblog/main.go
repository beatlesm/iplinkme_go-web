package main

import (
	"beeblog/controllers"
	"beeblog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
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
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	//创建附件目录
	os.Mkdir("attachment", os.ModePerm)
	//作为控制器
	beego.Router("/attachment/:all", &controllers.AttachController{})
	/*	//作为静态文件
		beego.SetStaticPath("/attachment", "attachment")*/
	//运行beego
	beego.Run()
}
