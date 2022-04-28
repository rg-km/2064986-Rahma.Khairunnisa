package main

import (
	"fmt"
	"net/http"
)

var (
	username = "aditira"
	password = "aditira123"
)

func main() {
	fmt.Println("Starting Server at port :8080")
	http.ListenAndServe(":8080", Routes())
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok {
			// TODO: answer here
			fmt.Println("Error parsing basic auth")
			w.WriteHeader(401)
		    return
		}
		if u != username {
			// TODO: answer here
			fmt.Printf("Username provided is correct: %s\n", u)
			w.WriteHeader(401)
			return
		}
		if p != password {
			// TODO: answer here
			fmt.Printf("Password provided is correct: %s\n", u)
		    w.WriteHeader(401)
		    return
		}
		fmt.Printf("Username: %s\n", u)
		fmt.Printf("Password: %s\n", p)

		// TODO: answer here
		w.WriteHeader(200)

	})

	return mux
}

// Encode auth aditira:aditira123 disini -> https://www.base64encode.org/

// $ curl -v -X POST http://localhost:8080/login -H "Authorization: Basic YWRpdGlyYTphZGl0aXJhMTIz"
