package com

import (
	"encoding/json"
)

//Json must convert value to json string
func Json(v interface{}) string {
	bys, _ := json.Marshal(v)
	return string(bys)
}
