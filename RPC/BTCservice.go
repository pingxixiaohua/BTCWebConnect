package RPC

import (
	"bytes"
	"fmt"
)

//GetBlockCount
func GetBlockCount(int) interface{} {
	RPCjson := PrepareJSON("getblockcount")
	data, err := DoPost(RPCURL, bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return data.Result
}

//GetBlockHash
func GetBlockHash(height int) interface{} {
	RPCjson :=PrepareJSON("getblockhash",height)
	date,err:=DoPost(RPCURL,bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return date.Result
}

//GetBlock
func GetBlock(hash string) interface{}{
	RPCjson := PrepareJSON("getblock",hash)
	fmt.Println(string(RPCjson))

	data,err:=DoPost(RPCURL,bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return data.Result
}

//GetBestBlockHash
func GetBestBlockHash(string) interface{} {
	RPCjson := PrepareJSON("getbestblockhash")
	data, err := DoPost(RPCURL, bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return data.Result
}

//GetBlockHeader
func GetBlockHeader(hash string) interface{} {
	RPCjson := PrepareJSON("getblockheader",hash)
	data,err:=DoPost(RPCURL,bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return data.Result
}

//GetBlockChainInfo
func GetBlockChainInfo(string) interface{} {
	RPCjson := PrepareJSON("getblockchaininfo")
	data,err:=DoPost(RPCURL,bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return data.Result
}

//GetChainTips
func GetChainTips(string) interface{} {
	RPCjson := PrepareJSON("getchaintips")
	data,err:=DoPost(RPCURL,bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return data.Result
}

//GetDifficulty
func GetDifficulty(int) interface{} {
	RPCjson := PrepareJSON("getdifficulty")
	data,err:=DoPost(RPCURL,bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return data.Result
}

//GetMemPoolInfo
func GetMemPoolInfo(string) interface{} {
	RPCjson := PrepareJSON("getmempoolinfo")
	data,err:=DoPost(RPCURL,bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return data.Result
}

//GetTxOutSetInfo
func GetTxOutSetInfo(string) interface{} {
	RPCjson := PrepareJSON("gettxoutsetinfo")
	data,err:=DoPost(RPCURL,bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return data.Result
}

//GetPreciousBlock
func GetPreciousBlock(hash string) interface{} {
	RPCjson := PrepareJSON("preciousblock",hash)
	data,err:=DoPost(RPCURL,bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return data.Result
}

func GetBlockFilter(hash string) interface{} {
	RPCjson := PrepareJSON("getblockfilter",hash)
	data,err:=DoPost(RPCURL,bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return data.Result
}
func GetChainTxStats(string) interface{} {
	RPCjson := PrepareJSON("getchaintxstats")
	data,err:=DoPost(RPCURL,bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return data.Result
}
func GetWalletInfo(string) interface{} {
	RPCjson := PrepareJSON("getwalletinfo")
	data,err:=DoPost(RPCURL,bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return data.Result
}
func GetUnconfirmedBalance(int) interface{} {
	RPCjson := PrepareJSON("getunconfirmedbalance")
	data,err:=DoPost(RPCURL,bytes.NewBuffer(RPCjson))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return data.Result
}