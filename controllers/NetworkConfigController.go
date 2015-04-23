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

func (this *NetworkConfigController) Post() {
    this.TplNames = "config/network.html"
	beego.Info("NetworkConfigController Post")
	ip := this.GetString("ip")

	beego.Info(ip)
    //this.Data["json"] = "{\"ObjectId\":\"" + "9" + "\"}"
    //this.ServeJson()
    //this.Ctx.WriteString("{\"ObjectId\":\"" + "9" + "\"}")
    //return
    //this.Ctx.Output.SetStatus(200)
}

