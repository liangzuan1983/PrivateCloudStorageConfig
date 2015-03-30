package routers

import (
	"github.com/oikomi/PrivateCloudStorageConfig/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.LoginController{})
    beego.Router("/data", &controllers.DataController{})
    beego.Router("/config", &controllers.MainConfigController{})
    beego.Router("/config/network", &controllers.NetworkConfigController{})
}
