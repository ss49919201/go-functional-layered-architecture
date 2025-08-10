package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/ss49919201/go-functional-layered-architecture/in-memory/internal/infra"
)

func TestNewRetriveReservation(t *testing.T) {
	type args struct {
		id int

		retriveReservation infra.RetriveReservation
		retriveRoom        infra.RetriveRoom
	}
	tests := []struct {
		name string
		args args
		want *RetriveReservationResult
	}{
		{
			name: "retrive reservation",
			args: args{
				id: 10,
				retriveReservation: func(_ int) (*infra.Reservation, error) {
					return &infra.Reservation{
						ID:                  10,
						RoomID:              20,
						ConfirmedDateTime:   time.Date(2025, 1, 1, 1, 0, 0, 0, time.UTC),
						ReservationDateTime: time.Date(2025, 2, 1, 1, 0, 0, 0, time.UTC),
					}, nil
				},
				retriveRoom: func(_ int) (*infra.Room, error) {
					return &infra.Room{
						ID:   20,
						Name: "normal",
					}, nil
				},
			},
			want: &RetriveReservationResult{
				ID:                  10,
				RoomName:            "normal",
				ConfirmedDateTime:   time.Date(2025, 1, 1, 1, 0, 0, 0, time.UTC),
				ReservationDateTime: time.Date(2025, 2, 1, 1, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "reservation not found",
			args: args{
				id: 10,
				retriveReservation: func(_ int) (*infra.Reservation, error) {
					return nil, infra.ErrNotFound
				},
				retriveRoom: func(_ int) (*infra.Room, error) {
					return nil, nil
				},
			},
			want: nil,
		},
		{
			name: "room not found",
			args: args{
				id: 10,
				retriveReservation: func(_ int) (*infra.Reservation, error) {
					return &infra.Reservation{
						ID:                  10,
						RoomID:              20,
						ConfirmedDateTime:   time.Date(2025, 1, 1, 1, 0, 0, 0, time.UTC),
						ReservationDateTime: time.Date(2025, 2, 1, 1, 0, 0, 0, time.UTC),
					}, nil
				},
				retriveRoom: func(_ int) (*infra.Room, error) {
					return nil, infra.ErrNotFound
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRetriveReservation(tt.args.retriveReservation, tt.args.retriveRoom)(tt.args.id)
			if (tt.want == nil) == (err == nil) {
				t.Errorf("NewRetriveReservation() expects error")
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRetriveReservation() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
