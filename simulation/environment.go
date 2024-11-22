package simulation

import (
	"fmt"
	"log"
)

// Cage représente une cage dans le zoo qui peut contenir un animal
type Cage struct {
	id     int
	isOpen bool
}

func NewCage(id int) *Cage {
	return &Cage{id: id, isOpen: false}
}

func (c *Cage) Open() {
	c.isOpen = true
}

func (c *Cage) Close() {
	c.isOpen = false
}

func (c *Cage) IsOpen() bool {
	return c.isOpen
}

// Environment contient tous les éléments du zoo pour la simulation
type Environment struct {
	cages    map[int]*Cage
	visitors []*Visitor
	animals  []*Animal
	//keepers  []*agents.Keeper
	escaped []*Animal
}

// NewEnvironment initialise un nouvel environnement de zoo
func NewEnvironment() *Environment {
	return &Environment{
		cages:    make(map[int]*Cage),
		visitors: []*Visitor{},
		animals:  []*Animal{},
	}
}

// AddCage ajoute une nouvelle cage dans l'environnement
func (env *Environment) AddCage(cage *Cage) {
	env.cages[cage.id] = cage
}

func (env *Environment) AddAnimal(animal *Animal) {
	env.animals = append(env.animals, animal)
}

// OpenCage ouvre une cage et libère l'animal s'il y en a un
func (env *Environment) OpenCage(id int) {
	cage, exists := env.cages[id]
	if exists && !cage.isOpen {
		cage.isOpen = true
		fmt.Printf("La cage %d a été ouverte par un animal !\n", id)
	}
}

// AddEscapedAnimal ajoute un animal à la liste des animaux échappés
func (env *Environment) AddEscapedAnimal(animal *Animal) {
	env.escaped = append(env.escaped, animal)
}

// ResetPanic réinitialise la panique en vidant la liste des animaux échappés
func (env *Environment) ResetPanic() {
	env.escaped = []*Animal{}
	fmt.Println("Tous les animaux échappés ont été capturés. La panique est terminée.")
}

// AddVisitor et AddKeeper permettent d'ajouter des agents dans l'environnement
func (env *Environment) AddVisitor(visitor *Visitor) {
	env.visitors = append(env.visitors, visitor)
	log.Println(env.visitors)
}

//func (env *Environment) AddKeeper(keeper *agents.Keeper) {
//	env.keepers = append(env.keepers, keeper)
//}

// GetEscapedAnimals retourne la liste des animaux échappés
func (env *Environment) GetEscapedAnimals() []*Animal {
	return env.escaped
}
func (env *Environment) GetAnimals() []*Animal {
	return env.animals
}
func (env *Environment) GetVisitors() []*Visitor {
	return env.visitors
}
