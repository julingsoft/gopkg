package xsign

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetSign(jsonStr, secret string, queryData ...map[string]interface{}) (string, error) {
	queryStr := MustGetQueryStr(queryData...)
	signStr := gmd5.MustEncryptString(secret + queryStr + jsonStr + secret)
	return gstr.ToUpper(signStr), nil
}

func MustGetQueryStr(queryData ...map[string]interface{}) (queryStr string) {
	var dataMap = make(map[string]json.RawMessage)
	if len(queryData) > 0 {
		for _, qd := range queryData {
			for k, v := range qd {
				dataMap[k] = json.RawMessage(gconv.String(v))
			}
		}
	}

	if len(dataMap) > 0 {
		var keys = make([]string, len(dataMap))
		for key := range dataMap {
			if key != "sign" {
				keys = append(keys, key)
			}
		}
		sort.Strings(keys)

		for _, key := range keys {
			queryStr = queryStr + fmt.Sprintf("%v%v", key, string(dataMap[key]))
		}
	}

	return queryStr
}
