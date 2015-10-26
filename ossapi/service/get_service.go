/**
* Author: CZ cz.theng@gmail.com
 */

package service

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"strconv"
	"strings"
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

func GetServiceWith(prefix, marker string, maxKeys int) (bucketsInfo *BucketsInfo, ossapiError *ossapi.Error) {
	var args []string
	path := "/"
	if "" != prefix {
		args = append(args, "prefix="+prefix)
	}
	if "" != marker {
		args = append(args, "marker="+marker)
	}
	if 0 < maxKeys && maxKeys <= 1000 {
		args = append(args, "maxkeys="+strconv.FormatUint(uint64(maxKeys), 10))
	}

	if args != nil {
		path += "?" + strings.Join(args, "&")
	}
	req := &ossapi.Request{Host: "oss.aliyuncs.com", Path: path, Method: "GET", Resource: "/"}
	rsp, err := req.Send()
	if err != nil {
		if _, ok := err.(*ossapi.Error); !ok {
			ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
			ossapiError = ossapi.OSSAPIError
			return
		}
	}
	if rsp.Result != ossapi.ESUCC {
		ossapiError = err.(*ossapi.Error)
		return
	}
	bodyLen, err := strconv.Atoi(rsp.HttpRsp.Header["Content-Length"][0])
	if err != nil {
		ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	body := make([]byte, bodyLen)
	rsp.HttpRsp.Body.Read(body)
	bucketsInfo = new(BucketsInfo)
	xml.Unmarshal(body, bucketsInfo)
	return
}

func GetService() (bucketsInfo *BucketsInfo, err *ossapi.Error) {
	bucketsInfo, err = GetServiceWith("", "", 0)
	return
}
