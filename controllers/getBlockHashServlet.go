package controllers

import (
	"BitData/RPC"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type GetblockHashController struct {
	beego.Controller
}
func (g *GetblockHashController)Post(){
	g.Ctx.Request.ParseForm()
	heightStr := g.Ctx.Request.FormValue("height")
	heightInt,err:=strconv.Atoi(heightStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	Str := RPC.GetBlockHash(heightInt)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}
