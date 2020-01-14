package responsecustom

import (
	"constants"
	"net/http"
)

// ResponseCustom function
func ResponseCustom(w http.ResponseWriter, status int, rs []byte) {
	w.Header().Set(constants.CntType, constants.AppJson)
	w.WriteHeader(status)
	w.Write(rs)
}
