package response

import (
	"encoding/json"
	"net/http"
)

func Encode(rw http.ResponseWriter, response any, statusCode int) {
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(response)
}
