package simulation

import (
	"fmt"
	"math/rand"
)

// État du visiteur
type VisitorState int

const (
	Calm VisitorState = iota
	Alert
	Panicked
)

// Visitor représente un visiteur dans le zoo
type Visitor struct {
	Id         int64
	State      VisitorState
	nearbyExit bool
	isInZoo    bool
	nbEscaped  int
}

// NewVisitor crée un nouveau visiteur
func NewVisitor(id int64) *Visitor {
	return &Visitor{Id: id, State: Calm, nearbyExit: false, isInZoo: true}
}

// Start initialise le visiteur dans la simulation
func (v *Visitor) Start() {
	fmt.Printf("Visiteur %d entre dans le zoo\n", v.Id)
}

// Percept permet au visiteur de percevoir l'environnement
func (v *Visitor) Percept(env *Environment) {
	v.nbEscaped = len(env.GetEscapedAnimals())
}

// Deliberate décide de la prochaine action en fonction de l'état du visiteur
func (v *Visitor) Deliberate() {
	if v.nbEscaped == 0 {
		v.State = Calm
	} else if v.nbEscaped < 3 {
		v.State = Alert
	} else {
		v.State = Panicked
	}
	// Mise à jour de la perception de la sortie proche (exemple)
	v.nearbyExit = rand.Float64() < 0.5 // suppose qu'il y a une chance de trouver une sortie proche
}

// Act exécute l'action en fonction de la décision prise
func (v *Visitor) Act(env *Environment) {
	if v.isInZoo {
		switch v.State {
		case Calm:
			fmt.Printf("Visiteur %d se balade tranquillement dans le Zoo.\n", v.Id)
		case Alert:
			// Le visiteur marche rapidement ou cherche à s'éloigner des animaux échappés
			fmt.Printf("Visiteur %d se dirige vers une zone plus éloignée des animaux.\n", v.Id)
		case Panicked:
			if v.nearbyExit {
				// Le visiteur quitte le zoo s'il est proche d'une sortie
				v.isInZoo = false
				fmt.Printf("Visiteur %d quitte le zoo en panique par la sortie la plus proche.\n", v.Id)
			} else {
				// Sinon, il continue de paniquer et de chercher une sortie
				fmt.Printf("Visiteur %d court partout, essayant de trouver une issue.\n", v.Id)
			}
		}
	}
}
