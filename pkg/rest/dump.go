package rest

import (
	"encoding/json"
	"net/http"
	"strings"
)

var (
	doNotDumpParams = []string{"access_token"}
)

type requestDump struct {
	Method  string            `json:"method"`
	Host    string            `json:"host"`
	Path    string            `json:"path"`
	Params  map[string]string `json:"params"`
	Headers map[string]string `json:"headers"`
}

type responseDump struct {
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Body       interface{}       `json:"body"`
}

func DumpRequestAsJson(req *http.Request) string {
	host := req.Host
	if host == "" && req.URL != nil {
		host = req.URL.Host
	}
	params := copyValuesMap(req.URL.Query())
	for _, k := range doNotDumpParams {
		delete(params, k)
	}
	headers := copyValuesMap(req.Header)
	dumpBytes, _ := json.Marshal(requestDump{
		Method:  req.Method,
		Path:    req.URL.Path,
		Host:    host,
		Headers: headers,
		Params:  params,
	})
	return string(dumpBytes)
}

func DumpResponseAsJson(res http.Response) string {
	headers := copyValuesMap(res.Header)
	dumpBytes, _ := json.Marshal(responseDump{
		StatusCode: res.StatusCode,
		Headers:    headers,
		Body:       res.Body,
	})
	return string(dumpBytes)
}

func copyValuesMap(a map[string][]string) map[string]string {
	cp := map[string]string{}
	for k, v := range a {
		cp[k] = strings.Join(v, ",")
	}
	return cp
}
