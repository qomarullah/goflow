// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"goflow/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		/*beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		*/
		beego.NSNamespace("/query",
			beego.NSInclude(
				&controllers.QueryController{},
			),
		),
		beego.NSNamespace("/goflow",
			beego.NSInclude(
				&controllers.ServicesController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
