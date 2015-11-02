/**
* Author: CZ cz.theng@gmail.com
 */

// service package support list action for bucket
// or list bucket packge

package service

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"strconv"
	"strings"
)

// bucket info
type Bucket struct {
	Name         string
	CreationDate string
	Location     string
}

// buckets
type Buckets struct {
	Bucket []Bucket
}

// owner info
type Owner struct {
	ID          string
	DisplayName string
}

// buckets info return form xml
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

// GetService will list bucket of a account
// buckets with prefix will be return if prefix is not ""
// marker mark the split for return
// at moste maxKeys will return ,default is 100
func GetService(prefix, marker string, maxKeys int) (bucketsInfo *BucketsInfo, ossapiError *ossapi.Error) {
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
	req := &ossapi.Request{
		Host:     "oss.aliyuncs.com",
		Path:     path,
		Method:   "GET",
		Resource: "/"}
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

// GetServiceDefault list all buckets with no prefix ,no marker and maxkeys to 100
func GetServiceDefault() (bucketsInfo *BucketsInfo, err *ossapi.Error) {
	bucketsInfo, err = GetService("", "", 0)
	return
}
