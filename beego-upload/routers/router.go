package routers

import (
	"beego-upload/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.UploadController{})
	beego.Router("/upload", &controllers.UploadController{})
}
