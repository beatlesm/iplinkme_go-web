package controllers

import (
	"beeblog/models"
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
	//this.Data["Topics"] = topics
	//topics, err := models.GetAllTopics(false)
	topics, err := models.GetAllTopics(this.Input().Get("cate"), this.Input().Get("label"), true)
	/*	n := len(topics)
		fmt.Printf("CC  文章数 len =%d\n", n)
	*/ //err, topics := models.GetAllTopics(false)//err
	if err != nil {
		beego.Error(err)
		//beego.Error(err.Error)
	} else {
		this.Data["Topics"] = topics
	}

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories
	//this.Data["TrueCond"] = true
	//this.Data["FalseCond"] = false
	//this.Data["Gmj"] = fmt.Sprint(checkAccount(this.Ctx))
}
