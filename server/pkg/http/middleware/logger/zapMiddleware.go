package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ErrMessage struct {
	Message string `json:"message"`
}

// bodyDumpResponseWriter wraps gin.ResponseWriter to capture response body.
type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

// Write writes to the response writer.
func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	// WriteHeader(int) is not called if our response implicitly returns 200 OK, so
	// we default to that status code.
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func ZapLogger(log *zap.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			//get request body
			reqBody := []byte{}
			contentType := r.Header.Get("Content-Type")
			if r.Body != nil {
				reqBody, _ = io.ReadAll(r.Body)
			}
			mapData := make(map[string]interface{})

			if len(reqBody) > 0 && contentType == "application/json" {
				if err := json.Unmarshal(reqBody, &mapData); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(ErrMessage{Message: "BAD_REQUEST"})
					return
				}
			}
			//reset request body
			r.Body = io.NopCloser(bytes.NewBuffer(reqBody))

			//masking credentials field
			doc := &Document{}
			bodyMasked := doc.throughMap(mapData)

			//get response body
			resBody := new(bytes.Buffer)
			mw := io.MultiWriter(w, resBody)
			writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: w}
			w = writer

			lrw := NewLoggingResponseWriter(w)
			next.ServeHTTP(lrw, r)

			// make map for res body
			resBodyMap := make(map[string]interface{})

			if len(resBody.Bytes()) > 0 {
				if err := json.Unmarshal(resBody.Bytes(), &resBodyMap); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(ErrMessage{Message: "INTERNAL_SERVER_ERROR"})
					return
				}
			}
			resBodyMasked := doc.throughMap(resBodyMap)

			id := r.Header.Get("X-Request-ID")

			fields := []zapcore.Field{
				zap.Int("status", lrw.statusCode),
				zap.String("latency", time.Since(start).String()),
				zap.String("id", id),
				zap.String("method", r.Method),
				zap.String("uri", r.RequestURI),
				zap.String("host", r.Host),
				zap.Any("body", bodyMasked),
				zap.Any("respBody", resBodyMasked),
				zap.String("remote_ip", r.RemoteAddr),
				zap.Any("headers", r.Header),
				zap.Any("request_body", reqBody),
			}

			n := lrw.statusCode
			switch {
			case n >= 500:
				log.Error("Server Error", fields...)
			case n >= 400:
				log.Warn("Client Error", fields...)
			case n >= 300:
				log.Info("Redirection", fields...)
			default:
				log.Info("Success", fields...)
			}
		})
	}
}
