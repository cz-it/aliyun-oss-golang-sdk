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
	ENone = "None"
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
