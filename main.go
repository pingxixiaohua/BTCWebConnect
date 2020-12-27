package main

import (
	_ "BitData/routers"
	"BitData/util"
	_ "BitData/util"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {



	fmt.Println("^-^")


	tips:=util.GetChainTips
	fmt.Println(tips)
	
	beego.Run()
}

