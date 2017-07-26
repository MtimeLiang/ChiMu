package routers

import (
	"ChiMu/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/wine/address_add", &controllers.AddressController{})
}
