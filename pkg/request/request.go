// Package request 存放辅助方法
package request

import (
	"net/http"
)


func Query(k string, defaultValue string, r *http.Request)string {
	v := r.URL.Query().Get(k)

	if v != "" {
		return v
	} else if v == "" && defaultValue != "" {
		return defaultValue
	}

	return ""
}
