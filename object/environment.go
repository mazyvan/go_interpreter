package object

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) GetEnvFor(name string) (env *Environment) {
	_, ok := e.store[name]
	if ok {
		return e
	}
	if !ok && e.outer != nil {
		return e.outer.GetEnvFor(name)
	}
	return nil
}

func (e *Environment) Set(name string, val Object) Object {
	env := e.GetEnvFor(name)
	if env == nil {
		env = e
	}
	env.store[name] = val
	return val
}
