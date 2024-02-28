package configurator

import "fmt"

// type operator int

// const (
// 	is operator = iota
// 	inNot
// )

type operator int

const (
	and operator = iota
	or
)

type condition interface {
	fulfilled(configuration Configuration) bool
}

type compositeCondition struct {
	left, right condition
	operator    operator
}

func newCompositeCondition(left condition, op operator, right condition) compositeCondition {
	return compositeCondition{left, right, op}
}

func (composite compositeCondition) fulfilled(configuration Configuration) bool {
	switch composite.operator {
	default:
		panic(fmt.Sprintf("unknown operator: %v", composite.operator))
	case and:
		return composite.left.fulfilled(configuration) && composite.right.fulfilled(configuration)
	case or:
		return composite.left.fulfilled(configuration) || composite.right.fulfilled(configuration)
	}
}

type parameterFinalCondition struct {
	parameterId int
}

func newFinalCondition(paramId int) parameterFinalCondition {
	return parameterFinalCondition{paramId}
}

func (condition parameterFinalCondition) fulfilled(configuration Configuration) bool {
	param, _ := configuration.ParameterById(condition.parameterId)
	return param.Final()
}

type execution interface {
	execute(configuration Configuration)
}

type setValueExecution struct {
	parameterId int
	value       value
}

func newSetValueExecution(paramId int, value value) setValueExecution {
	return setValueExecution{paramId, value}
}

func (execution setValueExecution) execute(configuration Configuration) {
	param, _ := configuration.mutableParameterById(execution.parameterId)
	change, _ := param.SetValue(execution.value)
	if change {
		for _, c := range param.constraints {
			c(configuration)
		}
	}
}

type constraint func(configuration Configuration) (bool, error)

func createContraint(condition condition, exexution execution) constraint {
	return func(configuration Configuration) (bool, error) {
		if condition.fulfilled(configuration) {
			exexution.execute(configuration)
		}
		return false, nil
	}
}
