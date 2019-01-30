package route

import (
    "encoding/json"
    "net/http"
    "config"
)

type TagRequest struct {
    Name string
}

func TagHandler(w http.ResponseWriter, req *http.Request) {
    _, tags, _ := config.Config()
    if req.Method == "POST" {
        decoder := json.NewDecoder(req.Body)
        var request TagRequest
        err := decoder.Decode(&request)
        if err != nil {
            http.Error(w, "Error converting results to json", http.StatusInternalServerError)
        }
        tags.Create(request.Name)
        return
    }

    if req.Method == "GET" {
        jsonBody, err := json.Marshal(tags.GetAll().Value)
        if err != nil {
            http.Error(w, "Error converting results to json", http.StatusInternalServerError)
        }
        w.Write(jsonBody)

        return
    }

    http.Error(w, "Error converting results to json", http.StatusNotFound)
}