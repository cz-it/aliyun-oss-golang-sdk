/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"encoding/xml"
	"errors"
)

var (
	//Inner Errors
	ESUCC    = errors.New("Success!")
	EFAIL    = errors.New("HTTP Request Failed(4xx/5xx)!")
	EUNKNOWN = errors.New("HTTP Request With Unknown Status (NOT 2xx/4xx/5xx)!")
	EARG     = errors.New("Invalied Argument!")
)

const (
	ENone                                = "None"
	EAccessDenied                        = "AccessDenied"
	EBucketAlreadyExists                 = "BucketAlreadyExists"
	EBucketNotEmpty                      = "BucketNotEmpty"
	EEntityTooLarge                      = "EntityTooLarge"
	EEntityTooSmall                      = "EntityTooSmall"
	EFileGroupTooLarge                   = "FileGroupTooLarge"
	EInvalidLinkName                     = "InvalidLinkName"
	ELinkPartNotExist                    = "LinkPartNotExist"
	EObjectLinkTooLarge                  = "ObjectLinkTooLarge"
	EFieldItemTooLong                    = "FieldItemTooLong"
	EFilePartInterity                    = "FilePartInterity"
	EFilePartNotExist                    = "FilePartNotExist"
	EFilePartStale                       = "FilePartStale"
	EIncorrectNumberOfFilesInPOSTRequest = "IncorrectNumberOfFilesInPOSTRequest"
	EInvalidArgument                     = "InvalidArgument"
	EInvalidAccessKeyId                  = "InvalidAccessKeyId"
	EInvalidBucketName                   = "InvalidBucketName"
	EInvalidDigest                       = "InvalidDigest"
	EInvalidEncryptionAlgorithmError     = "InvalidEncryptionAlgorithmError"
	EInvalidObjectName                   = "InvalidObjectName"
	EInvalidPart                         = "InvalidPart"
	EInvalidPartOrder                    = "InvalidPartOrder"
	EInvalidPolicyDocument               = "InvalidPolicyDocument"
	EInvalidTargetBucketForLogging       = "InvalidTargetBucketForLogging"
	EInternalError                       = "InternalError"
	EMalformedXML                        = "MalformedXML"
	EMalformedPOSTRequest                = "MalformedPOSTRequest"
	EMaxPOSTPreDataLengthExceededError   = "MaxPOSTPreDataLengthExceededError"
	EMethodNotAllowed                    = "MethodNotAllowed"
	EMissingArgument                     = "MissingArgument"
	EMissingContentLength                = "MissingContentLength"
	ENoSuchBucket                        = "NoSuchBucket"
	ENoSuchKey                           = "NoSuchKey"
	ENoSuchUpload                        = "NoSuchUpload"
	ENotImplemented                      = "NotImplemented"
	EPreconditionFailed                  = "PreconditionFailed"
	ERequestTimeTooSkewed                = "RequestTimeTooSkewed"
	ERequestTimeout                      = "RequestTimeout"
	ERequestIsNotMultiPartContent        = "RequestIsNotMultiPartContent"
	ESignatureDoesNotMatch               = "SignatureDoesNotMatch"
	ETooManyBuckets                      = "TooManyBuckets"
)

type Error struct {
	XMLName      xml.Name `xml:"Error"`
	ErrNo        string   `xml:"Code"`
	ErrMsg       string   `xml:"Message"`
	HttpStatus   int
	ErrDetailMsg string
}

func (e Error) Error() string {
	return e.ErrMsg
}
