package main

import (
	"fmt"
	"log"
	"zoo-panic/server"
	"zoo-panic/simulation"
)

// Fonction d'initialisation des éléments de la simulation
func setupSimulation() *simulation.SimulationMA {
	// Création de l'environnement
	env := simulation.NewEnvironment()

	// Création de la simulation
	simu := simulation.NewSimulation(env)

	cage1 := simulation.NewCage(1)
	cage2 := simulation.NewCage(2)
	cage3 := simulation.NewCage(3)
	env.AddCage(cage1)
	env.AddCage(cage2)
	env.AddCage(cage3)

	// Ajout d'animaux avec différents comportements
	lion := simulation.NewAnimal(1, "Lion", simulation.Aggressive, cage1)
	lion2 := simulation.NewAnimal(11, "Lion", simulation.Aggressive, cage1)
	lion3 := simulation.NewAnimal(12, "Lion", simulation.Aggressive, cage1)
	lion6 := simulation.NewAnimal(13, "Lion", simulation.Aggressive, cage1)
	lion5 := simulation.NewAnimal(14, "Lion", simulation.Aggressive, cage1)
	lion4 := simulation.NewAnimal(15, "Lion", simulation.Aggressive, cage1)
	monkey := simulation.NewAnimal(2, "Monkey", simulation.Clever, cage2)
	deer := simulation.NewAnimal(3, "Deer", simulation.Passive, cage3)
	deer3 := simulation.NewAnimal(31, "Deer", simulation.Passive, cage3)
	deer2 := simulation.NewAnimal(32, "Deer", simulation.Passive, cage3)
	//cage2.Open()
	// Ajout des animaux à l'environnement et à la simulation
	env.AddAnimal(lion)
	env.AddAnimal(lion3)
	env.AddAnimal(lion2)
	env.AddAnimal(monkey)
	env.AddAnimal(lion4)
	env.AddAnimal(lion5)
	env.AddAnimal(lion6)
	env.AddAnimal(deer)
	env.AddAnimal(deer2)
	env.AddAnimal(deer3)
	simu.AddAgent(lion)
	simu.AddAgent(lion2)
	simu.AddAgent(lion3)
	simu.AddAgent(monkey)
	simu.AddAgent(lion4)
	simu.AddAgent(lion5)
	simu.AddAgent(lion6)
	simu.AddAgent(deer)
	simu.AddAgent(deer2)
	simu.AddAgent(deer3)

	// Ajout des visiteurs
	visitor1 := simulation.NewVisitor(1)
	visitor2 := simulation.NewVisitor(2)
	log.Println("here")
	env.AddVisitor(visitor1)
	env.AddVisitor(visitor2)
	simu.AddAgent(visitor1)
	simu.AddAgent(visitor2)

	return simu
}

func main() {
	// Initialisation de la simulation
	simu := setupSimulation()
	restServer := server.NewServerRestAgent(":8080", simu)
	go restServer.Start()

	// Lancement de la boucle de simulation
	fmt.Println("Lancement de la simulation : Panique au Zoo")
	go simu.Run()

	// Garde le programme en exécution
	select {}
}
