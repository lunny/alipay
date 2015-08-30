package alipay

import (
	"bytes"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type Params map[string]string

func ParseParams(query string) (Params, error) {
	params := make(Params)
	err := parseParams(params, query)
	return params, err
}

func parseParams(m Params, query string) (err error) {
	for query != "" {
		fmt.Println("query:", query)
		key := query
		if i := strings.IndexAny(key, "&;"); i >= 0 {
			key, query = key[:i], key[i+1:]
		} else {
			query = ""
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, "="); i >= 0 {
			key, value = key[:i], key[i+1:]
		}
		key, err1 := url.QueryUnescape(key)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		value, err1 = url.QueryUnescape(value)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		m[key] = value
	}
	return err
}

func (v Params) Encode(escape bool) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}

		if escape {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(vs))
		} else {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(vs)
		}
	}
	return buf.String()
}
