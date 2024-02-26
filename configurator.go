package configurator

import (
	"fmt"
)

type Parameter struct {
	id   int
	name string
}

func (p Parameter) Id() int {
	return p.id
}

func (p Parameter) Name() string {
	return p.name
}

type Configuration struct {
	parameters []*Parameter
}

func (config Configuration) ParameterById(id int) (Parameter, error) {
	for _, p := range config.parameters {
		if p.id == id {
			return *p, nil
		}
	}

	var empty Parameter
	return empty, fmt.Errorf("parameter with ID %v could not be found", id)
}

func Start(model Model) Configuration {
	parameters := make([]*Parameter, 0, len(model.parameters))
	for _, pModel := range model.parameters {
		pInstance := pModel.toInstance()
		parameters = append(parameters, &pInstance)
	}

	return Configuration{
		parameters: parameters,
	}
}
