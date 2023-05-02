package route

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func Php(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    url := config.Server.PHP + getTarget(r, params)
    handleRequest(w, r, url)
}