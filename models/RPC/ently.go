package RPC

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const RPCUSER = "user"
const RPCPASSWORD = "pwd"
const RPCURL = "http://127.0.0.1:8332"


type  Rpc_Requst struct {
	Requst *http.Request
	Ahead Rpc_Ahead
}

type Rpc_Ahead struct {
	Id int          `json:"id"`
	Jsonrpc string  `json:"jsonrpc"`
	Method string   `json:"method"`
	Params []interface{}  `json:"params"`
}

//创造rpc请求并初始化

func NewRpcRequst()(Rpc_Requst,error){
	R := Rpc_Requst{
		Requst: nil,
		Ahead: Rpc_Ahead{
			Id:    time.Now().Nanosecond(),
			Jsonrpc: "2.0",
			Params:nil,
		},
	}
	var err error
	R.Requst,err = http.NewRequest("POST",RPCURL,nil)
	if err!=nil{
		log.Fatal(err)
		return R,err
	}
	R.Requst.Header.Add("Authorization","Basic "+Base64str(RPCUSER+":"+RPCPASSWORD))
	R.Requst.Header.Add("Encoding","UTF-8")
	R.Requst.Header.Add("Content-Type","application/json")

	return R,nil
}
/*
   发起rpc post请求 并返回结果
*/
func (R Rpc_Requst)Rpc_DoPost(method string,params []interface{})(interface{},error){
	client :=	http.Client{}
	R.Ahead.Method = method
	if(len(params)!=0) {
		R.Ahead.Params = params
	}else{
		R.Ahead.Params = make([]interface{},0)
	}
	fmt.Println("params:",R.Ahead.Params)
	/*
	   序列化请求体
	*/
	jsonbyte,err:=json.Marshal(R.Ahead)
	fmt.Println("send:"+string(jsonbyte))
	if err!=nil{
		fmt.Println(err)
		return "",err
	}
	/*
	   设置请求体
	*/
	read := bytes.NewBuffer(jsonbyte)
	R.Requst.Body = ioutil.NopCloser(read)
	/*
	  客户端发起网络请求
	*/
	response,err :=  client.Do(R.Requst)
	if err!=nil{
		fmt.Println("网络请求失败:",err)
		return "",err
	}
	/*
	   读取响应体
	*/
	resultbyte,err:= ioutil.ReadAll(response.Body)
	if err!=nil{
		fmt.Println("ReadAll Error::",err)
		return "",err
	}

	/*
	   解析请求到的数据保存到map里并返回结果
	*/
	result := make(map[string]interface{})
	json.Unmarshal(resultbyte,&result)
	fmt.Println(result)
	/*
	   检测网络请求状态
	*/
	statuscode  :=	response.StatusCode
	if(statuscode!=200){
		fmt.Println("Post Error! Code:"+string(resultbyte))
		return "",errors.New("Rpc请求失败 error:"+string(resultbyte))
	}

	return result["result"],nil
}

func Base64str(str string)string{
	return base64.StdEncoding.EncodeToString([]byte(str))
}
