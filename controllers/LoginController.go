package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplNames = "index.html"
}

func (this *LoginController) Post() {
	beego.Info("Login")
    // Get form value.
    uname := this.GetString("uname")
    password := this.GetString("password")

    // Check valid.
    if len(uname) == 0 {
        this.Redirect("/", 302)
        return
    }

    beego.Info(password)

    this.Redirect("/config.html", 302)

    return
}
