package configuration

import (
	"fmt"

	"github.com/gossie/configurator/configuration"
)

type InternalConfiguration struct {
	Parameters map[int]*InternalParameter
}

func NewInternalConfiguration(parameters map[int]*InternalParameter) InternalConfiguration {
	return InternalConfiguration{
		Parameters: parameters,
	}
}

func (config InternalConfiguration) ParameterById(id int) (configuration.Parameter, error) {
	pointer, err := config.MutableParameterById(id)
	if err != nil {
		var empty InternalParameter
		return empty, fmt.Errorf("parameter with ID %v could not be found", id)
	}

	return *pointer, nil
}

func (config InternalConfiguration) MutableParameterById(id int) (*InternalParameter, error) {
	if param, ok := config.Parameters[id]; ok {
		return param, nil
	}
	return nil, fmt.Errorf("parameter with ID %v could not be found", id)
}
