package app

import (
	"fmt"
	"log"
	"net/http"
)

func response200(w http.ResponseWriter, msg any) {
	response(w, http.StatusOK, msg)
}

func response400(w http.ResponseWriter, err error) {
	response(w, http.StatusBadRequest, err)
}

func response500(w http.ResponseWriter, err error) {
	log.Printf(err.Error())
	response(w, http.StatusInternalServerError, err)
}

func response(w http.ResponseWriter, code int, msg any) {
	w.WriteHeader(code)
	_, err := fmt.Fprint(w, msg)
	if err != nil {
		log.Printf("[RESPONSE]: %s", err)
	}
}
