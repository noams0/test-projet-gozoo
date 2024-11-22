package simulation

import (
	"fmt"
	"log"
)

type Action int64

const (
	Wait = iota
	Escape
	Act
)

// TypeAnimal définit les types de comportements possibles pour un animal
type TypeAnimal int

const (
	Passive TypeAnimal = iota
	Aggressive
	Clever
)

// AnimalBehavior interface définit les comportements de base de chaque animal
type AnimalBehavior interface {
	Escape(env *Environment)
	Act(env *Environment)
	Perceive(env *Environment)
}

// Animal représente un animal dans le zoo
type Animal struct {
	Id       int
	Name     string
	Behavior TypeAnimal
	decision Action
	cage     *Cage
	Escaped  bool
}

// NewAnimal crée un nouvel animal avec un comportement spécifique
func NewAnimal(id int, name string, behavior TypeAnimal, cage *Cage) *Animal {
	return &Animal{
		Id:       id,
		Name:     name,
		Behavior: behavior,
		cage:     cage,
		Escaped:  false,
	}
}

// Implémentation des méthodes de l'interface Agent pour Animal

// Start initialise l'animal dans la simulation
func (a *Animal) Start() {
	fmt.Printf("L'animal %s démarre dans la simulation.\n", a.Name)
}

// Percept permet à l'animal de percevoir son environnement
func (a *Animal) Percept(env *Environment) {
	escapedAnimals := env.GetEscapedAnimals()
	fmt.Printf("L'animal %s perçoit %d animaux échappés.\n", a.Name, len(escapedAnimals))
}

// Deliberate détermine la prochaine action de l'animal en fonction de son comportement
func (a *Animal) Deliberate() {
	// Selon le comportement, l'animal prend une décision
	if !a.Escaped && a.cage.isOpen {
		a.decision = Escape
	} else if a.Escaped {
		a.decision = Act
	} else {
		a.decision = Wait
	}
}

// Act exécute une action spécifique selon le comportement de l'animal
func (a *Animal) Act(env *Environment) {
	switch a.decision {
	case Wait:
		fmt.Printf("L'animal %s est enfermé et ne fait rien.\n", a.Name)
	case Escape:
		a.Escape(env)
	case Act:
		switch a.Behavior {
		case Passive:
			fmt.Printf("L'animal %s est passif et ne fait rien.\n", a.Name)
		case Aggressive:
			fmt.Printf("L'animal %s attaque les visiteurs à proximité !\n", a.Name)
			a.attackVisitors(env)
		case Clever:
			fmt.Printf("L'animal %s tente d'ouvrir d'autres cages !\n", a.Name)
			a.openNearbyCages(env)
		}
	}
}

// Méthodes additionnelles pour le comportement spécifique de l'animal

// Escape fait sortir l'animal de sa cage si elle est ouverte
func (a *Animal) Escape(env *Environment) {
	a.Escaped = true
	fmt.Printf("L'animal %s s'est échappé de la cage %d !\n", a.Name, a.cage.id)
	env.AddEscapedAnimal(a)
}

// attackVisitors gère le comportement d'attaque d'un animal agressif
func (a *Animal) attackVisitors(env *Environment) {
	//fmt.Println("Les visiteurs sont en panique !")
	// Code pour marquer les visiteurs comme en panique dans l'environnement
}

// openNearbyCages gère l'action d'ouverture de cages par un animal malin
func (a *Animal) openNearbyCages(env *Environment) {

	for _, cage := range env.cages {
		log.Println(cage)
		if !cage.isOpen {
			env.OpenCage(cage.id)
			fmt.Printf("L'animal malin %s a ouvert la cage %d !\n", a.Name, cage.id)
			break // Ouvrir une cage suffit pour cet acte
		}
	}
}
