/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"encoding/xml"
	"path"
)

type RefererListInfo struct {
	Referer []string
}

type RefererConfigurationInfo struct {
	XMLName           xml.Name        `xml:"RefererConfiguration"`
	AllowEmptyReferer bool            `xml:"AllowEmptyReferer"`
	RefererList       RefererListInfo `xml:"RefererList"`
}

func SetBucketReferer(name, location string, enable bool, urls []string) (ossapiError *Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name)
	refersInfo := RefererListInfo{Referer: urls}
	var info RefererConfigurationInfo
	if urls == nil {
		info = RefererConfigurationInfo{AllowEmptyReferer: enable}
	} else {
		info = RefererConfigurationInfo{AllowEmptyReferer: enable, RefererList: refersInfo}
	}
	body, err := xml.Marshal(info)
	if err != nil {
		Logger.Error("err := xml.Marshal(Info) Error %s", err.Error())
		ossapiError = OSSAPIError
		return
	}
	body = append([]byte(xml.Header), body...)
	req := &Request{
		Host:     host,
		Path:     "/?referer",
		Method:   "PUT",
		Resource: resource + "/",
		SubRes:   []string{"referer"},
		Body:     body,
		CntType:  "application/xml"}
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

	return
}
