package examplehttp

import (
	"fmt"
	"net/http"
)

func StatServer() {

	http.HandleFunc("/hello", helloworld)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("testing"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Println()
			w.Write([]byte(`{"success": true}`))
		} else {
			w.Write([]byte(`{"id": 23, "name": "Vikram"}`))
		}
	})
	http.ListenAndServe(":8080", nil)
}

func helloworld(w http.ResponseWriter, req *http.Request) {
	if value, ok := req.Header["Accept"]; ok {
		fmt.Println("header = ", value)
	}
	fmt.Println("My name is BIkram.")
}
