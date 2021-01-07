package controllers

import (
	"BitData/RPC"
	"github.com/astaxie/beego"
	"strconv"
)

type GetDifficultyController struct {
	beego.Controller
}

func (g *GetDifficultyController )Post(){
	g.Ctx.Request.ParseForm()
	difficultySTR := g.Ctx.Request.FormValue("difficulty")
	difficultyInt,_:=strconv.Atoi(difficultySTR)
	Str := RPC.GetDifficulty(difficultyInt)
	g.Data["rpcdate"] = Str
	g.TplName = "page.html"
}
