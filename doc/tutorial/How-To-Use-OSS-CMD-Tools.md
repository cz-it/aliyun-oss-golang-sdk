#How To Use OSS CMD Tools


## List Buckets

## Create Bucket

## Set Bucket's Attributes
osscmd has subcommand bucket and option -s to set bucket's attributes. 

Usage：
	
	./osscmd bucket -s -b bucket'name -a location [-p|]
	
Set Must Point out bucket's name and location	
### Set ACL

	./osscmd bucket -s -b testossscmd -a shenzhen -p RO
	Set Bucket testossscmd:RO Success !
	
Use -b to indicate bucket's name , -a to indicate bucket's loaction, location's value is descripted on [Create Bucket](#Create Bucket). -p to indicat bucket's permission , permission's value is descripted on [Create Bucket](#Create Bucket).

### Set Logging
To Open Logging. Use:

	./osscmd bucket -s -g -b testossscmd -a shenzhen --log-possition testossscmd --log-prefix osscmdlog
	Set Bucket testossscmdlog  Success !
	
To Close Logging. Use:

	./osscmd bucket -s -g -b testossscmd -a shenzhen
	Set Bucket testossscmdlog  Success !
	
-g option means set logging . if --log-possition present , logging will be open  and log to testossscmd bucket with object name's prefix osscmdlog.  if not ,it will close logging.

###Set WebSite
Set Bucket's WebSite pages

	./osscmd bucket -s -w -b testossscmd -a shenzhen --index index.html --error 404.html
	Set Bucket testossscmd website  Success !
	
-w option meas set website . set Bucket's Index to --index and 404 page to --error .

### Set Referer
Set Bucket's Referer urls

	./osscmd bucket -s -r -b testossscmd -a shenzhen baidu.com qq.com
	Set Bucket testossscmd refer Success !
-r option means set referer , The white list followed . Here is "baidu.com ","qq.com"

## Delete Bucket's Attributes
osscmd has subcommand bucket and option -d to delete bucket's attributes. 

Usage：
	
	./osscmd bucket -d [--logging|website|lifecycle] -b bucket'name -a location
	
### Dlete bucket

	./osscmd bucket -d   -b test-put-bucket2  -a hangzhou

	Delete Bucket test-put-bucket2 Success !	
Delete the bucket test-put-bucket2
	 
### Delete Website Info

	./osscmd bucket -d --website  -b testossscmd -a shenzhen
	Delete Bucket testossscmd  Website Success !
	
--webiste options  indicate to delete webiste info of bucket testosscmd

### Delete Lifecycle 

	./osscmd bucket -d --lifecycle  -b testossscmd -a shenzhen
	Delete Bucket testossscmd  LifeCycle Success !
--lifecycle options indicate to delete lifycycle info of bucket testosscmd

### Delete/Close Logging

	./osscmd bucket -d --logging -b testossscmd -a shenzhen
	Delete Bucket testossscmd  Logging Success !
	
--logging options indicate to close logging function of bucket testosscmd

## Query Bucket's Attributes
osscmd has subcommand bucket and option -q to get bucket's attributes. 

Usage：
	
	./osscmd bucket -q [--logging|website|lifecycle] -b bucket'name [-a location]

### Get Bucket's ACL
	
	./osscmd bucket -q --acl -b testossscmd -a shenzhen
	Owner is : {1415982622007927 1415982622007927}
	Grant is : public-read
--acl option indicate to get bucket's acl infomation. Grant show the ACL: public-read.
### Query objects in Bucket
	 ./osscmd bucket -q   -b testossscmd -a shenzhen
	Objects [0] are: {a.c 2015-11-03T09:49:10.000Z "B0222DA9C0BC538F896ED4441F9F7C24" Normal 67 Standard {1415982622007927 1415982622007927}}
	
This will Query all objects in bucket testossscmd.And List followed.

### Query where bucket is 

	./osscmd bucket -q --location -b testossscmd
	Location is  oss-cn-shenzhen
	
--location option indicate to query where bucket is . Here it is Shenzhen

###Query Logging info

	./osscmd bucket -q --logging -b testossscmd -a shenzhen
	Target Bucket testossscmd
	Target Prefix cmd

	./osscmd bucket -q --logging -b testossscmd -a shenzhen
	Bucket has not config logging
	
--logging option indicate to query bucket's logging info .
If bucket has config logging , it will return with Target Bucket name and Target Prefix name. If Not , return 
"Bucket has not config logging"

### Query Referer Whilte list

	./osscmd bucket -q --referer  -b testossscmd -a shenzhen
	If Allow Empyt :true
	White List [http://baidu.com http://qq.com]
	
--referer option indicate to query bucket's referer info . If bucket allow emtpy url , it return true , otherwise false. The white list is followed.

### Query Website info

	./osscmd bucket -q --website  -b testossscmd -a shenzhen
	Index is : index.html
	404page is : 404.html

--website option indicate to query bucket's website configurtion. index page and 404 error page is return .


### Query Lifecycle

	 ./osscmd bucket -q --lifecycle  -b testossscmd -a shenzhen
	Rule [0]:{97d7bc25-5ed3-444d-83f4-27f43c750ff0 cmd Enabled {30}}
	
--lifecycle option indicate to query bucket's lifecycle info .If bucket has no lifecycle , it return a NoSuchLifecycle Error . If bucket has lifecycles , it show evey rules.
	