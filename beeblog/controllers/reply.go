package controllers

import (
	"beeblog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add() {
	tid := this.Input().Get("tid") // get tid form html
	fmt.Printf("5555555555\n")
	err := models.AddReply(tid, this.Input().Get("nickname"), this.Input().Get("content"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic/view/"+tid, 302)
}

func (this *ReplyController) Delete() {
	//checkAccount
	if !checkAccount(this.Ctx) {
		return
	}
	tid := this.Input().Get("tid")
	err := models.DeleteReply(this.Input().Get("rid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic/view/"+tid, 302)
}
