/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"encoding/xml"
	"path"
)

type LoggingInfo struct {
	TargetBucket string
	TargetPrefix string
}

type OpenLoggingInfo struct {
	XMLName        xml.Name    `xml:"BucketLoggingStatus"`
	LoggingEnabled LoggingInfo `xml:"LoggingEnabled"`
}

type CloseLoggingInfo struct {
	XMLName xml.Name `xml:"BucketLoggingStatus"`
}

func OpenBucketLogging(name, location, targetBucket, targetPrefix string) (error *Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name)
	info := LoggingInfo{TargetBucket: targetBucket, TargetPrefix: targetPrefix}
	openInfo := &OpenLoggingInfo{LoggingEnabled: info}
	body, err := xml.Marshal(openInfo)
	if err != nil {
		Logger.Error("err := xml.Marshal(openInfo) Error %s", err.Error())
		error = OSSAPIError
		return
	}
	body = append([]byte(xml.Header), body...)
	req := &Request{Host: host, Path: "/?logging", Method: "PUT", Resource: resource + "/", Body: body, CntType: "application/xml"}
	rsp, err := req.Send()
	if err != nil {
		if _, ok := err.(*Error); !ok {
			Logger.Error("GetService's Send Error:%s", err.Error())
			error = OSSAPIError
			return
		}
	}
	if rsp.Result != ESUCC {
		error = err.(*Error)
		return
	}
	return
}

func CloseBucketLogging(name, location string) (error *Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name)
	closeInfo := &CloseLoggingInfo{}
	body, err := xml.Marshal(closeInfo)
	if err != nil {
		Logger.Error("err := xml.Marshal(closeInfo) Error %s", err.Error())
		error = OSSAPIError
		return
	}
	body = append([]byte(xml.Header), body...)
	req := &Request{Host: host, Path: "/?logging", Method: "PUT", Resource: resource + "/", Body: body, CntType: "application/xml"}
	rsp, err := req.Send()
	if err != nil {
		if _, ok := err.(*Error); !ok {
			Logger.Error("GetService's Send Error:%s", err.Error())
			error = OSSAPIError
			return
		}
	}
	if rsp.Result != ESUCC {
		error = err.(*Error)
		return
	}
	return
}
