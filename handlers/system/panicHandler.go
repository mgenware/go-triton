package system

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer recoverFromPanic(w, r)
		next.ServeHTTP(w, r)
	})
}

func recoverFromPanic(w http.ResponseWriter, r *http.Request) {
	if res := recover(); res != nil {
		panicHandler(w, r, res, debug.Stack())
	}
}

func panicHandler(w http.ResponseWriter, r *http.Request, result interface{}, stack []byte) {
	w.WriteHeader(http.StatusInternalServerError)

	msg := fmt.Sprintf("Fatal error：%v\nDetails: %v", result, string(stack))

	fmt.Print(w, msg)
}
