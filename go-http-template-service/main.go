package main

import (
	"strings"
	"test/firstservice"
	"fmt"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	initRoute(r)

	fmt.Println("Server start")
	http.ListenAndServe(":8080", handlers.CompressHandler(r))
}


func initRoute(r *mux.Router) {
	myservicetest.TopicRoute(r)
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Printf("ROUTE:%s Methods: %s", pathTemplate,strings.Join(methods, ","))
			fmt.Println()
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

}