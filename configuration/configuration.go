package configuration

type Configuration interface {
	ParameterById(id int) (Parameter, error)
}

type Parameter interface {
	Id() int
	Name() string
	Selectable() bool
	Final() bool
	Value() string
}
