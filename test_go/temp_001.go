package main

import (
    "flag"
    "github.com/astaxie/beego"
)

func init() {
    const (
        defaultPort = 8085
        usage       = "set port"
    )
    flag.IntVar(&beego.HttpPort, "port", defaultPort, usage)
}

type MainController struct {
    beego.Controller
}

func (this *MainController) Get() {
    this.Ctx.WriteString("hello world")
}

func main() {
    flag.Parse()
    println(beego.HttpPort)
    beego.Router("/", &MainController{})
    beego.Run()
}