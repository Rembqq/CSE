package main

type TimeResponse struct {
	Time string `json:"time"`
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