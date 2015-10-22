/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"fmt"
	"testing"
)

func TestOpenBucketLogging(t *testing.T) {
	if err := OpenBucketLogging("test-put-bucket", L_Hangzhou, "test-put-bucket", "test-"); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("Open Bucket Log Success")
	}
}

func TestCloseBucketLogging(t *testing.T) {
	if err := CloseBucketLogging("test-put-bucket", L_Hangzhou); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("Close Bucket Log Success")
	}

}
