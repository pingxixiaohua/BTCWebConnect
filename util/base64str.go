package util

import "encoding/base64"

func Base64str(msg string)string  {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}
