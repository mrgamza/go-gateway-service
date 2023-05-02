package route

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func Go(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    url := config.Server.GO + getTarget(r, params)
    handleRequest(w, r, url)
}