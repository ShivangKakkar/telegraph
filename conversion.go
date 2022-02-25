package telegraph

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
)

// ToDo : Use a better way, nope, much better way.

func HTMLToNode(html string) []Node {
	// Indentation shouldn't matter
	html = strings.ReplaceAll(html, "\r\n", " ")
	html = strings.ReplaceAll(html, "  ", "")
	tag := soup.HTMLParse(html).Find("body").Children()
	nodes := soupToNode(tag)
	return nodes
}

func soupToNode(children []soup.Root) []Node {
	nodes := []Node{}
	for _, c := range children {
		var el NodeElement
		tag := c.NodeValue
		var children []Node
		if c.Children() != nil {
			children = soupToNode(c.Children())
		}
		attrs := make(map[string]string)
		if c.Attrs() != nil {
			for k, v := range c.Attrs() {
				attrs[k] = v
			}
		}
		if tag != "" {
			if !strings.Contains(c.HTML(), "<") {
				if strings.TrimSpace(c.NodeValue) != "" {
					var val string
					if len(nodes) == 0 && string(c.NodeValue[0]) == " " {
						val = c.NodeValue[1:] // Usually formatters add an extra " ".
					} else {
						val = c.NodeValue
					}
					nodes = append(nodes, val)
					continue
				}
			} else {
				el.Tag = tag
			}
		}
		if attrs != nil {
			el.Attrs = attrs
		}
		if children != nil {
			el.Children = children
		}
		nodes = append(nodes, el)
	}
	return nodes
}

func NodeToQueryString(nodes []Node) string {
	str := "["
	for _, node := range nodes {
		switch n := node.(type) {
		case NodeElement:
			bs, _ := json.Marshal(n)
			str += string(bs) + ","
		case string:
			str += fmt.Sprintf(`"%v",`, n)
		}
	}
	str = str[0:len(str)-1] + "]"
	return str
}
