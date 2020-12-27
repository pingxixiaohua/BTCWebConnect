package util

import (
	"BitData/models"
	"encoding/json"
)


func GetData(body []byte)int  {
	var count models.Count
	json.Unmarshal([]byte(body),&count)
	return count.Result


}