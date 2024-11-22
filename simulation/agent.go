package simulation

type Agent interface {
	Start()
	Percept(env *Environment)
	Deliberate()
	Act(env *Environment)
}
