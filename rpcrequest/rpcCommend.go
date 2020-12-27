package rpcrequest

import (
	"ConnectBC/constant"
	"ConnectBC/controllers"
	"errors"
	"fmt"
)

/**
 * 	stop：将钱包关闭
 */
func Stop() {
	rpcResult, err := controllers.RpcConnect(constant.STOP)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("已关闭比特币节点",rpcResult)
}
/**
 * getblock：根据区块hash获取具体的区块内容
 */
func Getblock(params interface{}) error {
	//判断传入的参数params是否是int类型
	switch params.(type) {
	case string:
		//调用RpcConnect，发起连接
		rpcResult, err := controllers.RpcConnect(constant.GETBLOCK,params)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println("getblock返回的内容：",rpcResult)
		//用相对应的类型接受返回内容中的响应结果rpcResult.Result
		resultMap := rpcResult.Result.(map[string]interface{})
		//查看某一个特定的参数
		fmt.Println(resultMap["hash"])
	default:
		return errors.New("传入的类型不符合参数的类型，请输入string类型")
	}
	return nil
}

/**
 * getbestblockhash 获取最新区块的hash值
 */
func Getbestblockhash() {
	rpcResult, err := controllers.RpcConnect(constant.GETBESTBLOCKHASH)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("getbestblockhash返回的内容：",rpcResult)
	resultStr := rpcResult.Result
	fmt.Println(resultStr)
}

