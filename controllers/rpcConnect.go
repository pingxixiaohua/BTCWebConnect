package controllers

import (
	"ConnectBC/models"
	"ConnectBC/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func RpcConnect(method string, params ...interface{}) (models.RPCResult, error) {
		//0、定义一个空RPCResult 用来传
		rpcReqNil := models.RPCResult{}
	//1、准备一个json格式的数据（rpc通信协议规范）
	rpcReq := models.RPCRequest{
		Id:      time.Now().Unix(),
		Method:  method,
		Jsonrpc: beego.AppConfig.String("rpc_version"),
	}
	//判断是否有参数
	if params != nil {
		rpcReq.Params = params
	}
	//转换为json格式
	reqBytes, err := json.Marshal(&rpcReq)
	if err != nil {
		fmt.Println(err.Error())
		return rpcReqNil,err
	}
	//fmt.Println("json格式数据请求：",string(reqBytes))

	//2、发送post请求
	client := &http.Client{}
	request, err := http.NewRequest("POST", beego.AppConfig.String("rpc_url"), strings.NewReader(string(reqBytes)))
	if err != nil {
		fmt.Println(err.Error())
		return rpcReqNil,err
	}
	//设置请求头
	request.Header.Add("Encoding", "UTF-8")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Basic "+utils.Base64Str(beego.AppConfig.String("rpc_user") + ":" + beego.AppConfig.String("rpc_password")))

	//接受响应
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return rpcReqNil,err
	}

	code := response.StatusCode//请求状态码
	if code == 200 {
		fmt.Println(method,"请求成功")
		//接收请求体
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
			return rpcReqNil,err
		}
		defer response.Body.Close()

		//将响应反序列化后用RPCResult结构体接收
		RpcResult := models.RPCResult{}
		err = json.Unmarshal(body, &RpcResult)
		if err != nil {
			log.Fatal(err.Error())
			return  rpcReqNil,err
		}
		//fmt.Println("响应的内容：",RpcResult)

		return RpcResult, err
	}else {//请求失败原因，状态码
		fmt.Println("状态码：",code)
		fmt.Println("请求失败")
	}
	return rpcReqNil,errors.New("错误")

}
