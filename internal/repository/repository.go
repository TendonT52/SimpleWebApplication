package repository

import "github.com/TendonT52/SimpleWebApplication/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
