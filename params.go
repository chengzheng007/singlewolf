package singlewolf

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	if val, ok := p.data[key]; ok {
		v, _ := val.(string)
		return v
	}
	return ""
}

func (p *paramsData) GetAll() interface{} {
	return p.data
}
