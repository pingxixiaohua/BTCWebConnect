package controllers

import (
	"BitData/RPC"
	"github.com/astaxie/beego"
)

type GetblockHeaderController struct {
	beego.Controller
}
func (g *GetblockHeaderController)Post(){
	g.Ctx.Request.ParseForm()
	headerStr := g.Ctx.Request.FormValue("header")
	Str := RPC.GetBlockHeader(headerStr)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}
