package util

import "encoding/json"

func I2Json(i interface{}) string {
	data, _ := json.Marshal(i)
	return string(data)
}
