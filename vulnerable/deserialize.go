package main

import (
    "encoding/json"
    "net/http"
)

type CreateUserReq struct {
    Username string `json:"username"`
    IsAdmin  bool   `json:"is_admin"` // клиент может назначить себе админские права
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserReq
    // No size limits, no validation
    json.NewDecoder(r.Body).Decode(&req)

    // симулирует создание пользователя с переданным is_admin
    if req.IsAdmin {
        w.Write([]byte("created as admin"))
        return
    }
    w.Write([]byte("created as user"))
}


