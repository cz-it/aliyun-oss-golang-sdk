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
// redifine in put_bucket_logging.go
type LoggingInfo struct {
	TargetBucket string
	TargetPrefix string
}
*/

type LoggingStatus struct {
	XMLName        xml.Name `xml:"BucketLoggingStatus"`
	LoggingEnabled LoggingInfo
}

func GetBucketLogging(name, location string) (info *LoggingInfo, ossapiError *Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name) + "/"
	req := &Request{
		Host:     host,
		Path:     "/?logging",
		Method:   "GET",
		Resource: resource,
		SubRes:   []string{"logging"}}
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
	status := new(LoggingStatus)
	err = xml.Unmarshal(body, status)
	if err != nil {
		Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = OSSAPIError
		return
	}
	info = &status.LoggingEnabled
	return
}
