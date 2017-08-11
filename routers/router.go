package routers

import (
	"ChiMu/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	registerAddressControllers()
	registerBannerControllers()
	registerCartControllers()
	registerCommentControllers()
	registerCouponControllers()
	registerImageControllers()
	registerMyCenterControllers()
	registerMyCouponControllers()
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
	beego.Router("/wine/banner_modify", &controllers.BannerModifyController{})
	beego.Router("/wine/banner_delete", &controllers.BannerDeleteController{})
}

func registerCartControllers() {
	beego.Router("/wine/cart_add", &controllers.CartAddController{})
	beego.Router("/wine/cart", &controllers.CartController{})
	beego.Router("/wine/cart_delete", &controllers.CartDeleteController{})
}

func registerCommentControllers() {
	beego.Router("/wine/comment", &controllers.CommentController{})
	beego.Router("/wine/comment_add", &controllers.CommentAddController{})
}

func registerCouponControllers() {
	beego.Router("/wine/coupon", &controllers.CouponController{})
	beego.Router("/wine/coupon_add", &controllers.CouponAddController{})
	beego.Router("/wine/coupon_modify", &controllers.CouponModifyController{})
	beego.Router("/wine/coupon_delete", &controllers.CouponDeleteController{})
}

func registerImageControllers() {
	beego.Router("/wine/image", &controllers.ImageController{})
}

func registerMyCenterControllers() {
	beego.Router("/wine/my", &controllers.MyUserInfoController{})
}

func registerMyCouponControllers() {
	beego.Router("/wine/my_coupon_add", &controllers.MyCouponAddController{})
	beego.Router("/wine/my_coupon", &controllers.MyCouponController{})
}

func registerPointControllers() {
	beego.Router("/wine/point", &controllers.PointController{})
}
