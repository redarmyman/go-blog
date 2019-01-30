package route

import (
    "encoding/json"
    "net/http"
    "config"
)

type CategoryRequest struct {
    Name string
}

func CategoryHandler(w http.ResponseWriter, req *http.Request) {
    categories, _, _ := config.Config()
    if req.Method == "POST" {
        decoder := json.NewDecoder(req.Body)
        var request CategoryRequest
        err := decoder.Decode(&request)
        if err != nil {
            http.Error(w, "Error converting results to json", http.StatusInternalServerError)
        }
        categories.Create(request.Name)
        return
    }

    if req.Method == "GET" {
        jsonBody, err := json.Marshal(categories.GetAll().Value)
        if err != nil {
            http.Error(w, "Error converting results to json", http.StatusInternalServerError)
        }
        w.Write(jsonBody)

        return
    }

    http.Error(w, "Error converting results to json", http.StatusNotFound)
}