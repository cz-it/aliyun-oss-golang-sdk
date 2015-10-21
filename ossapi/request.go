/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"fmt"
	"net/http"
	"path"
	"time"
)

type Requester interface {
	AddXOSS(key string, value string)
	Signature() (sig string, err error)
	Send() (rsp *Response, err error)
}

type Request struct {
	Host string
	Path string
	Date string

	httpreq *http.Request
}

func (req *Request) Send() (rsp *Response, err error) {
	URL := "http://" + path.Join(req.Host, req.Path)
	req.httpreq, err = http.NewRequest("GET", URL, nil)
	if err != nil {
		return
	}
	req.httpreq.ProtoMinor = 1
	req.Date = time.Now().UTC().Format(DATE_FMT)
	req.httpreq.Header.Add("Date", req.Date)
	//req.httpreq.Header.Add("Host", req.Host)
	auth, err := req.Auth()
	if err != nil {
		return
	}
	req.httpreq.Header.Add("Authorization", auth)
	//fmt.Println("Req head:", req.httpreq.Header)
	httprsp, err := httpClient.Do(req.httpreq)
	if err != nil {
		return
	}
	rsp = &Response{httpRsp: httprsp}
	body := make([]byte, 10240)
	//fmt.Println("httpresponse:", httprsp.Header)
	//fmt.Println("httpresponse:", httprsp.Status)
	httprsp.Body.Read(body)
	fmt.Println("Body:", string(body))
	return
}

func (req *Request) Auth() (authStr string, err error) {
	authStr = "OSS " + accessKeyID + ":"
	sigStr, err := req.Signature()
	if err != nil {
		return
	}
	authStr += sigStr
	return
}

func (req *Request) Signature() (sig string, err error) {
	sigStr := "GET\n"
	cntMd5, err := Base64AndMd5([]byte(""))
	if err != nil {
		return
	}
	println(cntMd5)
	sigStr += "" + "\n"
	sigStr += "\n"
	sigStr += req.Date + "\n"
	sigStr += "/"
	fmt.Println("sigStr:", sigStr)
	sig, err = Base64AndHmacSha1([]byte(accessKeySecret), []byte(sigStr))
	return
}

func (req *Request) AddXOSS(key string, value string) {

}
