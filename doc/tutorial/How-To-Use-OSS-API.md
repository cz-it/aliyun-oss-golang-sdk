#How To Use OSS API
To use OSS API . First use `go get ` to download sdk.

	go get github.com/cz-it/aliyun-oss-golang-sdk/ossapi
	
Then import sdk in your golang files.

	package main
	
	import (
	    "fmt"
	    "github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	) 
	
	func main() {
	    fmt.Println(ossapi.Version())
	}

OSSAPI has four main parts:

	github.com/cz-it/aliyun-oss-golang-sdk/ossapi/service
	github.com/cz-it/aliyun-oss-golang-sdk/ossapi/bucket
	github.com/cz-it/aliyun-oss-golang-sdk/ossapi/object
	github.com/cz-it/aliyun-oss-golang-sdk/ossapi/mupload
	github.com/cz-it/aliyun-oss-golang-sdk/ossapi/cors

	
* [service](https://github.com/cz-it/aliyun-oss-golang-sdk/blob/master/doc/tutorial/API-Reference-Service.md)Describe  how to get all buckets
* [bucket](https://github.com/cz-it/aliyun-oss-golang-sdk/blob/master/doc/tutorial/API-Reference-Bucket.md) Describe how to deal with bucket.Such as create, query, delete etc.
* [object] (https://github.com/cz-it/aliyun-oss-golang-sdk/blob/master/doc/tutorial/API-Reference-Object.md) Describe how to deal with obejct. Such as create, append, query etc.
* [mupload](https://github.com/cz-it/aliyun-oss-golang-sdk/blob/master/doc/tutorial/API-Reference-MultipartUpload.md) Describe how to upload a big file .
* [cors](https://github.com/cz-it/aliyun-oss-golang-sdk/blob/master/doc/tutorial/API-Reference-CORS.md) Describe how to access source to remote domain.


Before All This. You Should Invoke ossapi.Init to init SDK.

	 Init(ID string, secret string) error
Init SDK with your AccessKeyID and AccessKeySecort. then you can invoke from above.

OSSAPI return all OSS error in  ossapi.Error object.

	type Error struct {
	    XMLName      xml.Name `xml:"Error"`
	    ErrNo        string   `xml:"Code"`
	    ErrMsg       string   `xml:"Message"`
	    HttpStatus   int
	    ErrDetailMsg string
	}
[ErrNo](https://github.com/cz-it/aliyun-oss-golang-sdk/blob/master/doc/tutorial/API-Reference-Errno.md) Descript Error name . 

ErrMsg show Aliyun's OSS Error in brief. You cant use it to contact with aliyun's help.

ErrDetailMsg show Aliyun's OSS Deatil msg which formt in xml . 


Here is an example to show how to create a bucket and a object on it .More exapmle can find on [osscmd](github.com/cz-it/aliyun-oss-golang-sdk/cmd) And API Reference above.
		package main
		
		import (
			"fmt"
			"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
			"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/bucket"
			"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/object"
			"os"
		)
		
		func Exit(err *ossapi.Error) {
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(-1)
			}
		}
		
		func main() {
			fmt.Println(ossapi.Version())
			err := ossapi.Init("v8P430U3UcILPA", "EB9v8yL2aM07YOgtO1BdfrXtdxa1")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			bucketName := "testossapi"
			e := bucket.Create(bucketName, bucket.L_Shenzhen, bucket.P_PublicRW)
			if e != nil {
				Exit(e)
			}
			fmt.Println("Create Bucket ", bucketName)
		
			objName := "testossapi"
			objInfo := &object.ObjectInfo{
				CacheControl:       "no-cache",
				ContentDisposition: "attachment;",
				ContentEncoding:    "utf-8",
				Expires:            " Fri, 28 Nov 2015 05:38:42 GMT",
				Encryption:         "AES256",
				ACL:                bucket.P_PublicRW,
				Body:               []byte("<html></html>"),
				Type:               "text/html"}
			e = object.Create(objName, bucketName, bucket.L_Shenzhen, objInfo)
			if e != nil {
				Exit(e)
			}
			fmt.Println("Create Object", objName, " on ", bucketName, " Success.")
		
			data, e := object.Query(objName, bucketName, bucket.L_Shenzhen, nil, nil)
			if e != nil {
				Exit(e)
			}
			fmt.Println("Objetct ", objName, " Data:", string(data))
		}


run :

	go run main.go
	1.0.0
	Create Bucket  testossapi
	Create Object testossapi  on  testossapi  Success.
	Objetct  testossapi  Data: <html></html>