package route

import (
    "encoding/json"
    "net/http"
    "config"
)

type PostRequest struct {
    Title      string
    Text       string
    CategoryID int
    Tags       []int
}

func PostHandler(w http.ResponseWriter, req *http.Request) {
    _, _, posts := config.Config()
    if req.Method == "POST" {
        decoder := json.NewDecoder(req.Body)
        var request PostRequest
        err := decoder.Decode(&request)
        if err != nil {
            http.Error(w, "Error converting results to json", http.StatusInternalServerError)
        }

        posts.Create(request.Title, request.Text, request.CategoryID, request.Tags)
        return
    }

    if req.Method == "GET" {
        jsonBody, err := json.Marshal(posts.GetAll().Value)
        if err != nil {
            http.Error(w, "Error converting results to json", http.StatusInternalServerError)
        }
        w.Write(jsonBody)

        return
    }

    http.Error(w, "Error converting results to json", http.StatusNotFound)
}