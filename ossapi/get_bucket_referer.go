/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"encoding/xml"
	"path"
	"strconv"
)

/*
//redefine in put_bucket_referer
type RefererListInfo struct {
	Referer []string
}

type RefererConfigurationInfo struct {
	XMLName           xml.Name        `xml:"RefererConfiguration"`
	AllowEmptyReferer bool            `xml:"AllowEmptyReferer"`
	RefererList       RefererListInfo `xml:"RefererList"`
}
*/

func GetBucketReferer(name, location string) (info *RefererConfigurationInfo, ossapiError *Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name) + "/"
	req := &Request{
		Host:     host,
		Path:     "/?referer",
		Method:   "GET",
		Resource: resource,
		SubRes:   []string{"referer"}}
	rsp, err := req.Send()
	if err != nil {
		if _, ok := err.(*Error); !ok {
			Logger.Error("GetService's Send Error:%s", err.Error())
			ossapiError = OSSAPIError
			return
		}
	}
	if rsp.Result != ESUCC {
		ossapiError = err.(*Error)
		return
	}
	bodyLen, err := strconv.Atoi(rsp.httpRsp.Header["Content-Length"][0])
	if err != nil {
		Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = OSSAPIError
		return
	}
	body := make([]byte, bodyLen)
	rsp.httpRsp.Body.Read(body)
	info = new(RefererConfigurationInfo)
	err = xml.Unmarshal(body, info)
	if err != nil {
		Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = OSSAPIError
		return
	}
	return
}
