/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
)

const (
	D_COPY    = "COPY"
	D_REPLACE = "REPLACE"
)

type CopyConditionInfo struct {
	ETAG         string
	Date         string
	LastModify   string
	LastUnModify string
}

type CopyInfo struct {
	ObjectName string
	BucketName string
	Location   string
	Source     string
	Directive  string
	Encryption string
	ACL        string
}

type CopyResultInfo struct {
	XMLName      xml.Name `xml:"CopyObjectResult"`
	ETag         string   `xml:"ETag"`
	LastModified string   `xml:"LastModified"`
}

func CopyObject(copyInfo *CopyInfo, copyConnInfo *CopyConditionInfo) (info *CopyResultInfo, ossapiError *ossapi.Error) {
	if copyInfo == nil {
		ossapiError = ossapi.ArgError
		return
	}
	resource := path.Join("/", copyInfo.BucketName, copyInfo.ObjectName)
	host := copyInfo.BucketName + "." + copyInfo.Location + ".aliyuncs.com"
	req := &ossapi.Request{
		Host:     host,
		Path:     "/" + copyInfo.ObjectName,
		Method:   "PUT",
		Resource: resource}
	if copyConnInfo != nil {
		req.AddXOSS("x-oss-copy-source-if-match", copyConnInfo.ETAG)
		req.AddXOSS("x-oss-copy-source-if-none-match", copyConnInfo.Date)
		req.AddXOSS("x-oss-copy-source-if-unmodified-since", copyConnInfo.LastUnModify)
		req.AddXOSS("x-oss-copy-source-if-modified-since", copyConnInfo.LastModify)
	}
	if copyInfo.ObjectName != "" {
		req.AddXOSS("x-oss-copy-source", copyInfo.Source)
	}
	if copyInfo.Directive != "" {
		req.AddXOSS("x-oss-metadata-directive", copyInfo.Directive)
	}
	if copyInfo.Encryption != "" {
		req.AddXOSS("x-oss-server-side-encryption", copyInfo.Encryption)
	}
	if copyInfo.ACL != "" {
		req.AddXOSS("x-oss-object-acl", copyInfo.ACL)
	}

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
	return
}
