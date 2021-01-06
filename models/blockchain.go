package models

type Chian struct {
	Result BlockChain `json:"result"`
	Error string `json:"error"`
	Id int `json:"id"`
}

type BlockChain struct {
	Chain string `json:"chain"`
	Blocks int `json:"blocks"`
	Headers int `json:"headers"`
	Bestblockhash string `json:"bestblockhash"`
	Difficulty float64 `json:"difficulty"`
	Mediantime int `json:"mediantime"`
	Verificationprogress float64 `json:"verificationprogress"`
	Initialblockdownload bool `json:"initialblockdownload"`
	Chainwork string `json:"chainwork"`
	Size_on_disk int `json:"size_on_disk"`
	Pruned bool `json:"pruned"`
	Pruneheight int `json:"pruneheight"`
	Automatic_pruning bool `json:"automatic_pruning"`
	Prune_target_size int `json:"prune_target_size"`
	softforks Softforks `json:"softforks"`
	warnings string `json:"warnings"`
}
type Softforks struct {
	Bip_34 Bip34 `json:"bip34"`
	Bip_66 Bip66 `json:"bip66"`
	Bip_65 Bip65 `json:"bip65"`
	Csv Csv `json:"csv"`
	Segwit Segwit `json:"segwit"`
}
type Bip34 struct {
	Type string `json:"type"`
	Active bool `json:"active"`
	Height int `json:"height"`
}
type Bip66 struct {
	Type string `json:"type"`
	Active bool `json:"active"`
	Height int `json:"height"`
}
type Bip65 struct {
	Type string `json:"type"`
	Active bool `json:"active"`
	Height int `json:"height"`
}
type Csv struct {
	Type string `json:"type"`
	Active bool `json:"active"`
	Height int `json:"height"`
}
type Segwit struct {
	Type string `json:"type"`
	Active bool `json:"active"`
	Height int `json:"height"`
} 