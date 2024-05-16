package utils

import "encoding/json"

func ByteToMap(payload []byte)(result map[string]interface{},err error){
	if err = json.Unmarshal(payload, &result); err != nil {
		return result, err
	}
	return result, err
}
