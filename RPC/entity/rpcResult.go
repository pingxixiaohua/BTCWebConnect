package entity

type Data struct {
	Id     int         `json:"id"`
	Error  error       `json:"error"`
	Result interface{} `json:"result"`
}

type RPCresult struct {
	code int
	msg  string
	data Data
}
