package serial

import (
	"encoding/json"
)

type DataMap map[string]interface{}

func Json2Map(str string) *DataMap {
	dataMap := DataMap{}
	_ = json.Unmarshal([]byte(str), &dataMap)
	return &dataMap
}

func Object2Json(obj interface{}) string {
	res, _ := json.Marshal(obj)
	return string(res)
}

// Json2Instant make json to a type.
// 	str := "{\"redis\":{\"host\":\"46.198.212.111\",\"port\":\"6379\"}}"
//	r := yaml.R{}
//	_ = serial.Json2Instant(str, &r)
//	fmt.Println(r.Redis.Host)
// json deserialized
func Json2Instant(str string, param interface{}) *interface{} {
	_ = json.Unmarshal([]byte(str), &param)
	return &param
}
