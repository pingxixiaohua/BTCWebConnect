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

func GetChainTips()(models.Tips,models.Tips)  {
	rpcrequest:=models.RPCRequest{
		Id:      time.Now().Unix(),
		Method:models.GETCHAINTIPS ,
		Jsonrpc: models.RPCVERSION,
	}
	reqBytes,err:=json.Marshal(&rpcrequest)
	if err!=nil {
		fmt.Println(err.Error())

	}
	fmt.Println("准备好的json数据：",string(reqBytes))

	//发送一个post请求
	reader:=bytes.NewReader(reqBytes)
	client:=&http.Client{}
	request,err:=http.NewRequest("POST",models.RPCURL,reader)
	if err!=nil {
		fmt.Println(err.Error())

	}

	msg:=models.RPCUSER+":"+models.RPCPASSOWRD
	msg=Base64str(msg)
	request.Header.Add("Encoding","UTF-8")
	request.Header.Add("Content-Type","application/json")
	request.Header.Add("Authorization",("Basic "+msg))

	response,err:=client.Do(request)
	if err!=nil {
		fmt.Println(err.Error())

	}
	code:=response.StatusCode
	defer response.Body.Close()
	body,err:=ioutil.ReadAll(response.Body)

	data1,data2:=gettips(body)

	if code==200 {
		fmt.Println("请求成功")
		return data1,data2
	}else {
		fmt.Println("请求失败")
	}

	return data1,data2
}
func gettips(body []byte)(models.Tips,models.Tips) {
	var tips models.Chaintips
	json.Unmarshal([]byte(body),&tips)
	return tips.Result[0],tips.Result[1]
}
