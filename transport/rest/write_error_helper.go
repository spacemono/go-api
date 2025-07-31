package rest

import (
	"encoding/json"
	"github.com/spacemono/go-api/transport"
	"net/http"
)

func writeError(w http.ResponseWriter, err error) {
	apiError := transport.FromError(err)
	w.WriteHeader(apiError.Code)
	json.NewEncoder(w).Encode(apiError.Message)
}
