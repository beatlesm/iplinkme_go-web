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
	topics, err := models.GetAllTopics(false)
	//err, topics := models.GetAllTopics(false)//err
	if err != nil {
		beego.Error(err)
		//beego.Error(err.Error)
	} else {
		this.Data["Topics"] = topics
	}
}
func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	// 解析表单
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, category, content)
	} else {
		err = models.ModifyTopic(tid, title, category, content)
	}
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
func (this *TopicController) View() {
	this.TplName = "topic_view.html"
	//this.Ctx.WriteString(fmt.Sprint(this.Ctx.Input.Param("0")))
	//this.Ctx.WriteString(fmt.Sprint(this.Ctx.Input.Params("0"))) //err
	//this.Ctx.WriteString("add")
	//fmt.Printf("444444\n")
	//fmt.Println(this.Ctx.Input.Param("0"))
	//fmt.Println(this.Ctx.Input.Params("0"))

	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	//topic, err := models.GetTopic(this.Ctx.Input.Params["0"])
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	//this.Data["Tid"] = tid
	this.Data["Tid"] = this.Ctx.Input.Param("0")

}

func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
}
func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	/*
		//方法1:
		err := models.DeleteTopic(this.Input().Get("tid"))
		HTML:
		<th>
			<a href="/topic/delete?tid={{.Id}}">删除</a>
		</th>
	*/
	//方法2:
	err := models.DeleteTopic(this.Ctx.Input.Param("0"))
	//<th>
	//<a href="/topic/delete/{{.Id}}">删除</a>
	//</th>
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/", 302)
}
