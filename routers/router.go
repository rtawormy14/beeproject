package routers

import (
	"github.com/rtawormy14/beeproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
