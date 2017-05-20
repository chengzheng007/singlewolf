package singlewolf

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

type paramsData struct {
	data map[string]interface{}
}

func getRequestParams(r *http.Request) paramsData {
	params := paramsData{data: make(map[string]interface{})}

	if r == nil {
		return params
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logf("ioutil.ReadAll(r.Body) error(%v)", err)
		return params
	}
	defer r.Body.Close()

	if err := json.Unmarshal(buf, &params.data); err != nil {
		logf("json.Unmarshal(%s, &params.data) error(%v)", buf, err)
		return params
	}

	return params
}

func (p *paramsData) GetString(key string) string {
	if p.data == nil {
		return ""
	}
	if v, ok := p.data[key].(string); ok {
		return v
	}
	return ""
}

func (p *paramsData) GetInt64(key string) int64 {
	if p.data == nil {
		return 0
	}
	// map interface{} 默认将数字类型设置为float64
	if v, ok := p.data[key].(float64); ok {
		return int64(reflect.ValueOf(v).Float())
	}
	return 0
}

func (p *paramsData) GetBytes(key string) []byte {
	if p.data == nil {
		return []byte("")
	}

	if v, ok := p.data[key].(string); ok {
		return []byte(v)
	}
	return []byte("")
}

func (p *paramsData) GetBool(key string) bool {
	if p.data == nil {
		return false
	}

	if v, ok := p.data[key].(bool); ok {
		return v
	}
	return false
}

func (p *paramsData) GetFloat64(key string) float64 {
	if p.data == nil {
		return 0.0
	}

	if v, ok := p.data[key].(float64); ok {
		return v
	}
	return 0.0
}

func (p *paramsData) GetInterface(key string) interface{} {
	if p.data == nil {
		return nil
	}

	if v, ok := p.data[key]; ok {
		return v
	}
	return nil
}

func (p *paramsData) GetAll() map[string]interface{} {
	return p.data
}
