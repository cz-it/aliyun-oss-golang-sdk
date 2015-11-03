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


	
	