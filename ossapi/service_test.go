/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"fmt"
	"testing"
)

func TestGetService(t *testing.T) {
	if buckets, err := GetService(); err != nil {
		if err != nil {
			t.Log("Error :", err.Error())
			if error, ok := err.(*Error); ok {
				t.Log("error")
				fmt.Println(error.ErrNo, error.HttpStatus, error.ErrMsg, error.ErrDetailMsg)
			} else {
				t.Log("ok", ok)
			}
		}
	} else {
		fmt.Println(buckets)
	}
}
