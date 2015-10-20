/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"net/http"
)

type Responser interface {
}

type Response struct {
	httpRsp *http.Response
}
