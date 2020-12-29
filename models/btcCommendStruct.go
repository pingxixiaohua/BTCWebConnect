package models

/**
 *	getblock命令获取到的内容
 */
type Getblock struct {
	Hash          string   `json:"hash"`
	Confirmations int      `json:"confirmations"`
	Strippedsize  int      `json:"strippedsize"`
	Size          int      `json:"size"`
	Weight        int      `json:"weight"`
	Height        int      `json:"height"`
	Version       int      `json:"version"`
	VersionHex    string   `json:"versionHex"`
	Merkleroot    string   `json:"merkleroot"`
	Tx            []string `json:"tx"`
	Time          float64  `json:"time"`
	Mediantime    int      `json:"mediantime"`
	Nonce         int      `json:"nonce"`
	Bits          string   `json:"bits"`
	Difficulty    int      `json:"difficulty"`
	Chainwork     string   `json:"chainwork"`
	Ntx           int      `json:"nTx"`
}
