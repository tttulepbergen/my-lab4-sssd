package main

import (
    "log"
    "net/http"
    "os"
)
func APIKey() string {
    if k := os.Getenv("APP_API_KEY"); k != "" {
        return k
    }
    return "TEST_KEY_PLACEHOLDER"
}

func configHandler(w http.ResponseWriter, r *http.Request) {
    // Не возвращает  секреты клиенту. Демонстрируем только, то что ключ загружен

    if APIKey() == "" {
        http.Error(w, "config not set", http.StatusServiceUnavailable)
        return

    }
    log.Println("config endpoint accessed; API key present (hidden)")
    w.Write([]byte("config ok"))
}


