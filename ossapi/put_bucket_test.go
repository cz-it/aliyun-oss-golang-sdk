/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"fmt"
	"testing"
)

func TestPutBucketDefault(t *testing.T) {
	if p, err := PutBucketDefault("test-put-bucket2"); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		fmt.Println(p)
	}
}
