package route

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func Java(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    url := config.Server.JAVA + getTarget(r, params)
    handleRequest(w, r, url)
}