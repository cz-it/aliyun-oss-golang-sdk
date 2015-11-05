#How To Use OSS CMD Tools
OSS CMD tool is a command tool to help test OSSAPI. 
Use golang's get command can get it .

	go get github.com/cz-it/aliyun-oss-golang-sdk/osscmd
Use "osscmd -v" to test if install successfuly.

This tool has two layer option like git ,such as `git commit -m "commit"` .Use OSSCMD `osscmd init -i YOUR_ACCESS_KEY_ID` .Here init is the first layer option and -i is the second.

## 1.Init With Your ID And Secort
To use OSS CMD tool .First you should invoke init command .This will store your Access Key ID and Access Key Secort to ~/.osscmd/config.xml.Then you can use other commands without  Access Key ID and Access Key Secort.

	./osscmd init -i v8P430U3UcILPA -s EB9v8yL2aM07YOgtO1BdfrXtdxa1
	[SUCC]: Init Success!
Then cat config file like this:

	cat ~/.osscmd/config.xml
	<Config>
		<AccessKeyID>v8P430U3UcILPA</AccessKeyID>
		<AccessKeySecret>EB9v8yL2aM07YOgtO1BdfrXtdxa1</AccessKeySecret>
	</Config>	
After this , you have successfull init OSSCMD.
## 2.Buckets
`bucket` is the first option to do things with bucket.

### Query All Your Buckets

	./osscmd bucket -l
	Owner: 1415982622007927
		Bucket[0]:{test-cors 2015-10-27T17:38:40.000Z oss-cn-hangzhou}
		Bucket[1]:{test-mupload 2015-10-26T03:01:30.000Z oss-cn-hangzhou}
		Bucket[2]:{test-object-hz 2015-10-24T08:17:21.000Z oss-cn-hangzhou}
		
-l option will list all your buckets. Every one with a Name/Create Time And Location.
### Create Bucket
	./osscmd bucket -c -b testosscmd2  -a shenzhen
	Create Bucket testosscmd2 Success !
	
-c option will create a bucket named -b on location -a . Location now have shenzhen/beijin/shanghai/hongkong and qingdao which mean literal city's name.	
###Set Bucket's Attributes
osscmd has subcommand bucket and option -s to set bucket's attributes. 

Usage：
	
	./osscmd bucket -s -b bucket'name -a location [-p|]
	
Set Must Point out bucket's name and location	
#### Set ACL

	./osscmd bucket -s -b testossscmd -a shenzhen -p RO
	Set Bucket testossscmd:RO Success !
	
Use -b to indicate bucket's name , -a to indicate bucket's loaction, location's value is descripted on [Create Bucket](#Create Bucket). -p to indicat bucket's permission , permission's value is descripted on [Create Bucket](#Create Bucket).

#### Set Logging
To Open Logging. Use:

	./osscmd bucket -s -g -b testossscmd -a shenzhen --log-possition testossscmd --log-prefix osscmdlog
	Set Bucket testossscmdlog  Success !
	
To Close Logging. Use:

	./osscmd bucket -s -g -b testossscmd -a shenzhen
	Set Bucket testossscmdlog  Success !
	
-g option means set logging . if --log-possition present , logging will be open  and log to testossscmd bucket with object name's prefix osscmdlog.  if not ,it will close logging.

####Set WebSite
Set Bucket's WebSite pages

	./osscmd bucket -s -w -b testossscmd -a shenzhen --index index.html --error 404.html
	Set Bucket testossscmd website  Success !
	
-w option meas set website . set Bucket's Index to --index and 404 page to --error .

#### Set Referer
Set Bucket's Referer urls

	./osscmd bucket -s -r -b testossscmd -a shenzhen baidu.com qq.com
	Set Bucket testossscmd refer Success !
-r option means set referer , The white list followed . Here is "baidu.com ","qq.com"

### Delete Bucket's Attributes
osscmd has subcommand bucket and option -d to delete bucket's attributes. 

Usage：
	
	./osscmd bucket -d [--logging|website|lifecycle] -b bucket'name -a location
	
#### Dlete bucket

	./osscmd bucket -d   -b test-put-bucket2  -a hangzhou

	Delete Bucket test-put-bucket2 Success !	
Delete the bucket test-put-bucket2
	 
#### Delete Website Info

	./osscmd bucket -d --website  -b testossscmd -a shenzhen
	Delete Bucket testossscmd  Website Success !
	
