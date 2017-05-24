package routers

import (
	"PagingBybeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	// beego.Router("/all", &controllers.PersonController{})
	beego.Router("/user/all", &controllers.PersonController{},"GET:AllPeople")
}
