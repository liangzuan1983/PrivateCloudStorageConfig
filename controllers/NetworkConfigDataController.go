package controllers

import (
	"github.com/astaxie/beego"
)

type NetworkConfigDataController struct {
	beego.Controller
}

func (this *NetworkConfigDataController) Post() {
    beego.Info("NetworkConfigDataController Post")

    return
}


