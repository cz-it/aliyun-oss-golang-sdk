/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"net/http"
	"path"
)

type Requester interface {
	AddXOSS(key string, value string)
	Signature() (sig string, err error)
	Send() (rsp *Response, err error)
}

type Request struct {
	Host string
	Path string

	request *http.Request
}

func (req *Request) Send() (rsp *Response, err error) {
	URL := "http://" + path.Join(req.Host, req.Path)
	req.request, err = http.NewRequest("GET", URL, nil)
	if err != nil {
		return
	}
	httprsp, err := httpClient.Do(req.request)
	if err != nil {
		return
	}
	rsp = &Response{httpRsp: httprsp}
	print("httpresponse:", httprsp.Header)
	return
}

func (req *Request) Signature() (sig string, err error) {
	return
}

func (req *Request) AddXOSS(key string, value string) {

}
