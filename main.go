package main

import (
	_ "ConnectBC/routers"
	"ConnectBC/rpcrequest"
)

func main() {

	rpcrequest.Getblock("000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f")
	rpcrequest.Getbestblockhash()
	rpcrequest.Getblockcount()
	rpcrequest.Getblockchaininfo()
	//rpcrequest.Stop()

	//静态资源路径设置
	//beego.SetStaticPath("/js","./static/js")
	//beego.SetStaticPath("/css","./static/css")
	//beego.SetStaticPath("/img","./static/img")
	//
	//
	//beego.Run()
}

