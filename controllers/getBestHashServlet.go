package controllers

import (
	"BitData/RPC"
	"github.com/astaxie/beego"

)

type GetbestHashController struct {
	beego.Controller
}
func (g *GetbestHashController)Post(){
	g.Ctx.Request.ParseForm()
	bestStr := g.Ctx.Request.FormValue("best")
	Str := RPC.GetBestBlockHash(bestStr)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}
