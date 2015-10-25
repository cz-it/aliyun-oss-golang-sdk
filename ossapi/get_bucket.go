/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"encoding/xml"
	"path"
	"strconv"
	"strings"
)

/*
// redefined in get_service.go
type Owner struct {
	ID          string
	DisplayName string
}
*/

type ContentInfo struct {
	Key          string
	LastModified string
	ETag         string
	Type         string
	Size         string
	StorageClass string
	Owner        Owner
}

type CommonInfo struct {
	Prefix string
}

type BucktsInfo struct {
	XMLName        xml.Name `xml:"ListBucketResult"`
	Name           string   `xml:"Name"`
	Prefix         string   `xml:"Prefix"`
	Marker         string   `xml:"Marker"`
	MaxKeys        int      `xml:"MaxKeys"`
	EncodingType   string   `xml:"encoding-type"`
	IsTruncated    bool     `xml:"IsTruncated"`
	Contents       []ContentInfo
	CommonPrefixes CommonInfo `xml:"CommonPrefixes"`
}

func GetBucket(name, location string, prefix, marker, delimiter, encodingType string, maxKeys int) (info *BucktsInfo, ossapiError *Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name) + "/"
	urlPath := "/"
	var args []string
	if prefix != "" {
		args = append(args, "prefix="+prefix)
	}
	if marker != "" {
		args = append(args, "marker="+marker)
	}
	if delimiter != "" {
		args = append(args, "delimiter="+delimiter)
	}
	if encodingType != "" {
		args = append(args, "encoding-type"+encodingType)
	}
	if maxKeys > 0 {
		args = append(args, "max-keys="+strconv.FormatUint(uint64(maxKeys), 10))
	}
	if args != nil {
		urlPath += "?" + strings.Join(args, "&")
	}
	req := &Request{
		Host:     host,
		Path:     urlPath,
		Method:   "GET",
		Resource: resource}
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
	info = new(BucktsInfo)
	err = xml.Unmarshal(body, info)
	if err != nil {
		Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = OSSAPIError
		return
	}
	return
}
