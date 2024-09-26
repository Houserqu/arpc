package ahttp

import (
	"encoding/json"
	"net/http"
)

// HTTP middleware to wrap the response.
func responseWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Intercept response
		rw := &responseWriter{ResponseWriter: w}
		next.ServeHTTP(rw, r)

		// Check if the response is successful
		if rw.statusCode == http.StatusOK {
			var originalResponse interface{}
			json.Unmarshal(rw.body, &originalResponse)

			// Wrap the original response in the desired structure
			wrappedResponse := map[string]interface{}{
				"code": rw.statusCode,
				"data": originalResponse,
			}
			jsonResponse, _ := json.Marshal(wrappedResponse)

			// Write the wrapped response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(rw.statusCode)
			w.Write(jsonResponse)
		} else {
			w.WriteHeader(rw.statusCode)
			w.Write(rw.body)
		}
	})
}

// Custom response writer to capture the response body and status code
type responseWriter struct {
	http.ResponseWriter
	body       []byte
	statusCode int
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	rw.body = data
	return rw.ResponseWriter.Write(data)
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}
