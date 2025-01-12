package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipHandler struct {
}

func (gh *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			wrw := NewWrappedResponseWriter(rw)

			next.ServeHTTP(wrw, r)

			defer wrw.Flush()
			return

		}
		next.ServeHTTP(rw, r)

	})

}

type WrappedResponseWriter struct {
	rw      http.ResponseWriter
	gzip_rw *gzip.Writer
}

func NewWrappedResponseWriter(rw http.ResponseWriter) *WrappedResponseWriter {
	gz := gzip.NewWriter(rw)
	return &WrappedResponseWriter{rw: rw, gzip_rw: gz}
}

func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.rw.Header()
}

func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wr.gzip_rw.Write(d)
}

func (wr *WrappedResponseWriter) WriteHeader(statusCode int) {
	wr.rw.WriteHeader(statusCode)
}

func (wr *WrappedResponseWriter) Flush() {
	wr.gzip_rw.Flush()
	wr.gzip_rw.Close()
}
