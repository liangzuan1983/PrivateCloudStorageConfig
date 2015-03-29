package controllers

import (
	"github.com/astaxie/beego"
)

type MainConfigController struct {
	beego.Controller
}

func (this *MainConfigController) Get() {
	this.TplNames = "config.html"
}