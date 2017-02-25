package controllers

import (
	"beeblog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = "true"
	this.TplName = "topic.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	//this.Data["Topics"] = topics
	topics, err := models.GetAllTopics()
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}
}
func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	title := this.Input().Get("title")
	content := this.Input().Get("content")

	var err error
	err = models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
		return
	}
	fmt.Printf("CC检查是否有del操作\n")
	this.Redirect("/topic", 302)
}
func (this *TopicController) Add() {
	this.TplName = "topic_add.html"

	//this.Ctx.WriteString("add")
}
