/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"fmt"
	"testing"
)

func TestPutBucketACL(t *testing.T) {
	if err := PutBucketACL("test-put-bucket", P_Private, L_Hangzhou); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("PutBucketACL Success!")
	}

}
