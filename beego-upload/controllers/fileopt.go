package controllers

import (

	"github.com/astaxie/beego"
)

type FileOptUploadController struct {
	beego.Controller
}

type UploadController struct {
	beego.Controller
}
 
func (this *UploadController) Get() {
 
	this.TplName="Upload.html"
}
 
func (this *UploadController)Post()  {
	file,head,err:=this.GetFile("file")
	if err!=nil {
		this.Ctx.WriteString("获取文件失败")
		return
	}
	defer file.Close()
 
	filename:=head.Filename
	err =this.SaveToFile("file","static/uploadfile/"+filename)
	if err!=nil {
		this.Ctx.WriteString("上传失败1")
	}else {
		this.Ctx.WriteString("上传成功")
	}
}