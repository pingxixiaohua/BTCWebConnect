package models

/**
 * rpc的连接区块链节点的请求格式（rpc通信协议规范）
 *
  {
	"id":,
	"method":
	"jsonrpc":2.0,
	"params":
  }
*/
type RPCRequest struct {
	Id      int64         `json:"id"`      //请求的id
	Method  string        `json:"method"`  //要请求的具体的命令
	Jsonrpc string        `json:"jsonrpc"` //rpc通信协议的版本
	Params  []interface{} `json:"params"`  //请求携带的参数
}

/**
 *	返回响应的格式
 */
type RPCResult struct {
	Id     int         `json:"id"`//对应请求的id
	Error  string      `json:"error"`//是否有错误
	Result interface{} `json:"result"`//返回的额内容
}
