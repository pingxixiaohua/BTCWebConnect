package controllers

import (
	"BitData/RPC"
	"github.com/astaxie/beego"
)

type GetChainTipsController struct {
	beego.Controller
}

func (g *GetChainTipsController )Post(){
	g.Ctx.Request.ParseForm()
	tipsSTR := g.Ctx.Request.FormValue("tips")
	Str := RPC.GetChainTips(tipsSTR)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}

