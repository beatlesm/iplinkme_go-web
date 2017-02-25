package controllers

import (
	//"fmt"
	//"net/url"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	//fmt.Printf("222222\n")

	isExit := this.Input().Get("ex") == "true"
	//fmt.Printf(fmt.Sprint(isExit))
	//fmt.Printf("3\n")
	//this.Ctx.WriteString(fmt.Sprint(isExit))
	if isExit {
		//del Cookie
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 301)
		//fmt.Printf("444444\n")
		return
	}
	this.TplName = "login.html"
}
func (this *LoginController) Post() {
	//需要将this.Input()转换成字符串
	//this.Ctx.WriteString(this.Input())//err:cannot use this.Controller.Input() (type url.Values) as type string in argument to this.Controller.Ctx.WriteString

	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	//this.Ctx.WriteString("hello world")
	//get uname pwd autologin
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autologin := this.Input().Get("autologin") == "on"
	//fmt.Printf("5555555555\n")
	//fmt.Printf(fmt.Sprint(this.Input().Get("exit")))
	//this.Ctx.WriteString("<br>")
	//this.Ctx.WriteString(uname)
	//this.Ctx.WriteString(pwd)
	//this.Ctx.WriteString(fmt.Sprint(autologin))

	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autologin {
			maxAge = 1<<31 - 1
		}
		//fmt.Printf("5555555555\n")
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	//fmt.Printf(" 666\n")
	this.Redirect("/", 301)
	//this.Redirect("/login", 301)

	//return
}

func checkAccount(ctx *context.Context) bool {
	//	fmt.Printf("111111111111\n")

	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	varExit := beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd
	//fmt.Printf(fmt.Sprint(varExit))
	//fmt.Printf(" 111\n")
	return varExit
}
