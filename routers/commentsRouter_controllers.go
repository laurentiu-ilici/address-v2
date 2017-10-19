package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["address/controllers:AddressController"] = append(beego.GlobalControllerRouter["address/controllers:AddressController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["address/controllers:AddressController"] = append(beego.GlobalControllerRouter["address/controllers:AddressController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:postcode`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
