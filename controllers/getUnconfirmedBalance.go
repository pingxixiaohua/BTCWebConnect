package controllers

import (
	"BitData/RPC"
	"github.com/astaxie/beego"
	"strconv"
)

type GetUnconfirmedBalanceController struct {
	beego.Controller
}

func (g *GetUnconfirmedBalanceController )Post(){
	g.Ctx.Request.ParseForm()
	balanceSTR := g.Ctx.Request.FormValue("balance")
	balanceInt,_ := strconv.Atoi(balanceSTR)
	Str := RPC.GetUnconfirmedBalance(balanceInt)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}
