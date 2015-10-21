/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"net/http"
)

type Response struct {
	Result  error
	httpRsp *http.Response
}
