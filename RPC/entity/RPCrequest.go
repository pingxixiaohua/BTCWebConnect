package entity

type RPCrequest struct {
	ID int64 `json:"id"`
	Method string `json:"method"`
	JsonRpc string `json:"json_rpc"`
	Params []interface{} `json:"params"`
}
