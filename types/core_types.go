package types

type Closure struct {
	Args []interface{}
	Mutator func(mutator *Closure)
	ReturnValue interface{}
}
