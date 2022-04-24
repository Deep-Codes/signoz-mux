package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	middleware "go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"

	"github.com/gorilla/mux"
)

type User struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

func main() {
	r := mux.NewRouter()
	r.Use(middleware.Middleware("SIGNOZ_MUX_TEST"))
	r.HandleFunc("/users", AllUsers)
	log.Fatal(http.ListenAndServe(":4444", r))
}

var Users = []User{
	{Name: "Deep", Id: 1},
	{Name: "Duma", Id: 2},
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: /users")
	json.NewEncoder(w).Encode(Users)
}
