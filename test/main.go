package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	_ "test/routers"
)

func main() {
	i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	i18n.SetMessage("en-US", "conf/locale_zh-CN.ini")
	beego.Run()
}
