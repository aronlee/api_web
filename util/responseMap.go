package util

import (
	"reflect"
)

const (
	errorCode         = -1   // 处理失败
	successCode       = 1    // 处理成功
	noPermissionsCode = 1001 //没有权限
	dbErrorCode       = 1002 //数据库异常
	errorMsg          = "处理失败"
	successMsg        = "处理成功"
	noPermissionsMsg  = "没有权限"
	dbErrorMsg        = "数据库异常"
)

// Struct2Map struct to map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// ErrorMap return json data
func ErrorMap() map[string]interface{} {
	result := map[string]interface{}{
		"code": errorCode,
		"msg":  errorMsg,
		"data": "",
	}
	return result
}

// ErrorResMsgMap return json data
func ErrorResMsgMap(msg string) map[string]interface{} {
	result := map[string]interface{}{
		"code": errorCode,
		"msg":  msg,
		"data": "",
	}
	return result
}

// SuccessDataMap return json data
func SuccessDataMap(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code": successCode,
		"msg":  successMsg,
		"data": data,
	}
	return result
}

// SuccessMsgDataMap return json data
func SuccessMsgDataMap(msg string, data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code": successCode,
		"msg":  msg,
		"data": data,
	}
	return result
}
