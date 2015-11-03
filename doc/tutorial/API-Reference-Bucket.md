#API Reference: Bucket

## bucket.Create
Create a Bucket with name , location and permission.

	Create(name, location, permission string) (ossapiError *ossapi.Error)
* name : bucket's name
* location: bucket's location , it's a enum value on ossapi:
	
		const (
		   L_Hangzhou      = "oss-cn-hangzhou"
		   L_Shenzhen      = "oss-cn-shenzhen"
		   L_Beijing       = "oss-cn-beijing"
		   L_Qingdao       = "oss-cn-qingdao"
		   L_Shanghai      = "oss-cn-shanghai"
		   L_HongKong      = "oss-cn-hongkong"
		   L_SiliconValley = "oss-us-west-1"
		   L_Singapore     = "oss-ap-southeast-1"
		  )
	location indicate where is bucket. The name of location is enum literal. 
* permission: bucke's permission ,  it's a enum value on ossapi:
		
		const (
			P_Private        = "private"
    		P_PublicReadOnly = "public-read"
    		P_PublicRW       = "public-read-write"
		)
	permission has three enum value,P_Private can be access only by owner.  P_PublicReadOnly can be read  by everyone, and write only by owner.  P_PublicRW can be write and read by anyone .

return nil error when success. If failed return a ossapi.Error.
## bucket.SetACL

	SetACL(name, location, permission string) (error *ossapi.Error)
Set Bucket (name+location) with perission.

name and location point a unique bucket, set it with permission. permission is a ossapi's enum value with descripted above.

return nil error when success. If failed return a ossapi.Error.
## bucket.OpenLogging

	OpenLogging(name, location, targetBucket, targetPrefix string) (ossapiError *ossapi.Error)
Open bucket's logging .It will log to bucket with name targetBucket wiht a logging object , logging's object wiht targetPrefix prefix name.such as "MyLog-oss-example-2012-09-10-04-00-00-0000" . Format is "<TargetPrefix><SourceBucket>-YYYY-mm-DD-HH-MM-SS-UniqueString". Here, "Mylog" is the prefix by targetPrefix.

return nil error when success. If failed return a ossapi.Error.

## bucket.CloseLogging

	CloseLogging(name, location string) (ossapiError *ossapi.Error)
	
Close Bucket's Logging ,Bucket is indicated by name and location.

return nil error when success. If failed return a ossapi.Error.

##bucket.SetWebsite

	SetWebsite(name, location, indexPage, errorPage string) (ossapiError *ossapi.Error)
Set Bucket's Website pages. it will set index page to indexPage and 404 page to errorPage.

return nil error when success. If failed return a ossapi.Error.

##bucket.SetLifecycle

	SetLifecycle(name, location string, rules []RuleInfo) (ossapiError *ossapi.Error)
Set Bucket's Lifecycle limitation. rules is a rule list:

	const (
	    LifecycleStatsEnable  = "Enabled"
	    LifecycleStatsDisable = "Disabled"
	)
	
	type ExpirationDaysInfo struct {
	    Days uint
	}
	
	type RuleInfo struct {
	    ID         string
	    Prefix     string
	    Status     string
	    Expiration ExpirationDaysInfo
	}
	
ID is a rule's ID. you can custom it.ExpirationDaysInfo points out how many days to expire.Prefix figure out object's name with this prefix be setted. status is "Enabled" and  "Disabled" point weather this rule has function.
##bucket.Delete

	Delete(name, location string) (ossapiError *ossapi.Error)
Delte bucket which named name and on location. If bucket has objects ,this action may be failed .

return nil error when success. If failed return a ossapi.Error.

## bucket.DeleteLifecycle

	DeleteLifecycle(name, location string) (ossapiError *ossapi.Error)
Delte all lifecycle rules on bucket which named name .After this, no object will be delte automaticly.

return nil error when success. If failed return a ossapi.Error.

## bucket.DeleteLogging

	DeleteLogging(name, location string) (ossapiError *ossapi.Error)
Close bucket's logging function. it is the same as bucket.CloseLogging

return nil error when success. If failed return a ossapi.Error.

##bucket.DeleteWebsite

	DeleteWebsite(name, location string) (ossapiError *ossapi.Error)
Delete bucket's website info. It is the same as bucket.SetWebsite with index and error "";

return nil error when success. If failed return a ossapi.Error.

