package models

type BLOCKHEADER struct {
	Result BlockHeader `json:"result"`
	Error string `json:"error"`
	Id int `json:"id"`
}
type BlockHeader struct {
	HASH string `json:"hash"`
	CONFIRMATIONS int `json:"confirmations"`
	HEIGHT int `json:"height"`
	VERSION int `json:"version"`
	VERSIONHEX string `json:"versionHex"`
	MERKLEROOT string `json:"merkleroot"`
	TIME int `json:"time"`
	MEDIANTIME int `json:"mediantime"`
	NONCE int `json:"nonce"`
	BITS string `json:"bits"`
	DIFFICULTY int `json:"difficulty"`
	CHAINWORK string `json:"chainwork"`
	NTX int `json:"nTx"`
	PREVIOUSBLOCKHASH string `json:"previousblockhash"`
	NEXTBLOCKHASH string `json:"nextblockhash"`
}
