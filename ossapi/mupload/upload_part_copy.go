/**
* Author: CZ cz.theng@gmail.com
 */

package mupload

import (
	"encoding/xml"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/object"
	"path"
	"strconv"
)

//UploadPartCopyInfo is  copy part info
type UploadPartCopyInfo struct {
	ObjectName    string
	BucketName    string
	Location      string
	PartNumber    int
	UploadID      string
	SrcObjectName string
	SrcBucketName string
	SrcRangeBegin int
	SrcRangeEnd   int
}

//UploadPartCopyRstInfo  is  resoponse info
type UploadPartCopyRstInfo struct {
	XMLName      xml.Name `xml:"CopyObjectResult"`
	LastModified string   `xml:"LastModified"`
	ETag         string   `xml:"ETag"`
}

// Copy a data slice
// @param partInfo : source and dest info
// @param copyConnInfo : conditon to copy
// @return rstInfo : return response
// @return ossapiError : nil on success
func Copy(partInfo *UploadPartCopyInfo, copyConnInfo *object.CopyConditionInfo) (rstInfo *UploadPartCopyRstInfo, ossapiError *ossapi.Error) {
	if partInfo == nil {
		ossapiError = ossapi.ArgError
		return
	}
	resource := path.Join("/", partInfo.BucketName, partInfo.ObjectName)
	host := partInfo.BucketName + "." + partInfo.Location + ".aliyuncs.com"
	header := make(map[string]string)
	header["Content-Length"] = strconv.FormatUint(uint64(partInfo.SrcRangeEnd-partInfo.SrcRangeBegin), 10)
	req := &ossapi.Request{
		Host:      host,
		ExtHeader: header,
		Path:      "/" + partInfo.ObjectName + "?partNumber=" + strconv.FormatUint(uint64(partInfo.PartNumber), 10) + "uploadId=" + partInfo.UploadID,
		Method:    "PUT",
		//		Body:     partInfo.Data,
		//		CntType:  partInfo.CntType,
		SubRes:   []string{"partNumber=" + strconv.FormatUint(uint64(partInfo.PartNumber), 10) + "uploadId=" + partInfo.UploadID},
		Resource: resource}
	if copyConnInfo != nil {
		req.AddXOSS("x-oss-copy-source-if-match", copyConnInfo.ETAG)
		req.AddXOSS("x-oss-copy-source-if-none-match", copyConnInfo.Date)
		req.AddXOSS("x-oss-copy-source-if-unmodified-since", copyConnInfo.LastUnModify)
		req.AddXOSS("x-oss-copy-source-if-modified-since", copyConnInfo.LastModify)
	}
	if partInfo.SrcObjectName != "" && partInfo.SrcBucketName != "" {
		req.AddXOSS("x-oss-copy-source", path.Join("/", partInfo.SrcBucketName, partInfo.SrcObjectName))
	}

	if partInfo.SrcRangeBegin > 0 && partInfo.SrcRangeEnd > 0 {
		cntRange := "bytes=" + strconv.FormatUint(uint64(partInfo.SrcRangeBegin), 10) + "-" + strconv.FormatUint(uint64(partInfo.SrcRangeEnd), 10)
		req.AddXOSS("x-oss-copy-source-range", cntRange)
	}

	rsp, err := req.Send()
	if err != nil {
		if _, ok := err.(*ossapi.Error); !ok {
			ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
			ossapiError = ossapi.OSSAPIError
			return
		}
	}
	if rsp.Result != ossapi.ErrSUCC {
		ossapiError = err.(*ossapi.Error)
		return
	}
	bodyLen, err := strconv.Atoi(rsp.HTTPRsp.Header["Content-Length"][0])
	if err != nil {
		ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
		ossapiError = ossapi.OSSAPIError
		return
	}
	body := make([]byte, bodyLen)
	rsp.HTTPRsp.Body.Read(body)
	rstInfo = new(UploadPartCopyRstInfo)
	xml.Unmarshal(body, rstInfo)
	return
}
