package models

//Request data
type Request struct {
	DataAgendamento   string  `json: "dataAgendamento"`
	DataTransferencia string  `json: "dataTransferencia"`
	Valor             float32 `json: "valor"`
}
