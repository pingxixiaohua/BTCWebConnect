package controllers

import (
	"BitData/RPC"
	"github.com/astaxie/beego"
)

type GetPreciousController struct {
	beego.Controller
}

func (g *GetPreciousController )Post(){
	g.Ctx.Request.ParseForm()
	preciousSTR := g.Ctx.Request.FormValue("precious")
	Str := RPC.GetPreciousBlock(preciousSTR)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}

