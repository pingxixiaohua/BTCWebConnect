package controllers

import (
	"BitData/RPC"
	"github.com/astaxie/beego"
)

type GetBlockFilterController struct {
	beego.Controller
}

func (g *GetBlockFilterController )Post(){
	g.Ctx.Request.ParseForm()
	filterSTR := g.Ctx.Request.FormValue("filter")
	Str := RPC.GetBlockFilter(filterSTR)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}
