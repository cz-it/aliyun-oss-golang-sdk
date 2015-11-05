#API Reference: Service
Object has Create/Delete/Modify And Qeury API . 

##object.Create

	Create(objInfo *ObjectInfo) (ossapiError *ossapi.Error)
	
To create an object, you should point where it is (BucketName and Location), what's it's name(ObjName) and 
object's content.

	type ObjectInfo struct {
	    CacheControl       string
	    ContentDisposition string
	    ContentEncoding    string
	    Expires            string
	    Encryption         string
	    ACL                string
	    Body               []byte
	    Type               string
	}
Besides you can also set CacheControl to open cache , ContentEncoding such as "utf-8" , Expires to delete automatic , Encryption such as "AES256" .

ACL is the same as bukcet's ACL enum define . Type is Http content type such as "text/html".

return nil error when success. If failed return a ossapi.Error.

##object.Append

	Append(objName, bucketName, location string, objInfo *AppendObjInfo) (rstInfo *AppendObjRspInfo, ossapiError *ossapi.Error)

Append data to a object objName on bucketName with location. AppendObjInfo inherit form ObjectInfo .SO Use it just like create .

It has a Position more. Indicate that where to append.

On Success .It return new possition and crc . 

return nil error when success. If failed return a ossapi.Error.
##object.Copy

	Copy(copyInfo *CopyInfo, copyConnInfo *CopyConditionInfo) (rstInfo *CopyResultInfo, ossapiError *ossapi.Error)
Copy is another method to create an object. CopyInfo shows the source and target.

	type CopyInfo struct {
	    ObjectName string
	    BucketName string
	    Location   string
	    Source     string
	    Directive  string
	    Encryption string
	    ACL        string
	} 
copy Source to ObjectName on BucketName with Location . Directive is COPY or REPLACE. REPLACE will replace the new object if existed.Encryption is "AES256" and ACL is just like bucket.

if CopyConditionInfo present . Copy will do only it matched.

	type CopyConditionInfo struct {
	    ETAG         string
	    Date         string
	    LastModify   string
	    LastUnModify string
	}
ETAG is if matched. Date is the create time after. if last modify time older than LastModify time it matched. or if last modify time early than LastUnModify matched.
On Success object's ETag will return on CopyResultInfo.ETag.

return nil error when success. If failed return a ossapi.Error.

##object.Delete

	Delete(objName, bucketName, location string) (ossapiError *ossapi.Error)
	
Delete the object named objName on bucketName bucket with location.

return nil error when success. If failed return a ossapi.Error.

##object.DeleteObjects

	DeleteObjects(bucketName, location string, info *DeleteObjInfo) (rstInfo *DeleteObjRstInfo, ossapiError *ossapi.Error)
If you want to delete sever objects at once. You can put Ojbect in DeleteObjInfo. type DeleteObjInfo  :

	struct {
	    XMLName xml.Name `xml:"Delete"`
	    Quiet   bool     `xml:"Quiet"`
	    Object  []KeyInfo
	}

Put objects' key in a KeyInfo list.

On Success. Deleted Object's key will on DeleteObjRstInfo.Deleted , a KeyInfo list.

return nil error when success. If failed return a ossapi.Error.

##object.QueryMeta
	
	QueryMeta(objName, bucketName, location string, info *BriefConnInfo) (briefInfo *ObjectBriefInfo, ossapiError *ossapi.Error)
	
Query object's head information. 

##object.SetACL

	SetACL(objName, bucketName, location, permission string) (error *ossapi.Error)	
Set Object's ACL to bucket.P_Private P_PublicReadOnly or P_PublicRW

##object.QueryACL

	QueryACL(objName, bucketName, location string) (info *bucket.ACLInfo, ossapiError *ossapi.Error)

Query Object's ACL info. On Success , bucket.ACLInfo hold the ACL info .bucket.ACLInfo see bucket.