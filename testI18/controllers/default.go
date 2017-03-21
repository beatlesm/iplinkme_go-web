package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type baseController struct {
	beego.Controller
	i18n.Locale
}

func (this *baseController) Prepare() {
	lang := this.GetString("lang")

	if lang == "zh-CN" {
		this.Lang = lang
	} else {
		this.Lang = "en-US"
	}

	/*	if lang == "en-US" {
			this.Lang = lang
		} else {
			this.Lang = "zh-CN"
		}*/
	this.Data["Lang"] = this.Lang
}

type MainController struct {
	baseController
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["About"] = "about"
	this.Data["Hi"] = "hi"
	this.Data["Bye"] = "bye"
	this.TplName = "index.tpl"
}