##bucket.Query
	QueryObjects(name, location string, prefix, marker, delimiter, encodingType string, maxKeys int) (info     *BucktsInfo, ossapiError *ossapi.Error)
	
List All Object in Bucket . if prefix/marker/delimiter and encodingType is not "" .It Only return object match these. At most maxKeys items will return. objects store in info:

	type ContentInfo struct {
	    Key          string
	    LastModified string
	    ETag         string
	    Type         string
	    Size         string
	    StorageClass string
	    Owner        service.Owner
	}
	
	type CommonInfo struct {
	    Prefix string
	}
	
	type BucktsInfo struct {
	    XMLName        xml.Name `xml:"ListBucketResult"`
	    Name           string   `xml:"Name"`
	    Prefix         string   `xml:"Prefix"`
	    Marker         string   `xml:"Marker"`
	    MaxKeys        int      `xml:"MaxKeys"`
	    EncodingType   string   `xml:"encoding-type"`
	    IsTruncated    bool     `xml:"IsTruncated"`
	    Contents       []ContentInfo
	    CommonPrefixes CommonInfo `xml:"CommonPrefixes"`
	}
Real Object's meta info is in ContentInfo Such as ETag/ Size/ Type etc.
##bucket.QueryACL

	QueryACL(name, location string) (info *ACLInfo, ossapiError *ossapi.Error)
Query Bucket's ACL information. The result is stored in info.

	type OwnerInfo struct {
    ID          string
    DisplayName string
	}
	
	type AccessControlListInfo struct {
	    Grant string
	}
	
	type ACLInfo struct {
	    XMLName           xml.Name `xml:"AccessControlPolicy"`
	    Owner             OwnerInfo
	    AccessControlList AccessControlListInfo
	}
Owner store Owner Info, ACL is in AccessControlListInfo.Grant

return nil error when success. If failed return a ossapi.Error.

##bucket.QueryLocation

	QueryLocation(name string) (location string, ossapiError *ossapi.Error)
	
Query Where bucket is located. The return value is a string with "oss-cn" prefix and city name . Such as 
"oss-cn-shenzhen"	 means Shenzhen

return nil error when success. If failed return a ossapi.Error.

##bucket.QueryLogging

	QueryLogging(name, location string) (info *LoggingInfo, ossapiError *ossapi.Error)	
Query bucket's  Logging configure .If bucket has configure of logging , It is stored in info:

	type LoggingInfo struct {
	    TargetBucket string
	    TargetPrefix string
	}
	
If Not , info is nil 

return nil error when success. If failed return a ossapi.Error.

##bucket.QueryReferer

	QueryReferer(name, location string) (info *RefererConfigurationInfo, ossapiError *ossapi.Error)

Query bucket's referer white list . Urls stored in info.

	type RefererListInfo struct {
	    Referer []string
	}
	
	type RefererConfigurationInfo struct {
	    XMLName           xml.Name        `xml:"RefererConfiguration"`
	    AllowEmptyReferer bool            `xml:"AllowEmptyReferer"`
	    RefererList       RefererListInfo `xml:"RefererList"`
	}
	
If bucket allows empty access, AllowEmptyReferer will be true otherwise false. White Url List is on RefererList.Referer A string split.

##bucket.QueryLifcycle

	QueryLifecycle(name, location string) (infos []RuleInfo, ossapiError *ossapi.Error)
Query Bucket's lifecycle info .If buckt doesn't have a lifecycle, infos is nil and ossapiError is ENoSuchLifecycle.Otherwise  rules are in infos.

	type RuleInfo struct {
	    ID         string
	    Prefix     string
	    Status     string
	    Expiration ExpirationDaysInfo
	}
Expiration has the number of days to expire.

##bucket.QueryWebsite

	QueryWebsite(name, location string) (info *WebsiteInfo, ossapiError *ossapi.Error)
	
Query bucket's website info. The Index page and error page is stored in info.

	type IndexInfo struct {
	    Suffix string
	}
	type ErrorInfo struct {
	    Key string
	}
	type WebsiteInfo struct {
	    XMLName       xml.Name  `xml:"WebsiteConfiguration"`
	    IndexDocument IndexInfo `xml:"IndexDocument"`
	    ErrorDocument KeyInfo   `xml:"ErrorDocument"`
	}
	
Index Page is IndexDocument.Suffix and 404 Error Page is ErrorDocument.Key