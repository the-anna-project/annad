package text

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
)

// stream text

func streamTextDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.StreamTextRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, maskAny(err)
	}
	return request, nil
}

func streamTextEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return maskAny(err)
	}
	return nil
}
