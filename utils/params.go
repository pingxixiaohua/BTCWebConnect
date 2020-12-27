package utils

/**
 *	定义一个接口类型的数组，用于接收响应的内容
 *	参数可能是多个未知的类型，所以用结构体切片表示
 */
func Params(para interface{}) []interface{} {
	return []interface{}{para}
}
