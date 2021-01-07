package main

import (
	_"BitData/routers"
	_ "BitData/util"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {



	fmt.Println("^-^")

	//besthash:=util.GetBestBlockHash()
	//block:=util.GetBlock(besthash)
	//fmt.Println(block)
	//
	//tips1,tips2:=util.GetChainTips()
	//fmt.Println(tips1,tips2)

	//pool:=util.GetRawMemPool()
	//fmt.Println(pool)
	//

	//net:=util.GetNetWorkHashPs()
	//fmt.Println("\n",net)

	//header:=util.GetBlockHeader(besthash)
	//fmt.Println(header)
	//util.VerIfyChain(4,10000)

	//stop:=util.Stop()
	//fmt.Println(stop)

	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()
}

