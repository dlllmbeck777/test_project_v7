package response

import (
	"encoding/json"
	"reflect"
	"time"
)

type ErrorItem struct {
	Key           int    `json:"key"`
	AppMessage    string `json:"app_message"`
	SystemMessage string `json:"system_message"`
}

type Response struct {
	Status      bool          `json:"status"`
	Errors      []ErrorItem   `json:"errors"`
	Values      []interface{} `json:"values"`
	TmRequest   string        `json:"tm_req"`
	TmRequestSt time.Time     `json:"-"`
}

func InitResp() *Response {
	return &Response{
		TmRequestSt: time.Now(),
	}
}

func (r *Response) SetError(key int, mess string, systemess string) *Response {
	r.Errors = append(r.Errors, ErrorItem{
		Key:           key,
		AppMessage:    mess,
		SystemMessage: systemess,
	})
	return r
}

func (r *Response) SetValue(val interface{}) *Response {
	r.Values = append(r.Values, val)
	return r
}

func (r *Response) SetValues(vals interface{}) *Response {
	r.Values = r.InterfaceSlice(vals)
	return r
}

func (r *Response) InterfaceSlice(in interface{}) []interface{} {
	s := reflect.ValueOf(in)
	if s.Kind() != reflect.Slice {
		return []interface{}{}
	}
	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

func (r *Response) FormResponse() *Response {
	if len(r.Errors) > 0 {
		r.Status = false
	} else {
		r.Status = true
	}

	// not null array json
	if len(r.Values) == 0 {
		r.Values = []interface{}{}
	}

	// not null array json
	if len(r.Errors) == 0 {
		r.Errors = []ErrorItem{}
	}
	return r
}

func (r *Response) Json() []byte {
	r.TmRequest = time.Now().Sub(r.TmRequestSt).String()
	if bts, err := json.Marshal(r); err == nil {
		return bts
	}
	return []byte{}
}
