package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginControllers struct {
	beego.Controller
}
type User struct {
	name string `json:"name"`
	password string `json:"password"`
}
func (l *LoginControllers)Get(){
	l.TplName="newenter.html"
}
func (l *LoginControllers)Post(){
	l.Ctx.Request.ParseForm()
	loginUserName := l.Ctx.Request.FormValue("user")
	loginUserPassword := l.Ctx.Request.FormValue("password")
	fmt.Printf("%T %T",loginUserPassword,loginUserName)
	if loginUserName =="123" && loginUserPassword == "123" {
		l.TplName="page.html"
		return
	}
	l.Ctx.WriteString("登陆错误，请重试！")

}