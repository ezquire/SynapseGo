package wrapper

import (
	"github.com/parnurzeal/gorequest"
)

/********** GLOBAL VARIABLES **********/
const version = "v3.1"

// const _url = "https://api.synapsefi.com/" + version
const _url = "https://uat-api.synapsefi.com/" + version

var goreq = gorequest.New()

// http methods used
const (
	GET   = "GET"
	POST  = "POST"
	PATCH = "PATCH"
)

/********** METHODS **********/

func request(method, url string, headers map[string]interface{}, params []string, data ...string) []byte {
	var req = gorequest.New()
	req = setMethod(method, url)
	req = setParams(req, params, data)
	req = setHeader(req, headers)

	res, body, errs := req.EndBytes()

	if len(errs) > 0 {
		errorLog(errs)
	}

	if res.StatusCode != 200 {
		handleHTTPError(read(body)["http_code"].(string))
	}

	return body
}

func setHeader(r *gorequest.SuperAgent, h map[string]interface{}) *gorequest.SuperAgent {

	for k := range h {
		r.Set(k, h[k].(string))
	}

	return r
}

func setParams(req *gorequest.SuperAgent, params, data []string) *gorequest.SuperAgent {
	var p, d string

	if len(params) > 0 {
		p = params[0]
	}

	if len(data) > 0 {
		d = data[0]
	}

	return req.Query(p).Send(d)
}

func setMethod(m, u string) *gorequest.SuperAgent {
	switch m {
	case POST:
		return goreq.Post(u)

	case PATCH:
		return goreq.Patch(u)

	default:
		return goreq.Get(u)
	}
}
