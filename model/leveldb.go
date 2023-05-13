package model

import (
	"encoding/json"
)

func GetKey(key string) string {

	LevelDB.Put([]byte(key), []byte("value"), nil)

	var m map[string]interface{}
	m = make(map[string]interface{})
	m["str"] = "string"
	m["int"] = 100
	userJson, _ := json.Marshal(m)
	LevelDB.Put([]byte(key), userJson, nil)
	value, _ := LevelDB.Get([]byte(key), nil)
	json.Unmarshal(value, &m)

	return string(value)

}
