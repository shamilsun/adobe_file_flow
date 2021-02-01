package helpers

import (
	"encoding/json"
	"log"
)

func ContainsInt(s []uint64, e uint64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ToJSONString(v interface{}) string {
	res2B, _ := json.Marshal(v)
	return string(res2B)
}

func ItemExists(arr []string, v string) bool {
	if len(arr) > 0 {
		for i := 0; i < len(arr); i++ {
			if arr[i] == v {
				//				log.Println("exist")
				return true
			}
		}
	}
	return false
}

func RemoveItem(arr []string, v string) []string {
	var list []string

	if len(arr) > 0 {
		for i := 0; i < len(arr); i++ {
			if arr[i] != v {
				list = append(list, arr[i])
			}
		}
	}
	log.Println("removed")
	log.Println(list)
	return list
}
