package util

import (
	"BitData/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetBestBlockHash() string {
	rpcrequest:=models.RPCRequest{
		Id:      time.Now().Unix(),
		Method:models.GETBESTBLOCKHASH ,
		Jsonrpc: models.RPCVERSION,
	}
	reqBytes,err:=json.Marshal(&rpcrequest)
	if err!=nil {
		fmt.Println(err.Error())
		return "错误"
	}
	fmt.Println("准备好的json数据：",string(reqBytes))

	//发送一个post请求
	reader:=bytes.NewReader(reqBytes)
	client:=&http.Client{}
	request,err:=http.NewRequest("POST",models.RPCURL,reader)
	if err!=nil {
		fmt.Println(err.Error())
		return "错误"
	}

	msg:=models.RPCUSER+":"+models.RPCPASSOWRD
	msg=Base64str(msg)
	request.Header.Add("Encoding","UTF-8")
	request.Header.Add("Content-Type","application/json")
	request.Header.Add("Authorization",("Basic "+msg))

	response,err:=client.Do(request)
	if err!=nil {
		fmt.Println(err.Error())
		return "错误"
	}
	code:=response.StatusCode
	defer response.Body.Close()
	body,err:=ioutil.ReadAll(response.Body)


	data:=getdata(body)



	if code==200 {
		fmt.Println("请求成功")
		return data
	}else {
		fmt.Println("请求失败")
	}


	return "错误"

}
func getdata(body []byte)string  {
	var besthash models.BestHash
	json.Unmarshal([]byte(body),&besthash)
	return besthash.Result
}
