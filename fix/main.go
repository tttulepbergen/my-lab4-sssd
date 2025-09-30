package main
import (
    "log"
    "net/http"

    "time"
)

func withRecover(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        defer func() {
            if err := recover(); err != nil {
                log.Printf("panic recovered: %v", err)
                http.Error(w, "internal server error", http.StatusInternalServerError)
            }

        }()
        next.ServeHTTP(w, r)
    })
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
    // та же симуляция ошибкино теперь клиент не увидит деталей

    panic("some internal error happened")
}
func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", rootHandler)
    mux.HandleFunc("/config", configHandler)
    mux.HandleFunc("/create-user", createUser)

    srv := &http.Server{
        Addr:         ":8080",
        Handler:      withRecover(mux),
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,

        IdleTimeout:  60 * time.Second,
    }

    log.Println("[fix] Listening on :8080")
    log.Fatal(srv.ListenAndServe())
}


