/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"encoding/xml"
	"path"
	"strconv"
)

type LocationInfo struct {
	XMLName  xml.Name `xml:"LocationConstraint"`
	Location string   `xml:",chardata"`
}

func GetBucketLocation(name string) (location string, ossapiError *Error) {
	host := name + ".oss.aliyuncs.com"
	resource := path.Join("/", name) + "/"
	req := &Request{
		Host:     host,
		Path:     "/?location",
		Method:   "GET",
		Resource: resource,
		SubRes:   []string{"location"}}
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
	bodyLen, err := strconv.Atoi(rsp.HttpRsp.Header["Content-Length"][0])
	if err != nil {
		Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = OSSAPIError
		return
	}
	body := make([]byte, bodyLen)
	rsp.HttpRsp.Body.Read(body)
	locationInfo := new(LocationInfo)
	err = xml.Unmarshal(body, locationInfo)
	if err != nil {
		Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = OSSAPIError
		return
	}
	location = locationInfo.Location
	return
}
