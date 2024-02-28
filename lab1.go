package main

import (
  "encoding/json"
  "log"
  "net/http"
  "time"
)

type TimeResponse struct {
  Time string json:"time"
}

func main() {
  http.HandleFunc("/", api)
  server := &http.Server{
    Addr: ":8795",
  }

  err := server.ListenAndServe()
  if err != nil {
    log.Fatal(err)
  }
}

func api(w http.ResponseWriter, r *http.Request) {

  currentTime := time.Now().Format(time.RFC3339)

  response := TimeResponse{
    Time: currentTime,
  }

  jsonData, err := json.Marshal(response)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(jsonData)
}
