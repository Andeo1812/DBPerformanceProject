package pkg

import (
	"context"
	"net/http"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type HTTPMiddleware struct {
	log *logrus.Logger
}

func NewHTTPMiddleware(log *logrus.Logger) *HTTPMiddleware {
	return &HTTPMiddleware{
		log: log,
	}
}

func (m *HTTPMiddleware) SetDefaultLoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), LoggerKey, m.log)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *HTTPMiddleware) UpdateDefaultLoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value(LoggerKey).(*logrus.Logger)
		if !ok {
			logrus.Fatal("GetLoggerContext: err convert context -> *logrus.Logger")
		}

		requestID := uuid.NewV4().String()

		start := time.Now()

		upgradeLogger := logger.WithFields(logrus.Fields{
			"url":         r.URL.Path,
			"method":      r.Method,
			"remote_addr": r.RemoteAddr,
			RequestID:     requestID,
		})

		ctx := context.WithValue(r.Context(), LoggerKey, upgradeLogger)

		ctx = context.WithValue(ctx, RequestIDKey, requestID)

		h.ServeHTTP(w, r.WithContext(ctx))

		executeTime := time.Since(start).Milliseconds()
		upgradeLogger.Infof("work time [ms]: %v", executeTime)
	})
}

func (m *HTTPMiddleware) SetSizeRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		strLength := r.Header.Get("Content-Length")
		if strLength == "" {
			h.ServeHTTP(w, r)
			return
		}

		length, err := strconv.Atoi(strLength)
		if err != nil {
			DefaultHandlerHTTPError(r.Context(), w, ErrConvertLength)
			return
		}

		if length > BufSizeRequest {
			DefaultHandlerHTTPError(r.Context(), w, ErrBigRequest)
			return
		}

		h.ServeHTTP(w, r)
	})
}
