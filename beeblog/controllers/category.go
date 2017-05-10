package controllers

import (
	"beeblog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

// 检查是否有操作
func (this *CategoryController) Get() {
	//fmt.Printf("CC检查是否有操作\n")
	// 检查是否有操作
	op := this.Input().Get("op")
	switch op {
	case "add":
		name := this.Input().Get("name")
		fmt.Println(name)
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		//fmt.Printf("CC检查是否有del操作\n")
		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}
		//fmt.Printf("CC检查del操作\n")
		this.Redirect("/category", 301)
		return
	}

	this.Data["IsCategory"] = true
	this.TplName = "category.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	//fmt.Printf("CC show  Data\n")
	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	//topics, err := models.GetAllTopics(this.Input().Get("cate"), true)
	//n := len(topics)
	//fmt.Printf("CC  文章数 len =%d\n", n)
	//this.Data["TopicCount"] = len(topics)
	if err != nil {
		beego.Error(err)
	}
}
