package controllers

import (
	"BitData/RPC"
	"github.com/astaxie/beego"
)

type GetMemPoolInfoController struct {
	beego.Controller
}

func (g *GetMemPoolInfoController )Post(){
	g.Ctx.Request.ParseForm()
	poolinfoSTR := g.Ctx.Request.FormValue("poolinfo")
	Str := RPC.GetMemPoolInfo(poolinfoSTR)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}