--webiste options  indicate to delete webiste info of bucket testosscmd

#### Delete Lifecycle 

	./osscmd bucket -d --lifecycle  -b testossscmd -a shenzhen
	Delete Bucket testossscmd  LifeCycle Success !
--lifecycle options indicate to delete lifycycle info of bucket testosscmd

#### Delete/Close Logging

	./osscmd bucket -d --logging -b testossscmd -a shenzhen
	Delete Bucket testossscmd  Logging Success !
	
--logging options indicate to close logging function of bucket testosscmd

### Query Bucket's Attributes
osscmd has subcommand bucket and option -q to get bucket's attributes. 

Usage：
	
	./osscmd bucket -q [--logging|website|lifecycle] -b bucket'name [-a location]

#### Get Bucket's ACL
	
	./osscmd bucket -q --acl -b testossscmd -a shenzhen
	Owner is : {1415982622007927 1415982622007927}
	Grant is : public-read
--acl option indicate to get bucket's acl infomation. Grant show the ACL: public-read.
#### Query objects in Bucket
	 ./osscmd bucket -q   -b testossscmd -a shenzhen
	Objects [0] are: {a.c 2015-11-03T09:49:10.000Z "B0222DA9C0BC538F896ED4441F9F7C24" Normal 67 Standard {1415982622007927 1415982622007927}}
	
This will Query all objects in bucket testossscmd.And List followed.

#### Query where bucket is 

	./osscmd bucket -q --location -b testossscmd
	Location is  oss-cn-shenzhen
	
--location option indicate to query where bucket is . Here it is Shenzhen

####Query Logging info

	./osscmd bucket -q --logging -b testossscmd -a shenzhen
	Target Bucket testossscmd
	Target Prefix cmd

	./osscmd bucket -q --logging -b testossscmd -a shenzhen
	Bucket has not config logging
	
--logging option indicate to query bucket's logging info .
If bucket has config logging , it will return with Target Bucket name and Target Prefix name. If Not , return 
"Bucket has not config logging"

