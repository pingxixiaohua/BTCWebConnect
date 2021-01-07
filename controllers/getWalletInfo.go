package controllers

import (
	"BitData/RPC"
	"github.com/astaxie/beego"
)

type GetWalletInfoController struct {
	beego.Controller
}

func (g *GetWalletInfoController )Post(){
	g.Ctx.Request.ParseForm()
	walletSTR := g.Ctx.Request.FormValue("wallet")
	Str := RPC.GetWalletInfo(walletSTR)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}
