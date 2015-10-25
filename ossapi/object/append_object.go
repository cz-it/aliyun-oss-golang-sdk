/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
	"strconv"
)

/*
//redefine on put_object
type ObjectInfo struct {
	CacheControl       string
	ContentDisposition string
	ContentEncoding    string
	Expires            string
	Encryption         string
	ACL                string
	ObjName            string
	BucketName         string
	Location           string
	Body               []byte
	Type               string
}
*/

type AppendObjInfo struct {
	ObjectInfo
	Position uint64
}

type AppendObjRspInfo struct {
	Possition uint64
	crc64     uint64
}

func AppendObject(objInfo *AppendObjInfo) (ossapiError *ossapi.Error) {
	if objInfo == nil {
		ossapiError = ossapi.ArgError
		return
	}
	resource := path.Join("/", objInfo.BucketName, objInfo.ObjName)
	host := objInfo.BucketName + "." + objInfo.Location + ".aliyuncs.com"
	header := make(map[string]string)
	if objInfo != nil {
		header["Cache-Control"] = objInfo.CacheControl
		header["Content-Disposition"] = objInfo.ContentDisposition
		header["Content-Encoding"] = objInfo.ContentEncoding
		header["Expires"] = objInfo.Expires
		header["x-oss-server-side-encryption"] = objInfo.Encryption
		header["x-oss-object-acl"] = objInfo.ACL
	}
	req := &ossapi.Request{
		Host:      host,
		Path:      "/" + objInfo.ObjName + "?append&position=" + strconv.FormatUint(objInfo.Position, 10),
		Method:    "POST",
		Resource:  resource,
		Body:      objInfo.Body,
		CntType:   objInfo.Type,
		SubRes:    []string{"append&position=" + strconv.FormatUint(objInfo.Position, 10)},
		ExtHeader: header}
	req.AddXOSS("x-oss-object-acl", objInfo.ACL)
	req.AddXOSS("x-oss-server-side-encryption", objInfo.Encryption)

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
