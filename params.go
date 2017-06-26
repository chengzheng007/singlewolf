package singlewolf

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
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

	v, ok := p[key]
	if !ok {
		return 0
	}
	var i64 int64
	// map interface{} 默认将数字类型设置为float64
	switch inst := v.(type) {
	case float64:
		i64 = int64(inst)
	case string:
		i64, _ = strconv.ParseInt(inst, 10, 64)
	default:
		panic("unknow type")
	}
	return i64
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
	v, ok := p[key]
	if !ok {
		return 0.0
	}
	var f64 float64
	switch inst := v.(type) {
	case float64:
		f64 = inst
	case string:
		f64, _ = strconv.ParseFloat(inst, 10)
	default:
		panic("unknow type")
	}
	return f64
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

// GetInterfaces get slice of interface{}, then you can transfor type you want
func (p paramsData) GetInterfaces(key string) []interface{} {
	if p == nil {
		return nil
	}

	if v, ok := p[key].([]interface{}); ok {
		return v
	}
	return nil
}

// GetArray get array of paramsData, every element in the array you can call its own get-method
func (p paramsData) GetArray(key string) []paramsData {
	if p == nil {
		return nil
	}

	var list []paramsData
	for _, v := range p.GetInterfaces(key) {
		m, ok := v.(map[string]interface{})
		if ok {
			list = append(list, m)
		}
	}
	return list
}

// GetMap get an object
func (p paramsData) GetMap(key string) paramsData {
	if p == nil {
		return nil
	}
	var data paramsData
	if v, ok := p.GetInterface(key).(map[string]interface{}); ok {
		data = v
	}
	return data
}

// paramsData get all request data
func (p paramsData) GetAll() map[string]interface{} {
	return p
}
