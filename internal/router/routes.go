package router

import (
	"fmt"
	"net/http"
	"weekend.side/SocialMedia/internal/controller"
)

func Initialize() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)

	//router:=gin.New()
	//rootGroup := router.Group("/")
	//rootGroup.GET("", createAccount)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "I'm up!")
	case "/account":
		if r.Method == "POST" {
			controller.CreateAccount(w, r)
		}
		if r.Method == "DELETE" {
			controller.DeleteAccount(w, r)
		}
	case "/comment":
		if r.Method == "POST" {
			controller.CommentOnPost(w, r)
		}
	case "/post":
		if r.Method == "POST" {
			controller.CreatePost(w, r)
		}
		if r.Method == "GET" {
			controller.FetchPosts(w, r)
		}
	default:
		fmt.Print("hey")
	}
}
