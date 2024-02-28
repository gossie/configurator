package configuration

import (
	"github.com/gossie/configurator/internal/value"
)

type execution interface {
	execute(map[int]*InternalParameter)
	revert(map[int]*InternalParameter)
}

type setValueExecution struct {
	srcId, targetId int
	value           value.Value
}

func NewSetValueExecution(srcId, targetId int, value value.Value) setValueExecution {
	return setValueExecution{srcId, targetId, value}
}

func (execution setValueExecution) execute(config map[int]*InternalParameter) {
	param := config[execution.targetId]
	change, _ := param.RestrictValueFromConstraint(execution.srcId, execution.value)
	if change {
		for _, c := range param.Constraints {
			c(config)
		}
	}
}

func (execution setValueExecution) revert(config map[int]*InternalParameter) {
	param := config[execution.targetId]
	param.RemoveRestrictionsFromConstraint(execution.srcId)
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

func (execution disableExecution) revert(config map[int]*InternalParameter) {
	// TODO
}
