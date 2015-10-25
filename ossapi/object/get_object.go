/**
* Author: CZ cz.theng@gmail.com
 */

package object

import ()

type RspObjInfo struct {
	CntType      string
	LastModified string
	ETag         string
	Ranges       string
	Type         string
	Length       int
	Data         []byte
}

/*
func GetObject(objInfo *ObjectInfo) (ossapiError *ossapi.Error) {
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
		Path:      "/" + objInfo.ObjName,
		Method:    "PUT",
		Resource:  resource,
		Body:      objInfo.Body,
		CntType:   objInfo.Type,
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
*/
