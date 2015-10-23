/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"encoding/xml"
	"path"
)

type IndexInfo struct {
	Suffix string
}

type KeyInfo struct {
	Key string
}

type WebsiteInfo struct {
	XMLName       xml.Name  `xml:"WebsiteConfiguration"`
	IndexDocument IndexInfo `xml:"IndexDocument"`
	ErrorDocument KeyInfo   `xml:"ErrorDocument"`
}

func SetBucketWebsite(name, location, indexPage, errorPage string) (ossapiError *Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name)
	indexInfo := IndexInfo{Suffix: indexPage}
	var keyInfo KeyInfo
	keyInfo = KeyInfo{Key: errorPage}
	var info WebsiteInfo
	if "" == errorPage {
		info = WebsiteInfo{IndexDocument: indexInfo}
	} else {
		info = WebsiteInfo{IndexDocument: indexInfo, ErrorDocument: keyInfo}
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
		Path:     "/?website",
		Method:   "PUT",
		Resource: resource + "/",
		SubRes:   []string{"website"},
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
