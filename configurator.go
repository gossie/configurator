package configurator

import (
	"fmt"
	"strconv"
)

type Configuration struct {
	parameters []*Parameter
}

func (config Configuration) ParameterById(id int) (Parameter, error) {
	pointer, err := config.mutableParameterById(id)
	if err != nil {
		var empty Parameter
		return empty, fmt.Errorf("parameter with ID %v could not be found", id)
	}

	return *pointer, nil
}

func (config Configuration) mutableParameterById(id int) (*Parameter, error) {
	for _, p := range config.parameters {
		if p.id == id {
			return p, nil
		}
	}

	return nil, fmt.Errorf("parameter with ID %v could not be found", id)
}

func Start(model Model) Configuration {
	parameters := make([]*Parameter, 0, len(model.parameters))
	for _, pModel := range model.parameters {
		pInstance := pModel.toInstance(model)
		parameters = append(parameters, &pInstance)
	}

	return Configuration{
		parameters: parameters,
	}
}

func SetValue(configuration Configuration, parameterId int, value string) (Configuration, error) {
	parameter, err := configuration.mutableParameterById(parameterId)
	if err != nil {
		return configuration, err
	}

	intValue, _ := strconv.Atoi(value)
	err = parameter.SetValue(intValues{[]int{intValue}})
	for _, c := range parameter.constraints {
		c(configuration)
	}
	return configuration, err
}
