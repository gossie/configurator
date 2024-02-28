package configuration

import (
	"github.com/gossie/configurator/internal/value"
)

type execution interface {
	execute(map[int]*InternalParameter)
}

type setValueExecution struct {
	parameterId int
	value       value.Value
}

func NewSetValueExecution(paramId int, value value.Value) setValueExecution {
	return setValueExecution{paramId, value}
}

func (execution setValueExecution) execute(config map[int]*InternalParameter) {
	param := config[execution.parameterId]
	change, _ := param.SetValue(execution.value)
	if change {
		for _, c := range param.Constraints {
			c(config)
		}
	}
}

type disableExecution struct {
	parameterId int
}

func NewDisableExecution(paramId int) disableExecution {
	return disableExecution{paramId}
}

func (execution disableExecution) execute(config map[int]*InternalParameter) {
	param := config[execution.parameterId]
	param.selectable = false
	for _, c := range param.Constraints {
		c(config)
	}
}
