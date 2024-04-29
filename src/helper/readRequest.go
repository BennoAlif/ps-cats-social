package helper

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadRequest(s interface{}, r *http.Request) error {
	b, err := io.ReadAll(r.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(b, s)
}
