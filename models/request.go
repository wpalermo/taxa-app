package models

import "time"

//Request data
type Request struct {
	dataAgendamento   time.Time `json: "dataAgendamento"`
	dataTransferencia time.Time `json: "dataTransferencia"`
	valor             float32   `json: "valor"`
}
