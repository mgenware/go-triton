package systemRoute

import (
	"context"
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
	ctx := r.Context()
	if res := recover(); res != nil {
		ctx = context.WithValue(ctx, "recoverReturn", res)
		ctx = context.WithValue(ctx, "recoverStack", string(debug.Stack()))
		panicHandler(w, r)
	}
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	msgValue := ctx.Value("recoverReturn")
	stackInfo := ctx.Value("recoverStack")
	msg := fmt.Sprintf("Fatal error：%v\nDetails: %v", msgValue, stackInfo)

	w.WriteHeader(http.StatusInternalServerError)
	fmt.Print(w, msg)
}
