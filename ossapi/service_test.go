/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"testing"
)

func TestGetService(t *testing.T) {
	if 0 == GetService() {
		t.Log("Get Service Success!")
	}
}
