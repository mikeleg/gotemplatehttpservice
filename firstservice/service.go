package firstservice

import (
	"net/http"
)

func AddUSer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func DeleteUSer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
