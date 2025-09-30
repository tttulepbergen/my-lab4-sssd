package main
import (
    "encoding/json"
    "io"
    "net/http"
)

type CreateUserReq struct {

    Username string `json:"username"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
    // ограничиваем размер тела
    r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1MB

    dec := json.NewDecoder(r.Body)

    dec.DisallowUnknownFields()

    var req CreateUserReq
    if err := dec.Decode(&req); err != nil {
        if err == io.EOF {
            http.Error(w, "empty body", http.StatusBadRequest)
            return
        }

        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }

    if req.Username == "" {
        http.Error(w, "username required", http.StatusBadRequest)

        return
    }

    // флаг is_admin определяется только серверной логикой не из запроса
    w.Write([]byte("created"))
}


