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
	
-w option meas set website . set Bucket's Index to --index and 404 page to --error .
	
	