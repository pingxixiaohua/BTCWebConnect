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

func GetNetWorkHashPs() float64 {
	rpcrequest:=models.RPCRequest{
		Id:      time.Now().Unix(),
		Method:models.GETNETWORKHASHPS ,
		Jsonrpc: models.RPCVERSION,
	}
	reqBytes,err:=json.Marshal(&rpcrequest)
	if err!=nil {
		fmt.Println(err.Error())
		return 0
	}
	fmt.Println("准备好的json数据：",string(reqBytes))

	//发送一个post请求
	reader:=bytes.NewReader(reqBytes)
	client:=&http.Client{}
	request,err:=http.NewRequest("POST",models.RPCURL,reader)
	if err!=nil {
		fmt.Println(err.Error())
		return 0
	}

	msg:=models.RPCUSER+":"+models.RPCPASSOWRD
	msg=Base64str(msg)
	request.Header.Add("Encoding","UTF-8")
	request.Header.Add("Content-Type","application/json")
	request.Header.Add("Authorization",("Basic "+msg))

	response,err:=client.Do(request)
	if err!=nil {
		fmt.Println(err.Error())
		return 0
	}
	code:=response.StatusCode
	defer response.Body.Close()
	body,err:=ioutil.ReadAll(response.Body)


	data:=getnetworkhashps(body)



	if code==200 {
		fmt.Println("请求成功")
		return data
	}else {
		fmt.Println("请求失败")
	}


	return 0

}
func getnetworkhashps(body []byte)float64  {
	var net models.NetWorkHashPs
	json.Unmarshal([]byte(body),&net)
	return net.Result
}

