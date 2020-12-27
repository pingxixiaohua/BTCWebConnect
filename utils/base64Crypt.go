package utils

import "encoding/base64"

/**
 *	base64编码，将请求的格式转换
 */
func Base64Str(msg string) string {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}
