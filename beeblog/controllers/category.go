package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	fmt.Printf("CCCCCCCC\n")
	this.TplName = "category.html"
}

func (this *CategoryController) Post() {
	
}