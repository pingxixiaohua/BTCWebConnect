package models

type Mempollinfo struct {
	Result Mem `json:"result"`
	Error  string  `json:"error"`
	Id     int     `json:"id"`
}

type Mem struct {
	Loaded bool `json:"loaded"`
	Size int `json:"size"`
	Bytes int `json:"bytes"`
	usage int `json:"usage"`
	Maxmempool int `json:"maxmempool"`
	Mempoolminfee float64 `json:"mempoolminfee"`
	Minrelaytxfee float64 `json:"minrelaytxfee"`
}
