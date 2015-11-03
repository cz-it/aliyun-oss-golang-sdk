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