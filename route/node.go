package route

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func Node(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    url := config.Server.NODE + getTarget(r, params)
    handleRequest(w, r, url)
}