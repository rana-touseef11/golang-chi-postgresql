package middleware

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
)

var log = slog.New(
	slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}),
)

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

func getIP(r *http.Request) string {
	ip := r.Header.Get("CF-Connecting-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			status:         200,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start).Milliseconds()

		attrs := []any{
			"method", r.Method,
			"path", r.RequestURI,
			"status", rw.status,
			"size", rw.size,
			"ip", getIP(r),
			"duration", fmt.Sprintf("%v MS", duration),
		}

		// Choose log level based on status
		switch {
		case rw.status >= 500:
			log.Error("server_error", attrs...)
		case rw.status >= 400:
			log.Warn("client_error", attrs...)
		default:
			log.Info("request", attrs...)
		}
	})
}
