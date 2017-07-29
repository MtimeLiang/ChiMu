package routers

import (
	"ChiMu/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	registerAddressControllers()
	registerBannerControllers()
}

func registerAddressControllers() {
	beego.Router("/wine/address_add", &controllers.AddressAddController{})
	beego.Router("/wine/address_modify", &controllers.AddressModifyController{})
	beego.Router("/wine/address", &controllers.AddressController{})
	beego.Router("/wine/address_selected", &controllers.AddressSelectedController{})
	beego.Router("/wine/address_delete", &controllers.AddressDeleteController{})
}

func registerBannerControllers() {
	beego.Router("/wine/banner", &controllers.BannerListController{})
	beego.Router("/wine/banner_detail", &controllers.BannerDetailController{})
	beego.Router("/wine/banner_add", &controllers.BannerAddController{})
}
