package lib

import (
	"fmt"
	"net/http"
)

func RegisterHandler(path string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(path, handlerFunc)
	fmt.Printf("Registered handler for path: %s\n", path)
}

func StartServer(port string) {
	addr := fmt.Sprintf(":%s", port)
	panic(http.ListenAndServe(addr, nil))
}
