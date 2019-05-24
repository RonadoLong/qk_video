package jsonResult

import jsoniter "github.com/json-iterator/go"

// StructsToJsonStr 打印
// query 结构体
func StructsToJsonStr(query interface{}) string {
	bytes, _ := jsoniter.Marshal(query)
	return string(bytes)
}
