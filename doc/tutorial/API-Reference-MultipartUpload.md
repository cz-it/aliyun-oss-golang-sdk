#API Reference: MultipartUpload

MultipartUpload allow you to upload big files .Such 1G file. You can split 1G to 100 10M file pices to upload in one uploading context.


##muload.Init

	Init(objName, bucketName, location string, initInfo *InitInfo) (rstInfo *InitRstInfo, ossapiError *ossapi.Error)
	
At first you should Init a uploading context. If complete upload. All data will store in object objName on bucket bucketName with location .  initInfo can set context attributes.

	type InitInfo struct {
	    CacheControl       string
	    ContentDisposition string
	    ContentEncoding    string
	    Expires            string
	    Encryption         string
	}
such as if have cache , content's encoding.

On Success. A important ID is in InitRstInfo.UploadId. This ID represent a context. 

return nil error when success. If failed return a ossapi.Error.
##mupload.Append

	Append(partInfo *UploadPartInfo) (rstInfo *UploadPartRstInfo, ossapiError *ossapi.Error)
	
After you have a uploading context. You can append a data pice to the context. Here you should give a data pice's number ID and a uploanding context ID to indicated upload to which one.

	type UploadPartInfo struct {
	    ObjectName string
	    BucketName string
	    Location   string
	    PartNumber int
	    UploadID   string
	    Data       []byte
	    CntType    string
	}
	
pice's data is on UploadPartInfo.Data .Also you should show where to store this pice with ObjectName, BucketName and Location.

On Success. A pice ETag will return on UploadPartRstInfo.Etag . Becarefull , you should store it  with the nubmer ID for complete uploading .

return nil error when success. If failed return a ossapi.Error.

## mupload.Complete

	Complete(objName, bucketName, location string, uploadId string, info *PartsInfo) (rstInfo *PartsCompleteInfo, ossapiErro    r *ossapi.Error)
After append all your data pices.You should invoke complete to combine a object to objName on bucketName with location. uploadId indicate which uploading context to complete. 

Besides, The ETag-NubmerID above should give to PartsInfo to descript who are the final valied data pices.

	type PartInfo struct {
	    PartNumber int
	    ETag       string
	}
On success. Object's Etag will return .

	type PartsCompleteInfo struct {
	    XMLName  xml.Name `xml:"CompleteMultipartUploadResult"`
	    Location string   `xml:"Location"`
	    Bucket   string   `xml:"Bucket"`
	    Key      string   `xml:"Key"`
	    ETag     string   `xml:"ETag"`
	}
	
Here Key may be the object's name on Bucket with Location.

return nil error when success. If failed return a ossapi.Error.
## mupload.Abort
	Abort(objName, bucketName, location, uploadID string) (ossapiError *ossapi.Error)
Abort a uploading context.

All uploaded pices will be delete. And the context ID will be invalied.

return nil error when success. If failed return a ossapi.Error.

## mupload.QueryParts
	QueryParts(objName, bucketName, location string, uploadID string, filter *PartsFilterInfo) (rstInfo *PartsResultInfo, os    sapiError *ossapi.Error)
Query Parts will query all uploaded and uploading data pices for uploading context uploadID. PartsFilterInfo can give the query condition.

	type PartsFilterInfo struct {
	    MaxParts         int
	    PartNumberMarker int
	    Encoding         string
	}
MaxParts at most MaxParts pices return . PartNumberMarker indicated Onle return NubmerID bigger than .

All parts meta information on PartsFilterInfo

	type PartsResultInfo struct {
	    XMLName              xml.Name `xml:"ListPartsResult"`
	    Bucket               string   `xml:"Bucket"`
	    Key                  string   `xml:"Key"`
	    UploadId             string   `xml:"UploadId"`
	    NextPartNumberMarker string   `xml:"NextPartNumberMarker"`
	    MaxParts             int      `xml:"MaxParts"`
	    IsTruncated          bool     `xml:"IsTruncated"`
	    Part                 []PartListInfo
	}
	
Such as UploadID ,MaxParts etc. Part is a list of PartListInfo:

	type PartListInfo struct {
	    PartNumber   int
	    LastModified string
	    ETag         string
	    Size         uint64
	}
With part NumberID, ETag and Size.


return nil error when success. If failed return a ossapi.Error.

## mupload.QueryObject
	QueryObjects(bucketName, location string, filter *FilterInfo) (rstInfo *MultipartUploadsResultInfo, ossapiError *ossapi.    Error)
	
Query all uploading context.FilterInfo give the query condition.

	type FilterInfo struct {
	    Delimiter      string
	    MaxUploads     int
	    KeyMarker      string
	    Prefix         string
	    UploadIDMarker string
	    Encoding       string
	}
Delimiter is COPY or REPLACE. KeyMarker is a dictonary sort of object's name or pices' key.Prefix is key's or object's prefix.

On  Success uploading context mete infomation on MultipartUploadsResultInfo 

	type MultipartUploadsResultInfo struct {
	    XMLName            xml.Name `xml:"ListMultipartUploadsResult"`
	    Bucket             string   `xml:"Bucket"`
	    KeyMarker          string   `xml:"KeyMarker"`
	    UploadIdMarker     string   `xml:"UploadIdMarker"`
	    NextKeyMarker      string   `xml:"NextKeyMarker"`
	    NextUploadIdMarker string   `xml:"NextUploadIdMarker"`
	    Delimiter          string   `xml:"Delimiter"`
	    Prefix             string   `xml:"Prefix"`
	    MaxUploads         int      `xml:"MaxUploads"`
	    IsTruncated        bool     `xml:"IsTruncated"`
	    Upload             []UploadInfo
	}

Such as Delimiter COPY or REPLACE. Context is on Upload a list of UploadInfo:

	type UploadInfo struct {
	    Key       string
	    UploadId  string
	    Initiated string
	}
	
With weather initialed ,or context UploadId and Key.

return nil error when success. If failed return a ossapi.Error.

##mupload.Copy

	Copy(partInfo *UploadPartCopyInfo, copyConnInfo *object.CopyConditionInfo) (rstInfo *UploadPartCopyRstInfo, ossapiError *ossapi.Error)
	
Besides append  a data slice. You can also copy a existed data slice to a uploading context. source is:

	type UploadPartCopyInfo struct {
	    ObjectName string
	    BucketName string
	    Location   string
	    PartNumber int
	    UploadID   string
	    SrcObjectName string
	    SrcBucketName  string
	    SrcRangeBegin int
	    SrcRangeEnd   int
	}
SrcObjectName and SrcBucketName points to the source object.SrcRangeBegin and SrcRangeEnd calculate data slice. copy these data to uploading context UploadID with NumberID PartNumber . You can also give object.CopyConditionInfo to decided wheate to copy .

On Success. A Etag will return on UploadPartCopyRstInfo.ETag
	
return nil error when success. If failed return a ossapi.Error.