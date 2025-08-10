package controller

import (
	"net/http"

	"github.com/ss49919201/go-functional-layerd-architecture/in-memory/service"
)

type RetriveReservation func(w http.ResponseWriter, r *http.Request)

func NewRetriveReservation(
	retriveReservation service.RetriveReservation,
) RetriveReservation

func retriveReservation(w http.ResponseWriter, r *http.Request) {}
