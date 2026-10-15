package server

import (
    "context"
    "log/slog"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

// Run starts a health-checkable scaffold. No Redis connection or job processing yet.
func Run(role string) {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })
    // Scaffold readiness only indicates HTTP availability. Add Redis PING later.
    mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })
    if role == "api" {
        mux.HandleFunc("POST /jobs", func(w http.ResponseWriter, r *http.Request) {
            http.Error(w, "TODO: enqueue a word-count job", http.StatusNotImplemented)
        })
        mux.HandleFunc("GET /jobs/{id}", func(w http.ResponseWriter, r *http.Request) {
            http.Error(w, "TODO: retrieve job status and result", http.StatusNotImplemented)
        })
    }
    addr := os.Getenv("HTTP_ADDR")
    if addr == "" { addr = ":8080" }
    srv := &http.Server{Addr: addr, Handler: mux, ReadHeaderTimeout: 5*time.Second}
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()
    go func() {
        <-ctx.Done()
        shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        if err := srv.Shutdown(shutdownCtx); err != nil { slog.Error("shutdown", "error", err) }
    }()
    slog.Info("starting scaffold", "role", role, "addr", addr)
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        slog.Error("serve", "error", err)
        os.Exit(1)
    }
}
