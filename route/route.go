package route

import (
    cfg "gateway-service/config"
    "github.com/julienschmidt/httprouter"
    "io/ioutil"
    "net/http"
    "strings"
)

var config *cfg.Config

// - Public

func Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    _, _ = w.Write([]byte("Hello, gateway"))
}

// - Private

func getTarget(r *http.Request, params httprouter.Params) string {
    path := params.ByName("path")
    pullUrl := r.URL.String()
    pullLength := len(pullUrl)
    index := strings.Index(pullUrl, path)
    return pullUrl[index:pullLength]
}

func handleRequest(w http.ResponseWriter, r *http.Request, url string) {
    request, err := http.NewRequest(r.Method, url, r.Body)
    if request == nil || err != nil {
        write404(w)
        return
    }
    
    request.Header.Add("Accept", r.Header.Get("Accept"))
    
    client := &http.Client{}
    response, err := client.Do(request)
    
    if response == nil || err != nil {
        handle404(w, response)
        return
    }
    
    read, _ := ioutil.ReadAll(response.Body)
    contentType := response.Header.Get("Content-Type")
    response.Body.Close()
    w.Header().Set("Content-Type", contentType)
    w.WriteHeader(response.StatusCode)
    write(w, read)
}

func handle404(w http.ResponseWriter, response *http.Response) {
    if response != nil {
        response.Body.Close()
    }
    
    write404(w)
}

func write404(w http.ResponseWriter) {
    w.WriteHeader(http.StatusNotFound)
    write(w, []byte("404 Not Found"))
}

func write(w http.ResponseWriter, message []byte) {
    _, _ = w.Write(message)
}

// - Init

func New()  {
    var err error
    config, err = cfg.New("./config/config.yaml")
    if err != nil {
        panic(err)
    }
}