package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ss49919201/go-functional-layered-architecture/in-memory/internal/controller"
	"github.com/ss49919201/go-functional-layered-architecture/in-memory/internal/infra"
	"github.com/ss49919201/go-functional-layered-architecture/in-memory/internal/service"
)

func NewHandler() http.Handler {
	retriveReservationInfra := infra.NewRetriveReservation()
	retriveRoomInfra := infra.NewRetriveRoom()

	retriveReservationService := service.NewRetriveReservation(
		retriveReservationInfra,
		retriveRoomInfra,
	)

	retriveReservationController := controller.NewRetriveReservation(
		retriveReservationService,
	)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /reservation/{id}", retriveReservationController)
	return mux
}

func ListenAndServe(port int) error {
	handler := NewHandler()

	fmt.Printf("Server starting on :%d\n", port)
	return http.ListenAndServe(":"+strconv.Itoa(port), handler)
}
