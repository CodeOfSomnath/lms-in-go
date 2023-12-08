package webapi

import (
	"fmt"
	"log"
	"net/http"
)

type WebHandler struct {
	port    int
	address string
	Mux http.ServeMux
}

func NewWebHandler(port int) WebHandler  {
	addr := fmt.Sprintf("localhost:%d", port)
	return WebHandler {
		port: port,
		address: addr,
		Mux: *http.NewServeMux(),
	}
	
}


func (w *WebHandler) Start()  {
	log.Println(http.ListenAndServe(w.address, &w.Mux))
}

