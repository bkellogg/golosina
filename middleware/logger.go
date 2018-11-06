package middleware

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/eaperezc/golosina/framework"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (lrw *loggingResponseWriter) WriteHeader(status int) {
	lrw.status = status
	lrw.ResponseWriter.WriteHeader(status)
}

type RequestLogger struct {
	skipRoutePrefixes map[string]bool
	skipRouteSuffixes map[string]bool
	skipMethods       map[string]bool
	output            io.Writer
}

func LogRequests() framework.Middleware {
	return NewLogger(os.Stdout)
}

func NewLogger(output io.Writer) *RequestLogger {
	return &RequestLogger{
		skipRoutePrefixes: make(map[string]bool),
		skipRouteSuffixes: make(map[string]bool),
		skipMethods:       make(map[string]bool),
		output:            output,
	}
}

func (rl *RequestLogger) GetMWFunc() framework.MiddlewareFunc {
	return framework.MiddlewareFunc(rl.logRequests)
}

func (rl *RequestLogger) SkipMethods(methods ...string) *RequestLogger {
	for _, method := range methods {
		rl.skipMethods[method] = true
	}
	return rl
}

func (rl *RequestLogger) SkipPrefixes(prefixes ...string) *RequestLogger {
	for _, prefix := range prefixes {
		rl.skipMethods[prefix] = true
	}
	return rl
}

func (rl *RequestLogger) SkipSuffixes(suffixes ...string) *RequestLogger {
	for _, suffix := range suffixes {
		rl.skipMethods[suffix] = true
	}
	return rl
}

func (rl *RequestLogger) logRequests(context *framework.Context, next http.Handler) {
	lrw := &loggingResponseWriter{
		ResponseWriter: context.Response.ResponseWriter,
		status:         http.StatusOK,
	}

	startTime := time.Now()
	next.ServeHTTP(lrw, context.Request.Request)

	if _, found := rl.skipMethods[context.Request.Request.Method]; found {
		return
	}

	for prefix := range rl.skipRoutePrefixes {
		if strings.HasPrefix(context.Request.Request.URL.Path, prefix) {
			return
		}
	}

	for suffix := range rl.skipRouteSuffixes {
		if strings.HasSuffix(context.Request.Request.URL.Path, suffix) {
			return
		}
	}

	rl.output.Write(
		[]byte(
			fmt.Sprintf("%d %s %s %s %v",
				lrw.status,
				context.Request.Request.Method,
				context.Request.Request.RequestURI,
				context.Request.Request.RemoteAddr,
				time.Now().Sub(startTime))))
}
