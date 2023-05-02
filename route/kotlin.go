package route

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func Kotlin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    url := config.Server.KOTLIN + getTarget(r, params)
    handleRequest(w, r, url)
}