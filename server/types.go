package server

import (
	"sync"
	"zoo-panic/simulation"
)

type ServerRestAgent struct {
	sync.Mutex
	id         string
	addr       string
	simulation *simulation.SimulationMA
}

type AnimalDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Behavior int    `json:"behavior"`
	Escaped  bool   `json:"escaped"`
}

type VisitorDTO struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	IsPanicked bool   `json:"is_panicked"`
}

func ToAnimalDTO(a *simulation.Animal) AnimalDTO {
	return AnimalDTO{
		ID:       a.Id,
		Name:     a.Name,
		Behavior: int(a.Behavior),
		Escaped:  a.Escaped,
	}
}

func ToVisitorDTO(v simulation.Visitor) VisitorDTO {
	return VisitorDTO{
		ID:         int(v.Id),
		Name:       "Visitor",
		IsPanicked: v.State == 2,
	}
}
