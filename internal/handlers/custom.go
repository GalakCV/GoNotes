package handlers

import (
	"errors"
	"net/http"
	"udemy.com/galakcv/aulago/internal/handlers/apperror"
)
type HandlerWithError func (w http.ResponseWriter, r *http.Request) error 

func (f HandlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if err := f(w, r); err != nil {
		var statusErr apperror.StatusError
		if errors.As(err, &statusErr){
			http.Error(w, err.Error(), statusErr.StatusCode())
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

