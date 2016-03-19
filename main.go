package main

import (
	"fmt"
	"github.com/graphql-go/handler"
	"github.com/sogko/fossasia-2016-graphql-demo/app"
	"net/http"
)

func main() {

	// simplest GraphQL server HTTP handler
	h := handler.New(&handler.Config{
		Schema: &app.Schema,
		Pretty: true,
	})

	// static file server handler to serve GraphiQL in-browser editor
	fs := http.FileServer(http.Dir("static"))

	// wire up our handlers
	http.Handle("/graphql", h)
	http.Handle("/", fs)

	// start our HTTP server
	fmt.Println("Your GraphQL server is now running at http://localhost:8080")
	fmt.Println("You can test `/graphql` endpoint like so: curl -g 'http://localhost:8080/graphql?query={posts{title}}' ")
	http.ListenAndServe(":8080", nil)
}
