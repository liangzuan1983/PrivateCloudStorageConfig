package controllers

import (
	"github.com/astaxie/beego"
)

type NetworkConfigController struct {
	beego.Controller
}

func (this *NetworkConfigController) Get() {
	this.TplNames = "config/network.html"
}