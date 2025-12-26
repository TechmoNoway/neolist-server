package middleware

import (
	"bytes"
	"io"
	"net/http"
)

func MetadataMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rawBody, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "can not read body", http.StatusBadRequest)
				return
			}

			r.Body = io.NopCloser(bytes.NewReader(rawBody))

			if len(bytes.TrimSpace(rawBody)) == 0 {
				http.Error(w, "invalied JSON body", http.StatusBadRequest)
				return
			}

			

		})
	}

}
