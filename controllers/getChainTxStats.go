package controllers

import (
	"BitData/RPC"
	"github.com/astaxie/beego"
)

type GetChainTxStatsController struct {
	beego.Controller
}

func (g *GetChainTxStatsController )Post(){
	g.Ctx.Request.ParseForm()
	txstatsSTR := g.Ctx.Request.FormValue("txstats")
	Str := RPC.GetChainTxStats(txstatsSTR)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}

