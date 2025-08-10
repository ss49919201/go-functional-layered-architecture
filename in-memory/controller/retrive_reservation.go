package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/ss49919201/go-functional-layerd-architecture/in-memory/infra"
	"github.com/ss49919201/go-functional-layerd-architecture/in-memory/service"
)

type RetriveReservation func(w http.ResponseWriter, r *http.Request)

func NewRetriveReservation(
	retriveReservation service.RetriveReservation,
) RetriveReservation {
	return func(w http.ResponseWriter, r *http.Request) {
		retriveReservationHandler(w, r, retriveReservation)
	}
}

func retriveReservationHandler(
	w http.ResponseWriter,
	r *http.Request,
	retriveReservation service.RetriveReservation,
) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}

	result, err := retriveReservation(id)
	if err != nil {
		if errors.Is(err, infra.ErrNotFound) {
			http.Error(w, "reservation not found", http.StatusNotFound)
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
