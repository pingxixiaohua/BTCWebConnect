package models

type RPCRequest struct {
	Id   int64  `json:"id"`
	Method string `json:"method"`
	Jsonrpc string `json:"jsonrpc"`
	Params []byte `json:"params"`
}
