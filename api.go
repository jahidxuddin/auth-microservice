package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type ApiServer struct {
	port uint
}

func NewApiServer(port uint) *ApiServer {
	return &ApiServer{port: port}
}

func (s *ApiServer) Start() {
	port := strconv.FormatUint(uint64(s.port), 10)
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	err := http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
