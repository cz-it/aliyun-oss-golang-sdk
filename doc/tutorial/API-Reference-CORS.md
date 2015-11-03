#API Reference: Service


##Query 

	./osscmd cors -q -b testossscmd  -a shenzhen
	Rule [0]:
	    AllowedOrigin: [* /]
	    AllowedMethod: [GET POST PUT DELETE HEAD]
	    AllowedHeader: [access-control-requet-headers authorization access-control-request-method]
	    ExposeHeader: [x-oss-test]
	    MaxAgeSeconds: 1000
##Delete

	./osscmd cors -d -b testossscmd  -a shenzhen
	Delete CORS Success.

## Option

	./osscmd cors -o -b testossscmd  -a shenzhen --object a.c --origin "/" --method "GET"  --headers "authorization"
	Option Info: &{* GET, POST, PUT, DELETE, HEAD authorization x-oss-test 1000}