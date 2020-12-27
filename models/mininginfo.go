package models



type MiningInfo struct {
	Blocks int64 `json:"blocks"`
	Difficulty float64 `json:"difficulty"`
	Networkhashps float32 `json:"networkhashps"`
	Pooledtx int `json:"pooledtx"`
	Chain string `json:"chain"`
	Warnings string `json:"warnings"`
}
type Mining struct {
	Result MiningInfo `json:"result"`
	Error string `json:"error"`
	Id int `json:"id"`
}