/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"encoding/xml"
)

type Bucket struct {
	Name         string
	CreationDate string
	Location     string
}

type Buckets struct {
	Bucket []Bucket
}

type Owner struct {
	ID          string
	DisplayName string
}

type BucketsInfo struct {
	XMLName     xml.Name `xml:"ListAllMyBucketsResult"`
	Prefix      string   `xml:"Prefix"`
	Marker      string   `xml:"Marker"`
	MaxKeys     int      `xml:"MaxKeys"`
	IsTruncated bool     `xml:"IsTruncated"`
	NextMarker  string   `xml:"NextMarker"`
	Owner       Owner    `xml:"Owner"`
	Buckets     Buckets  `xml:"Buckets"`
}

func GetServiceWith(prefix, marker string, maxKeys int) (bucketsInfo *BucketsInfo, error *Error) {
	args := ""
	path := "/"
	if "" != prefix {
		args += "prefix=" + prefix
	}
	if "" != marker {
		args += "marker=" + marker
	}
	if 0 < maxKeys && maxKeys <= 1000 {
		args += "maxkeys=" + string(maxKeys)
	}

	if "" != args {
		path += "?" + args
	}
	Logger.Info("path is %s", path)
	req := &Request{Host: "oss.aliyuncs.com", Path: "/", Method: "GET", Resource: "/"}
	rsp, err := req.Send()
	if err != nil {
		if _, ok := err.(*Error); !ok {
			Logger.Error("GetService's Send Error:%s", err.Error())
			error = &Error{ErrNo: EAPI,
				ErrMsg:       "OSSAPI SDK's Inner Error,You Can Find More Details In Log Files",
				ErrDetailMsg: "OSSAPI SDK's Inner Error,You Can Find More Details In Log Files"}
			return
		}
	}
	if rsp.Result != ESUCC {
		error = err.(*Error)
		return
	}
	body := make([]byte, 10000)
	rsp.httpRsp.Body.Read(body)
	bucketsInfo = new(BucketsInfo)
	xml.Unmarshal(body, bucketsInfo)
	return
}

func GetService() (bucketsInfo *BucketsInfo, err *Error) {
	bucketsInfo, err = GetServiceWith("", "", 0)
	return
}
