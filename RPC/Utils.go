package RPC

import (
	"BitData/RPC/entity"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//准备rpc连接的json数据
func PrepareJSON(method string, params ...interface{}) []byte {
	request := entity.RPCrequest{
		ID:      time.Now().Unix(),
		Method:  method,
		JsonRpc: "2.0",
		Params:  params,
	}
	ReqBytes, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return ReqBytes
}

//发起rpc请求
func DoPost(url string, body io.Reader) (entity.Data, error) {
	var data entity.Data
	client := &http.Client{}
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println(err.Error())
		return data, err
	}
	request.Header.Add("Encoding", "UTF-8")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Basic "+Base64Encode(RPCUSER+":"+RPCPASSWORD))

	respon, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return data, err
	}
	result, err := ioutil.ReadAll(respon.Body)
	if err != nil {
		fmt.Println(err.Error())
		return data, err
	}
	//responResult:=make(map[string]string)
	fmt.Println(string(result))

	code := respon.StatusCode
	if code == 200 {
		err = json.Unmarshal(result, &data)
		if err != nil {
			log.Fatal(err.Error())
		}
		return data, nil
	} else {
		fmt.Println(code)
	}
	return data, nil
}

//Base64编码工具
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
func Base64Decode(str string) string {
	result, _ := base64.StdEncoding.DecodeString(str)
	return string(result)
}
