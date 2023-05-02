package route

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func Ruby(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    url := config.Server.RUBY + getTarget(r, params)
    handleRequest(w, r, url)
}