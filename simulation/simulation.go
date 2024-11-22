package simulation

import (
	"log"
	"time"
)

type SimulationMA struct {
	Env    *Environment
	agents []Agent
}

func NewSimulation(envt *Environment) *SimulationMA {
	agents := []Agent{}
	return &SimulationMA{
		Env:    envt,
		agents: agents,
	}
}

// AddAgent ajoute un agent à la simulation
func (simu *SimulationMA) AddAgent(agent Agent) {
	simu.agents = append(simu.agents, agent)
}

func (simu *SimulationMA) Run() {
	for step := 0; step < 30; step++ {

		//log.Println("//////////////////////////////////", step)
		//log.Println("//////////////////////////////////")
		//log.Println("//////////////////////////////////")
		for _, agent := range simu.agents {
			log.Println(agent)
			agent.Percept(simu.Env)
			agent.Deliberate()
			agent.Act(simu.Env)
		}
		time.Sleep(5000 * time.Millisecond)

	}
}
