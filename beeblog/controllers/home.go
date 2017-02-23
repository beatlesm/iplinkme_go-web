package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	//this.Data["IsCategory"] = true
	//this.Data["IsTopic"] = true
	this.TplName = "home.html"
	//this.Ctx.WriteString(fmt.Sprint(checkAccount(this.Ctx)))
	this.Data["IsLogin"] = checkAccount(this.Ctx)




	//this.Data["TrueCond"] = true
	//this.Data["FalseCond"] = false
	//this.Data["Gmj"] = fmt.Sprint(checkAccount(this.Ctx))
}
