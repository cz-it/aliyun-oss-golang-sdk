/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"testing"
)

func TestAppendObject(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}

	objInfo := &AppendObjInfo{ObjectInfo: ObjectInfo{
		CacheControl:       "no-cache",
		ContentDisposition: "attachment;filename=oss_download.jpg",
		ContentEncoding:    "utf-8",
		Expires:            "Fri, 28 Feb 2012 05:38:42 GMT",
		Encryption:         "AES256",
		ACL:                ossapi.P_Private,
		ObjName:            "append2",
		Location:           ossapi.L_Hangzhou,
		Body:               []byte("<html><head></head><body>test</body></html>"),
		Type:               "text/html",
		BucketName:         "test-object-hz"},

		Position: 43}
	if err := AppendObject(objInfo); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("AppendObject Success!")
	}

}
