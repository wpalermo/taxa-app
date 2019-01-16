package interfaces

import "time"

type ITaxaService interface {
	CalcularTaxa(dataTransferencia time.Date, dataAgendamento time.Date, valor float32) float32
}
