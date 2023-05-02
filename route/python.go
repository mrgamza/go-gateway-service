package route

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func Python(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    url := config.Server.PYTHON + getTarget(r, params)
    handleRequest(w, r, url)
}