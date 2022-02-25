package telegraph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

var baseURL string = "https://api.telegra.ph/"

type response struct {
	Ok     bool          `json:"ok"`
	Result AllValueTypes `json:"result,omitempty"`
	Error  string        `json:"error,omitempty"`
}

func Get(route string, opts interface{}) (*AllValueTypes, error) {
	vs := url.Values{}
	v := reflect.ValueOf(opts)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i).Interface()
		key := t.Field(i).Name
		if !isZeroOfType(val) {
			var str string
			switch vt := val.(type) {
			case int64:
				str = strconv.FormatInt(vt, 10)
			case bool:
				str = strconv.FormatBool(vt)
			case []string:
				str = fmt.Sprintf(`["%v"]`, strings.Join(vt, `","`))
			case []Node:
				str = NodeToQueryString(vt)
			default:
				str = vt.(string)
			}
			vs.Add(snaking(key), str)
		}
	}
	u := baseURL + route + "?" + vs.Encode()
	var jsonData response
	r, e := http.Get(u)
	bs, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if e != nil {
		return &AllValueTypes{}, e
	}
	json.Unmarshal(bs, &jsonData)
	if !jsonData.Ok {
		return &AllValueTypes{}, fmt.Errorf(jsonData.Error)
	}
	return &jsonData.Result, nil
}

func Post(route string, opts interface{}) (*AllValueTypes, error) {
	u := baseURL + route
	var jsonData response
	bs, _ := json.Marshal(opts)
	r, e := http.Post(u, "application/json", bytes.NewReader(bs))
	bs, _ = ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if e != nil {
		return &AllValueTypes{}, e
	}
	json.Unmarshal(bs, &jsonData)
	if !jsonData.Ok {
		return &AllValueTypes{}, fmt.Errorf(jsonData.Error)
	}
	return &jsonData.Result, nil
}

// ------------------- Helpers ------------------ //

func isZeroOfType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func Prettify(t interface{}) string {
	b, _ := json.MarshalIndent(t, "", "    ")
	return string(b)
}

func snaking(s string) string {
	c := s[1:]
	for _, char := range c {
		if unicode.IsUpper(char) {
			c = strings.Replace(c, string(char), "_"+strings.ToLower(string(char)), -1)
		}
	}
	return strings.ToLower(string(s[0])) + c
}
