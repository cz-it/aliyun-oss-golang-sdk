/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"errors"
	"net/http"
)

var (
	// global Access Key ID
	accessKeyID string
	// global Access Key Secret
	accessKeySecret string
)

var (
	EARG = errors.New("Invalied Argument !")
)

// http client for http request
var httpClient http.Client

/**
* Init ossapi with Access Key's ID and secret
* @param ID : Access Key's ID
* @param secret : Access Key's secret
 */
func Init(ID string, secret string) error {
	if "" == ID || "" == secret {
		return EARG
	}
	accessKeyID = ID
	accessKeySecret = secret
	return nil
}
