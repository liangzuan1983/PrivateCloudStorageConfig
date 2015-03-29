package routers

import (
	"github.com/oikomi/PrivateCloudStorageConfig/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    beego.Router("/", &controllers.LoginController{})
    beego.Router("/config", &controllers.MainConfigController{})
}
