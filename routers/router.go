package routers

import (
	"golangERP/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 首页,返回的为html，其他页面的请求返回的都为json
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login/?:id", &controllers.LoginContriller{})
	beego.Router("/menu", &controllers.MenuController{})
	beego.Router("/address/country/?:id", &controllers.AddressCountryContriller{})
	beego.Router("/address/province/?:id", &controllers.AddressProvinceContriller{})
	beego.Router("/address/city/?:id", &controllers.AddressCityContriller{})
	beego.Router("/address/district/?:id", &controllers.AddressDistrictContriller{})
}
