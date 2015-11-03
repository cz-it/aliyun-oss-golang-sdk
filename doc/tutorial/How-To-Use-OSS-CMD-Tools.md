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