#### Query Referer Whilte list

	./osscmd bucket -q --referer  -b testossscmd -a shenzhen
	If Allow Empyt :true
	White List [http://baidu.com http://qq.com]
	
--referer option indicate to query bucket's referer info . If bucket allow emtpy url , it return true , otherwise false. The white list is followed.

#### Query Website info

	./osscmd bucket -q --website  -b testossscmd -a shenzhen
	Index is : index.html
	404page is : 404.html

--website option indicate to query bucket's website configurtion. index page and 404 error page is return .


#### Query Lifecycle

	 ./osscmd bucket -q --lifecycle  -b testossscmd -a shenzhen
	Rule [0]:{97d7bc25-5ed3-444d-83f4-27f43c750ff0 cmd Enabled {30}}
	
--lifecycle option indicate to query bucket's lifecycle info .If bucket has no lifecycle , it return a NoSuchLifecycle Error . If bucket has lifecycles , it show evey rules.
	
##3.Object

`objcet` is a first layer option which do things with object.
###Create Object

	./osscmd object -n -b testossscmd -a shenzhen --file ./main.go  --encoding utf-8 --expire "Fri, 28 Feb 2016 05:38:42 GMT" -p RO --type "text/html"
	Create Object./main.go Success !
-n means new .Create a Ojbect in bucket testossscmd. Object's content is in --file . --encoding points content's encoding type.and --type is the HTTP file type. --expire is expire time in format GMT such as "Fri, 28 Feb 2016 05:38:42 GMT" 

###Query Object's data

	./osscmd object -q -b testossscmd -a shenzhen --object a.c
	Data: #include <stdio.h>
-q means query object's data . object --object is on bucket -b .

int main(int argc, char *argv[])
{
	return 0;
}

###Copy Object
Copy a object to create a new one

	./osscmd object -c -b testossscmd -a shenzhen --object copynew2.go --source "/testossscmd/main.go" --directive COPY
	Copy Object Success , New Object is  &{{ CopyObjectResult} "3BECD44293735912D10E269B6FFDF273" 2015-11-04T08:00:03.000Z}
	
-c means copy . Copy from --srouce to generage a --object object on -b bucket .--directive influncse when object is existed. Only COPY or REPLACE.

### Append Object

	 ./osscmd object -a -b testossscmd -a shenzhen --file ./cors.go  --encoding utf-8 --expire "Fri, 28 Feb 2016 05:38:42 GMT" -p RO --type "text/html" --position 11352 --obejct cors.go
	Append Success. resuult: &{15136 9223372036854775807}
	
Append file to object --object. It is just like create .But a --position more . position should be equeal to result info Possition.
	
### Delete Object

	./osscmd object -d --object copynew2.go -b testossscmd -a shenzhen
	Delete copynew2.go Success
	
Delete Object with name --objcet on -b bucket.

### Query Object meat Info
	./osscmd object -m --object copynew.go -b testossscmd -a shenzhen
	Meta copynew.go  is  &{Normal text/html Wed, 04 Nov 2015 07:55:07 GMT "3BECD44293735912D10E269B6FFDF273" 1107}
Query Ojbect's Meta Infomation. Such as 

* object type: normal or appendable
* content type : "text/htl"
* expire time
* ETag

### Set Object's ACL

	./osscmd object -s --object copynew.go -b testossscmd -a shenzhen -p RO
	Set Object's ACL Success
	
Set Ojbect permission to RO(Pubic ReadOnly) , RW(Public Write And Read) and PT(Private)

### Query Object's ACL

	./osscmd object -q --acl --object copynew.go -b testossscmd -a shenzhen
	Object's ACL Owner: {1415982622007927 1415982622007927}
	Object's ACL: {public-read}
	
Query object's ACL info . --acl option should be given. Object's Owner and permission info will return.
	
##4.MultiUpload
`mupload` is a first layer optin which do things with multiupload .

### Init a MulitplyUpload Context.

	./osscmd mupload -i --object init -b testossscmd -a shenzhen
	Init Mupload Success:
	Key is: init
	Id is:  A2D59D7FD7F24660A08449C65F53241D

Init a MulitplyUpload Context . ID And Key will be used to do the following actions.

### Append a slice to context

	./osscmd mupload -a --object mupload -b testossscmd -a shenzhen --number 1 --file cors.go
	Add Part Success, Tag is  "E304006340F83A79E3D6D2D36B2FFA69"
		
-a Add --file's content previous init .--number indicates the sequence.--number should be [1,1000], 0 is not valied.

### Complate Upload

	./osscmd mupload -c --object mupload -b testossscmd -a shenzhen
	Success Done Complete .
	
-c will finish upload .
	
## Query all uploading objects

	./osscmd  mupload -m  -b testossscmd -a shenzhen
	Uploading Slices:
	Part [0]{mupload 3AE43B3C0FE24F4D8E74C198C506B3BE 2015-11-04T17:37:23.000Z}
	
-m options list all uploading objects.

## Query all uploading parts

	./osscmd  mupload -l --object mupload  -b testossscmd -a shenzhen
	Uploaded Parts:
	Part [0]{1 2015-11-04T17:38:14.000Z "E304006340F83A79E3D6D2D36B2FFA69" 3784}
	Part [1]{2 2015-11-04T17:38:19.000Z "E304006340F83A79E3D6D2D36B2FFA69" 3784}
-l option List all uploading parts.
	
## Abort Upload

	./osscmd mupload -d --object mupload -b testossscmd -a shenzhen
	Success Cacnel Upload Task
	
Cancle will Abort Uploading , And objects will not be stored.

##5. CORS
CORS allow other domain to access resources 

`cors` is a first layer option which do things with CORS.
### Query Bucket's CORS

	./osscmd cors -q -b testossscmd  -a shenzhen
	Rule [0]:
	    AllowedOrigin: [* /]
	    AllowedMethod: [GET POST PUT DELETE HEAD]
	    AllowedHeader: [access-control-requet-headers authorization access-control-request-method]
	    ExposeHeader: [x-oss-test]
	    MaxAgeSeconds: 1000
-q option means query . Query bucket testossscmd's CORS Info.

Info is a Rule list. With AllowedOrigin  AllowedMethod AllowedHeader ExposeHeader MaxAgeSeconds

### Delete Bucket's CORS

	./osscmd cors -d -b testossscmd  -a shenzhen
	Delete CORS Success.
	
-d options means delete . Delete Bucket's CORS attribute.

### Query object's CORS

	./osscmd cors -o -b testossscmd  -a shenzhen --object a.c --origin "/" --method "GET"  --headers "authorization"
	Option Info: &{* GET, POST, PUT, DELETE, HEAD authorization x-oss-test 1000}
	
-o options means option, Query ojbect's bucket's CORS information

--object : object name
--origin : origin path
--method : http method
--headers: http headers
