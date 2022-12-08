package main

import (
	"net/http"
	"os"

	"golang.org/x/exp/slog"
)

func process(w http.ResponseWriter, r *http.Request) {
	logger  := slog.FromContext(r.Context())
	message := r.Header
	logger.Info("received the message", "key", message)

}

func main() {
	// logger := slog.New(slog.NewJSONHandler(os.Stdout))
	// logger.Info("get msg", "key", 2)
	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		logger := slog.Default().With("URL", r.URL)
		logger.Info("start")
		ctx := slog.NewContext(r.Context(), logger)
		r = r.WithContext(ctx)
		process(w, r)
	})

	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	err := http.ListenAndServe(":9876", nil)
	if err != nil {
		logger.Error("failed to listen the port", err)
		return
	}
}
