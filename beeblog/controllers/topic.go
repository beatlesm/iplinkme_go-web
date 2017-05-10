package controllers

import (
	"beeblog/models"
	"fmt"
	"github.com/astaxie/beego"
	"path"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = "true"
	this.TplName = "topic.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	//this.Data["Topics"] = topics
	topics, err := models.GetAllTopics("", "", false)
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
	label := this.Input().Get("label")
	//获取附件
	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	if fh != nil {
		//保存 attactment
		attachment = fh.Filename
		beego.Info(attachment) //print info
		//err = this.SaveToFile("attachment", attachment)// save to /
		err = this.SaveToFile("attachment", path.Join("attachment", attachment))

		if err != nil {
			beego.Error(err)
		}
	}

	// var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, category, label, content, attachment)
	} else {
		err = models.ModifyTopic(tid, title, category, label, content, attachment)
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

func (this *TopicController) Modify() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/", 302)
	}
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
	this.Data["IsLogin"] = true
}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"
	//this.Ctx.WriteString(fmt.Sprint(this.Ctx.Input.Param("0")))
	//this.Ctx.WriteString(fmt.Sprint(this.Ctx.Input.Params("0"))) //err
	//this.Ctx.WriteString("add")
	//fmt.Printf("444444\n")
	//fmt.Println(this.Ctx.Input.Param("0"))
	//fmt.Println(this.Ctx.Input.Params("0"))
	reqUrl := this.Ctx.Request.RequestURI
	i := strings.LastIndex(reqUrl, "/")
	tid := reqUrl[i+1:]
	//tid := this.Ctx.Input.Param("0")//f2: get tid
	topic, err := models.GetTopic(tid)
	//topic.Labels
	//topic, err := models.GetTopic(this.Ctx.Input.Params["0"])
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	//this.Data["Tid"] = tid
	//this.Data["Tid"] = this.Ctx.Input.Param("0")
	this.Data["Labels"] = strings.Split(topic.Labels, " ")
	replies, err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)
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
