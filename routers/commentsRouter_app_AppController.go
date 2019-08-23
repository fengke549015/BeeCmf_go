package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/BeeCmf/app/AppController:IndexController"] = append(beego.GlobalControllerRouter["github.com/BeeCmf/app/AppController:IndexController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/BeeCmf/app/AppController:IndexController"] = append(beego.GlobalControllerRouter["github.com/BeeCmf/app/AppController:IndexController"],
		beego.ControllerComments{
			Method:           "Main",
			Router:           `/main`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/BeeCmf/app/AppController:LoginController"] = append(beego.GlobalControllerRouter["github.com/BeeCmf/app/AppController:LoginController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/BeeCmf/app/AppController:LoginController"] = append(beego.GlobalControllerRouter["github.com/BeeCmf/app/AppController:LoginController"],
		beego.ControllerComments{
			Method:           "LoginPost",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/BeeCmf/app/AppController:MenuController"] = append(beego.GlobalControllerRouter["github.com/BeeCmf/app/AppController:MenuController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/menu`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/BeeCmf/app/AppController:UserController"] = append(beego.GlobalControllerRouter["github.com/BeeCmf/app/AppController:UserController"],
		beego.ControllerComments{
			Method:           "User",
			Router:           `/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
