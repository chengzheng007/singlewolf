package singlewolf

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

func getRequestParams(r *http.Request) paramsData {
	params := make(paramsData)

	if r == nil {
		return params
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logf("ioutil.ReadAll(r.Body) error(%v)", err)
		return params
	}
	defer r.Body.Close()

	// 没传参数，无需解析
	if len(buf) == 0 {
		return params
	}

	if err := json.Unmarshal(buf, &params); err != nil {
		logf("json.Unmarshal(%s, &params.data) error(%v)", buf, err)
		return params
	}

	return params
}

// paramsData storage paramsters data client sent, and it ha been Json Unmarshaled to map
type paramsData map[string]interface{}

func (p paramsData) GetString(key string) string {
	if p == nil {
		return ""
	}
	if v, ok := p[key].(string); ok {
		return v
	}
	return ""
}

func (p paramsData) GetInt64(key string) int64 {
	if p == nil {
		return 0
	}
	// map interface{} 默认将数字类型设置为float64
	if v, ok := p[key].(float64); ok {
		return reflect.ValueOf(v).Int()
	}
	return 0
}

func (p paramsData) GetBytes(key string) []byte {
	if p == nil {
		return []byte("")
	}

	if v, ok := p[key].(string); ok {
		return []byte(v)
	}
	return []byte("")
}

func (p paramsData) GetBool(key string) bool {
	if p == nil {
		return false
	}

	if v, ok := p[key].(bool); ok {
		return v
	}
	return false
}

func (p paramsData) GetFloat64(key string) float64 {
	if p == nil {
		return 0.0
	}

	if v, ok := p[key].(float64); ok {
		return v
	}
	return 0.0
}

func (p paramsData) GetInterface(key string) interface{} {
	if p == nil {
		return nil
	}

	if v, ok := p[key]; ok {
		return v
	}
	return nil
}

func (p paramsData) GetInterfaces(key string) []interface{} {
	if p == nil {
		return nil
	}

	if v, ok := p[key].([]interface{}); ok {
		return v
	}
	return nil
}

func (p paramsData) GetAll() map[string]interface{} {
	return p
}
