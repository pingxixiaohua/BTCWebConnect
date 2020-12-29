package main

import (
	_ "ConnectBC/routers"
	"ConnectBC/rpcrequest"
	"fmt"
)

func main() {

	//1.stop
	//rpcrequest.Stop()
	//2.getblockcount
	rpcrequest.Getblockcount()
	//3.getbestblockhash
	rpcrequest.Getbestblockhash()
	//4.getblock
	err := rpcrequest.Getblock("000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//5.getnewaddress





	//静态资源路径设置
	//beego.SetStaticPath("/js","./static/js")
	//beego.SetStaticPath("/css","./static/css")
	//beego.SetStaticPath("/img","./static/img")


	//beego.Run()
}

