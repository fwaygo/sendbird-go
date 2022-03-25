package sendbird

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/fwaygo/sendbird-go/api"
)

type ChannelListParams string

func formatArrayToParams(key string, values []interface{}) string {
	query := ""
	for i, val := range values {
		query += val.(string)
		if i != len(values)-1 {
			query += ","
		}
	}
	return query
}

func structToMap(data interface{}) map[string]interface{} {
	toRet := make(map[string]interface{})
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &toRet)
	return toRet
}

func EncodeParameters(getRequest *api.ChannelListRequest) string {
	paramString := "?"
	requestMap := structToMap(getRequest)

	for key, value := range requestMap {
		valType := reflect.TypeOf(value).String()
		paramString += key + "="

		switch valType {
		case "bool":
			paramString += strconv.FormatBool(value.(bool))
		case "string":
			paramString += value.(string)
		case "[]interface {}":
			paramString += formatArrayToParams(key, value.([]interface{}))
		}
		paramString += "&"
	}
	return paramString
}
