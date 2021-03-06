/**
* Author: CZ cz.theng@gmail.com
 */

package log

import (
	"fmt"
	"testing"
)

func TestDevice(t *testing.T) {
	fd, err := newFileDevice("a.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fd)
	fd.SetFileName("abc.txt")
	fd, err = newFileDevice("/dev/a.txt")
}
