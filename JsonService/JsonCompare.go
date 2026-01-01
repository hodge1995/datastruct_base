package JsonService

import (
	"fmt"
	"reflect"
)

var resBool bool

// CompareJsonDict 对比两个json
func CompareJsonDict(json1, json2 map[string]interface{}) bool {
	resBool = true
	for k1, v1 := range json1 {
		switch v1.(type) {
		case map[string]interface{}:
			if _, ok := json2[k1]; !ok {
				resBool = false
				fmt.Println(v1, json2[k1].(map[string]interface{}))
				return false
			} else {
				CompareJsonDict(json1[k1].(map[string]interface{}), json2[k1].(map[string]interface{}))
			}
		case []interface{}:
			CompareJsonSlice(json1[k1].([]interface{}), json2[k1].([]interface{}))
		default:
			if !reflect.DeepEqual(v1, json2[k1]) {
				fmt.Println(v1, json2[k1])
				resBool = false
				break
			}
		}
	}

	return resBool
}

func CompareJsonSlice(json1, json2 []interface{}) bool {
	for k1, v1 := range json1 {
		switch v1.(type) {
		case map[string]interface{}:
			if _, ok := json2[k1].(map[string]interface{}); !ok {
				fmt.Println(v1, json2[k1].(map[string]interface{}))
				resBool = false
				return false
			} else {
				CompareJsonDict(json1[k1].(map[string]interface{}), json2[k1].(map[string]interface{}))
			}
		case []interface{}:
			CompareJsonSlice(json1[k1].([]interface{}), json2[k1].([]interface{}))
		default:
			if !reflect.DeepEqual(v1, json2[k1]) {
				fmt.Println(v1, json2[k1])
				resBool = false
				return false
			}
		}
	}

	return resBool
}
