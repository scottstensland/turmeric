package main

import (
	"fmt"
	"net/http"
)

func main() {

	port := "8080"
	host_port := "localhost:" + port

	fmt.Println("Server running on http://" + host_port)

	http.ListenAndServe(":"+port, http.FileServer(http.Dir(".")))
}
