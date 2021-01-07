package controllers

import (
	"BitData/RPC"
	"fmt"
	"github.com/astaxie/beego"
)

type GetblockController struct {
	beego.Controller
}

func (g *GetblockController)Post(){
	g.Ctx.Request.ParseForm()
	hash := g.Ctx.Request.FormValue("hash")
	Str := RPC.GetBlock(hash)
	fmt.Println(Str)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}
