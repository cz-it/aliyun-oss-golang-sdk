#How To Get Aliyun OSS Golang SDK Working
Aliyun OSS Golang SDK is a warp for Aliyun OSS [Web API](https://docs.aliyun.com/?spm=5176.1980653.30105.5.rDaFkB#/pub/oss/api-reference/abstract) on Golang. To use this SDK you should install golang first, see how to [install golang](https://golang.org/doc/install) .The following articles assume that golang has been installed.

Use `go get` commands to get Aliyun OSS Golang SDK and it's command tools "osscmd". 

	go get github.com/cz-it/aliyun-oss-golang-sdk/ossapi
	go get github.com/cz-it/aliyun-oss-golang-sdk/osscmd
	
After installed SDK and command tool ,you can use command 

	osscmd -v 
	
To test if installed success.

Or you can `import github.com/cz-it/aliyun-oss-golang-sdk/ossapi` and invoke the `Version()` fucntion .

	package main

	import (
		"fmt"
		"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	) 
	
	func main() {
		fmt.Println(ossapi.Version())
	}
	
After testing installed successful.You can use API like demo this.


	package main
	
	import (
		"fmt"
		"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
		"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/service"
	)
	
	func main () {
		err := ossapi.Init("v8P430U3UcILP", "EB9v8yL2aM07YOgtO1BdfrXtdxa4")
		if err != nil {
			fmt.Println(err.Error())
			return 
		}
		
		info, e := service.GetServiceDefault()
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		fmt.Println("Owner:", info.Owner.DisplayName)
		for idx, bucket := range info.Buckets.Bucket {
			fmt.Printf("	Bucket[%v]:", idx)
			fmt.Println(bucket)
		}
	}
	
Use `go run ` may get output like this:

	Owner: 1415982622007927
		Bucket[0]:{python-sdk-test 2015-10-21T12:31:32.000Z oss-cn-hangzhou}
		Bucket[1]:{python-sdk-test2 2015-10-21T12:31:50.000Z oss-cn-hangzhou}
		Bucket[2]:{test-cors 2015-10-27T17:38:40.000Z oss-cn-hangzhou}
		
Here use "Init" with your Access Key ID and Access Secret to init ossapi. Then use "GetServiceDefault" to list all your buckets.

*What Next?* 

* See [How-To-Use-OSS-CMD-Tools](https://github.com/cz-it/aliyun-oss-golang-sdk/blob/master/doc/tutorial/How-To-Use-OSS-CMD-Tools.md) for SDK's command tool "osscmd".
* See [How-To-Use-OSS-API](https://github.com/cz-it/aliyun-oss-golang-sdk/blob/master/doc/tutorial/How-To-Use-OSS-API.md) for using SDK's API.