package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type WithError struct {
	Error string `json:"error"`
}

func UserError(w http.ResponseWriter, err error) {
	res := WithError{Error: err.Error()}
	msg, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err := w.Write(msg); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func DevError(w http.ResponseWriter, err error) {
	res := WithError{Error: err.Error()}
	msg, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(msg); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func Create(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func CreateWithMessage(w http.ResponseWriter, message []byte) {
	Create(w)

	if _, err := w.Write(message); err != nil {
		log.Println(err.Error())
	}
}

func Ok(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func OkWithMessage(w http.ResponseWriter, message []byte) {
	Ok(w)

	if _, err := w.Write(message); err != nil {
		log.Println(err.Error())
	}
}

