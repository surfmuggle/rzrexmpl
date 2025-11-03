package main

import (
	"fmt"
	"net/http"
	"rzrexmpl/models"
	// "rzrexmpl/tpl"
)

func main() {
	user := &models.User{}
	user.Name = "go"
	user.Email = "hello@world.com"
	user.Intro = "I love gorazor!"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "after running gorazor uncomment line below %v", user)
		// fmt.Fprintf(w, tpl.Home(1, user))
	})

	http.ListenAndServe(":8080", nil)
}
