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
	str:= l.Ctx.Request.Form
	fmt.Println(str)
	var u User
    fmt.Println(u)
	l.TplName="page.html"
}