package responsecustom

import "net/http"

// ResponseCustom function
func ResponseCustom(w http.ResponseWriter, status int, rs []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(rs)
}
