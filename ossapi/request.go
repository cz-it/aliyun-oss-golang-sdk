/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"encoding/xml"
	"io"
	"net/http"
	"path"
	"strconv"
	"time"
)

type Request struct {
	Host string
	Path string
	Date string

	httpReq *http.Request
}

type Signaturer func(*Request) (string, error)

func (req *Request) Send(signature Signaturer) (rsp *Response, err error) {
	if signature == nil {
		err = EARG
		return
	}
	URL := "http://" + path.Join(req.Host, req.Path)
	req.httpReq, err = http.NewRequest("GET", URL, nil)
	if err != nil {
		return
	}
	req.httpReq.ProtoMinor = 1
	req.Date = time.Now().UTC().Format(DATE_FMT)
	req.httpReq.Header.Add("Date", req.Date)
	//req.httpreq.Header.Add("Host", req.Host)
	auth, err := req.Auth(signature)
	if err != nil {
		return
	}
	req.httpReq.Header.Add("Authorization", auth)
	//fmt.Println("Req head:", req.httpreq.Header)
	httprsp, err := httpClient.Do(req.httpReq)
	if err != nil {
		return
	}
	rsp = &Response{httpRsp: httprsp}
	if httprsp.StatusCode/100 == 4 || httprsp.StatusCode/100 == 5 {
		var cntLen int
		rstErr := &Error{HttpStatus: httprsp.StatusCode, ErrNo: ENone, ErrMsg: "None", ErrDetailMsg: "None"}
		cntLen, err = strconv.Atoi(httprsp.Header["Content-Length"][0])
		if err != nil {
			cntLen = 1024
		}
		body := make([]byte, cntLen*10)
		_, err = httprsp.Body.Read(body)
		if err != nil && err != io.EOF {
			return
		}
		err = xml.Unmarshal(body, rstErr)
		if err != nil {
			return
		}
		rstErr.ErrDetailMsg = string(body)
		err = rstErr
		rsp.Result = EFAIL
		return
	} else if httprsp.StatusCode/100 == 2 {
		rsp.Result = ESUCC
	} else {
		rsp.Result = EUNKNOWN
	}
	return
}

func (req *Request) Auth(signature Signaturer) (authStr string, err error) {
	if signature == nil {
		err = EARG
		return
	}
	authStr = "OSS " + accessKeyID + ":"
	sigStr, err := signature(req)
	if err != nil {
		return
	}
	authStr += sigStr
	return
}

func (req *Request) AddXOSS(key string, value string) {
}
