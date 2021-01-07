package routers

import (
	"BitData/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.LoginControllers{})
    beego.Router("/login",&controllers.LoginControllers{})
    //getblockhash功能
    beego.Router("/getblockhash",&controllers.GetblockHashController{})
    //getblock功能
    beego.Router("/getblock",&controllers.GetblockController{})
    //getblockcount功能
    beego.Router("/getblockcount",&controllers.GetblockCountController{})
    //getbestblockhash功能
    beego.Router("/getbestblockhash",&controllers.GetbestHashController{})
    //getblockheader功能
    beego.Router("/getblockheader",&controllers.GetblockHeaderController{})
    //getblockchaininfo功能
    beego.Router("/getblockchaininfo",&controllers.GetblockChainInfoController{})
    //getchaintips功能
    beego.Router("/getchaintips",&controllers.GetChainTipsController{})
    //getdifficulty功能
    beego.Router("/getdifficulty",&controllers.GetDifficultyController{})
    //getmempoolinfo功能
    beego.Router("/getmempoolinfo",&controllers.GetMemPoolInfoController{})
    //gettxoutsetinfo功能
    beego.Router("/gettxoutsetinfo",&controllers.GetTxOutSetInfoController{})
    //preciousblock功能
    beego.Router("/preciousblock",&controllers.GetPreciousController{})
    //getblockfilter功能
    beego.Router("/getblockfilter",&controllers.GetBlockFilterController{})
    //getchaintxstats功能
    beego.Router("/getchaintxstats",&controllers.GetChainTxStatsController{})
    //getwalletinfo功能
    beego.Router("/getwalletinfo",&controllers.GetWalletInfoController{})
    //getunconfirmedbalance功能
    beego.Router("/getunconfirmedbalance",&controllers.GetUnconfirmedBalanceController{})
}
