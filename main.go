package main

import (
	"fmt"
	"net/http"

	"sortout-user/domain"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	u := domain.NewUserModel("hiroharu", "hh@exam.com")
	name := randomdata.SillyName()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "user: %s, and %s", u.Name, name)
	})
	http.ListenAndServe(":8080", nil)
}
