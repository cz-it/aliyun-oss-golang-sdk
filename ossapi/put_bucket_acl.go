/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"path"
)

func PutBucketACL(name, permission, location string) (error *Error) {
	resource := path.Join("/", name)
	host := name + "." + location + ".aliyuncs.com"
	req := &Request{Host: host, Path: "/?acl", Method: "PUT", Resource: resource + "/", SubRes: []string{"acl"}}
	req.AddXOSS("x-oss-acl", permission)

	rsp, err := req.Send()
	if err != nil {
		if _, ok := err.(*Error); !ok {
			Logger.Error("GetService's Send Error:%s", err.Error())
			error = OSSAPIError
			return
		}
	}
	if rsp.Result != ESUCC {
		error = err.(*Error)
		return
	}
	b := make([]byte, 1024)
	rsp.httpRsp.Body.Read(b)
	return
}
