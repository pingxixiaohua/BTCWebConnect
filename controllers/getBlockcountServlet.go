package controllers

import (
	"BitData/RPC"
	"github.com/astaxie/beego"
	"strconv"
)

type GetblockCountController struct {
	beego.Controller
}
func (g *GetblockCountController)Post(){
	g.Ctx.Request.ParseForm()
	CountStr := g.Ctx.Request.FormValue("Count")
	CountInt,_:=strconv.Atoi(CountStr)
	Str := RPC.GetBlockCount(CountInt)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}

