/**
* Author: CZ cz.theng@gmail.com
 */

package log

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	l, err := NewFileLogger(".ossapilog", "ossapi")
	if err != nil {
		fmt.Errorf("Create Logger Error\n")
		return
	}
	l.SetMaxFileSize(1024 * 1024 * 100) //100MB
	l.SetLevel(LDEBUG)
	l.Error("Error")
	l.Debug("Debug:int a is %d", 10)
	l.Fatal("fatal")
	l.Info("Info")
	l.SetCallDepth(2)
	l.Warning("Warning")
}
