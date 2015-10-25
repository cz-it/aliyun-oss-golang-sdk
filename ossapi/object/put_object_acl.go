/**
* Author: CZ cz.theng@gmail.com
 */

package object

import (
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
)

func SetObjectACL(objName, bucketName, permission, location string) (error *ossapi.Error) {
	resource := path.Join("/", bucketName, objName)
	host := bucketName + "." + location + ".aliyuncs.com"
	req := &ossapi.Request{
		Host:     host,
		Path:     "/" + objName + "?acl",
		Method:   "PUT",
		Resource: resource,
		SubRes:   []string{"acl"}}
	req.AddXOSS("x-oss-object-acl", permission)

	rsp, err := req.Send()
	if err != nil {
		if _, ok := err.(*ossapi.Error); !ok {
			ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
			error = ossapi.OSSAPIError
			return
		}
	}
	if rsp.Result != ossapi.ESUCC {
		error = err.(*ossapi.Error)
		return
	}
	return
}