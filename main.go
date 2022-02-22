package telegraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/anaskhan96/soup"
)

var baseURL string = "https://api.telegra.ph/"

// type Telegraph struct{}

type ResultStruct struct {
	Ok     bool          `json:"ok"`
	Result AllValueTypes `json:"result,omitempty"`
	Error  string        `json:"error,omitempty"`
}

func callAPI(route string, params string) (*AllValueTypes, error) {
	var url string
	if params == "" {
		url = baseURL + route
	} else {
		url = baseURL + route + "?" + params
	}
	var jsonData ResultStruct
	r, e := http.Get(url)
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

func isZeroOfType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func Prettify(t interface{}) string {
	b, _ := json.MarshalIndent(t, "", "    ")
	return string(b)
}

// ToDo : Use a better, nope, much better way. But yeah, it's reliable.

func HTMLToNodeString(html string) string {
	doc := soup.HTMLParse(html)
	children := doc.Find("body").Children()
	my := []string{}
	for _, c := range children {
		str := "{"
		tagStr := ""
		if c.NodeValue != "" {
			tagStr = fmt.Sprintf(`"%v"`, c.NodeValue)
		}
		childrenStr := ""
		if c.Text() != "" {
			childrenStr = fmt.Sprintf(`["%v"]`, c.Text())
		}
		// 	str += `"` + c.NodeValue + `"`
		// }
		attrsStr := ""
		if c.Attrs() != nil {
			attrsList := []string{}
			for k, v := range c.Attrs() {
				attrsList = append(attrsList, fmt.Sprintf(`"%v":"%v"`, k, v))
			}
			attrsStr = fmt.Sprintf(`{%v}`, strings.Join(attrsList, ","))
		}
		if tagStr != "" {
			if c.Text() == "" && c.NodeValue != "img" {
				str = fmt.Sprintf(`"%v"`, c.NodeValue)
				my = append(my, str)
				continue
			} else {
				str += fmt.Sprintf(`"tag":%v,`, tagStr)
			}
		}
		if attrsStr != "" {
			str += fmt.Sprintf(`"attrs":%v,`, attrsStr)
		}
		if childrenStr != "" {
			str += fmt.Sprintf(`"children":%v,`, childrenStr)
		}
		str = str[0 : len(str)-1]
		str += "}"
		my = append(my, str)
	}
	str := strings.Join(my, ",")
	return "[" + str + "]"
}
