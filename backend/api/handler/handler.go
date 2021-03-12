package handler

import (
	"encoding/json"
	"net/http"
)

func encodeResponse(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}

func decodeRequestBody(request *http.Request, payload interface{}) {
	json.NewDecoder(request.Body).Decode(payload)
}
