package controllers

import (
	"BitData/RPC"
	"github.com/astaxie/beego"
)

type GetblockChainInfoController struct {
	beego.Controller
}
func (g *GetblockChainInfoController)Post(){
	g.Ctx.Request.ParseForm()
	infoStr := g.Ctx.Request.FormValue("header")
	Str := RPC.GetBlockChainInfo(infoStr)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}
