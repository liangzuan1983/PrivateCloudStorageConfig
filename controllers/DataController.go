package controllers

import (
	"github.com/astaxie/beego"
)

type DataController struct {
	beego.Controller
}



func (this *DataController) Get() {
    beego.Info("DataController Get")
    mystruct := "{'a':'b'}"
    this.Data["json"] = &mystruct
    this.ServeJson()
}