package firstservice

import (
	"github.com/gorilla/mux"
)

func TopicRoute(r *mux.Router) {

	r.HandleFunc("/topic/add", AddUSer).Methods("POST")
	r.HandleFunc("/topic/delete", DeleteUSer).Methods("POST")

}
