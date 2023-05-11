package model

import (
	"encoding/json"
)

func GetKey(key string) string {

	LevelDB.Put([]byte(key), []byte("value"), nil)

	var m map[string]interface{}
	m = make(map[string]interface{})
	m["name"] = "韩信"
	m["age"] = 23
	m["address"] = "野区"
	userJson, _ := json.Marshal(m)
	LevelDB.Put([]byte(key), userJson, nil)
	value, _ := LevelDB.Get([]byte(key), nil)
	json.Unmarshal(value, &m)

	return string(value)

}
