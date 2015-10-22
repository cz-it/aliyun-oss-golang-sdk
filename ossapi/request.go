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
	Host     string
	Path     string
	Date     string
	Method   string
	CntType  string
	Resource string

	Body []byte

	httpReq *http.Request
}

func (req *Request) Send() (rsp *Response, err error) {
	URL := "http://" + path.Join(req.Host, req.Path)
	req.httpReq, err = http.NewRequest(req.Method, URL, nil)
	if err != nil {
		Logger.Error("http.NewRequest(req.Method,URL, nil) Error:%s", err.Error())
		return
	}
	req.httpReq.ProtoMinor = 1
	req.Date = time.Now().UTC().Format(DATE_FMT)
	req.httpReq.Header.Add("Date", req.Date)
	//req.httpreq.Header.Add("Host", req.Host)
	auth, err := req.Auth()
	if err != nil {
		Logger.Error("req.Auth() Error:%s", err.Error)
		return
	}
	req.httpReq.Header.Add("Authorization", auth)
	//fmt.Println("Req head:", req.httpreq.Header)
	httprsp, err := httpClient.Do(req.httpReq)
	if err != nil {
		Logger.Error("httpClient.Do(req.httpReq) Error:%s", err.Error())
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
			Logger.Error("httprsp.Body.Read(body) Error:%s", err.Error())
			return
		}
		err = xml.Unmarshal(body, rstErr)
		if err != nil {
			Logger.Error("xml.Unmarshal(body, rstErr) Error:%s", err.Error())
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

func (req *Request) Auth() (authStr string, err error) {
	authStr = "OSS " + accessKeyID + ":"
	sigStr, err := req.Signature()
	if err != nil {
		Logger.Error("sigStr, err := req.Signature() Error:%s", err.Error())
		return
	}
	authStr += sigStr
	return
}

func (req *Request) Signature() (sig string, err error) {
	sigStr := req.Method + "\n"
	var cntMd5 string
	if req.Body != nil {

	}
	sigStr += cntMd5 + "\n"
	sigStr += req.CntType + "\n"
	sigStr += req.Date + "\n"
	var ossHeaders string
	resources := req.Resource
	sigStr += ossHeaders + resources
	sig, err = Base64AndHmacSha1([]byte(accessKeySecret), []byte(sigStr))
	if err != nil {
		Logger.Error("sig, err = Base64AndHmacSha1([]byte(accessKeySecret), []byte(sigStr)) Error:%s", err.Error())
		return
	}
	return
}
func (req *Request) AddXOSS(key string, value string) {
}
