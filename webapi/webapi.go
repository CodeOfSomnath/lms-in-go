package webapi

import (
	"net/http"
	"strconv"
)

type WebHandler struct {
	port    int
	address string
	server http.Server
}

func NewWebHandler(port int, address string) WebHandler  {
	addr := address + strconv.Itoa(port)
	
	return WebHandler {
		port: port,
		address: address,
		server: http.Server{
			Addr: addr,
			Handler: http.NewServeMux(),
			
		},
	}
}