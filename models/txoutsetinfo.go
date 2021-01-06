package models

type TXOUTSET struct {
	Result Txoutsetinfo `json:"result"`
	Error string `json:"error"`
	Id int `json:"id"`
} 
type Txoutsetinfo struct {
	HEIGHT int `json:"height"`
	BESTBLOCK string `json:"bestblock"`
	TRANSACTONS int `json:"transactons"`
	TXOUTS int `json:"txouts"`
	BOGOSIZE int `json:"bogosize"`
	HASH_SERIALIZED_2 string `json:"hash_serialized_2"`
	DISK_SIZE int `json:"disk_size"`
	TOTAL_AMOUNT float64 `json:"total_amount"`
} 
