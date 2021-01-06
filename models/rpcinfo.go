package models

type Rpcinfo struct {
	Result RPC `json:"result"`
	Error string `json:"error"`
	Id int `json:"id"`

}
type RPC struct {
	Active_commands [1]Activecommandse `json:"active_commands"`
	Logpath string `json:"logpath"`
}
type Activecommandse struct {
	Method string `json:"method"`
	Duration int `json:"duration"`
}
