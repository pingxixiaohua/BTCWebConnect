package models

type Chaintips struct {
	Result [2]Tips `json:"result"`
	Error string `json:"error"`
	Id int `json:"id"`
}
//type RESult struct {
//	res [2]Tips
//}
type Tips struct {
	Height int `json:"height"`
	Hash string `json:"hash"`
	Branchlen int `json:"branchlen"`
	Status string `json:"status"`
} 
