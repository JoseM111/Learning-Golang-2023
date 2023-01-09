package main

import (
	"fmt"
	. "net/http"
)

func main() {
	HandleFunc("/", func(writer ResponseWriter, request *Request) {
		fmt.Println("Hola MUndo•..")
	},
	)
	err := ListenAndServe(":9090", nil)
	if err != nil {
		return
	}
}
