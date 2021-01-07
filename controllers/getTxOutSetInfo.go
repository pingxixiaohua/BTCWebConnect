package controllers

import (
	"BitData/RPC"
	"github.com/astaxie/beego"
)

type GetTxOutSetInfoController struct {
	beego.Controller
}

func (g *GetTxOutSetInfoController )Post(){
	g.Ctx.Request.ParseForm()
	txoutsetinfoSTR := g.Ctx.Request.FormValue("txoutsetinfo")
	Str := RPC.GetTxOutSetInfo(txoutsetinfoSTR)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}
