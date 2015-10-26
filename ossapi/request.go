/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"bytes"
	"encoding/xml"
	//	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	Host      string
	Path      string
	Date      string
	Method    string
	CntType   string
	Resource  string
	SubRes    []string
	XOSSes    map[string]string
	Body      []byte
	ExtHeader map[string]string
	RspHeader map[string]string

	httpReq *http.Request
}

func (req *Request) Send() (rsp *Response, err error) {
	URL := "http://"
	URL += path.Join(req.Host, req.Path)
	//fmt.Println("URL:", URL)
	req.httpReq, err = http.NewRequest(req.Method, URL, nil)
	if err != nil {
		Logger.Error("http.NewRequest(req.Method,URL, nil) Error:%s", err.Error())
		return
	}
	req.httpReq.ProtoMinor = 1
	req.Date = time.Now().UTC().Format(DATE_FMT)
	req.httpReq.Header.Add("Date", req.Date)
	auth, err := req.Auth()
	if err != nil {
		Logger.Error("req.Auth() Error:%s", err.Error)
		return
	}
	req.httpReq.Header.Add("Authorization", auth)
	for k, v := range req.XOSSes {
		req.httpReq.Header.Add(k, v)
	}
	if req.Body != nil {
		req.httpReq.Header.Add("Content-Length", strconv.FormatUint(uint64(len(req.Body)), 10))
		req.httpReq.Header.Add("Content-Type", req.CntType)
		var cntMd5 string
		if req.Body != nil {
			cntMd5, err = Base64AndMd5(req.Body)
			if err != nil {
				Logger.Error("Base64AndMd5(req.Body) Error:%s", err.Error())
				return
			}
		}
		req.httpReq.Header.Add("Content-MD5", cntMd5)
		req.httpReq.Body = ioutil.NopCloser(bytes.NewReader(req.Body))
	}
	if req.ExtHeader != nil {
		for k, v := range req.ExtHeader {
			req.httpReq.Header.Add(k, v)
		}
	}
	//fmt.Println("Req head:", req.httpReq.Header)
	httprsp, err := httpClient.Do(req.httpReq)
	if err != nil {
		Logger.Error("httpClient.Do(req.httpReq) Error:%s", err.Error())
		return
	}
	rsp = &Response{HttpRsp: httprsp}
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
		//fmt.Println("body:", string(body))
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
		cntMd5, err = Base64AndMd5(req.Body)
		if err != nil {
			Logger.Error("Base64AndMd5(req.Body) Error:%s", err.Error())
			return
		}
	}
	sigStr += cntMd5 + "\n"
	sigStr += req.CntType + "\n"
	sigStr += req.Date + "\n"
	var ossHeaders []string
	var ossHeadersStr string
	if req.XOSSes != nil {
		var ossHeaderKeys []string
		//fmt.Println("req.XOSSes : ", req.XOSSes)
		for k, _ := range req.XOSSes {
			ossHeaderKeys = append(ossHeaderKeys, strings.ToLower(k))
		}
		sort.Sort(sort.StringSlice(ossHeaderKeys))
		//fmt.Println("ossheaderKeys:", ossHeaderKeys)
		for i := 0; i < len(ossHeaderKeys); i++ {
			//fmt.Println("ossHeaderKeys[i]:", i, ": ", ossHeaderKeys[i])
			ossHeaders = append(ossHeaders, ossHeaderKeys[i]+":"+req.XOSSes[ossHeaderKeys[i]])
		}
		ossHeadersStr = strings.Join(ossHeaders, "\n")
		ossHeadersStr += "\n"
	}
	var subResStr string
	var resourcesStr string
	if req.SubRes != nil {
		subRes := req.SubRes
		sort.Sort(sort.StringSlice(subRes))
		subResStr = strings.Join(subRes, "&")
		subResStr = "?" + subResStr
	}
	resourcesStr = req.Resource + subResStr
	sigStr += ossHeadersStr + resourcesStr
	//fmt.Println("sigStr:", sigStr)
	sig, err = Base64AndHmacSha1([]byte(accessKeySecret), []byte(sigStr))
	if err != nil {
		Logger.Error("sig, err = Base64AndHmacSha1([]byte(accessKeySecret), []byte(sigStr)) Error:%s", err.Error())
		return
	}
	return
}
func (req *Request) AddXOSS(key string, value string) {
	if req.XOSSes == nil {
		req.XOSSes = make(map[string]string)
	}
	req.XOSSes[strings.TrimSpace(key)] = strings.TrimSpace(value)
}
