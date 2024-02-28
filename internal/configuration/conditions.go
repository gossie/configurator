package configuration

import (
	"fmt"

	"github.com/gossie/configurator/internal/value"
)

type compareOperator func(v1, v2 value.Value) bool

func Is(v1, v2 value.Value) bool {
	return v1.Subsumes(v2) && v2.Subsumes(v1)
}

func IsNot(v1, v2 value.Value) bool {
	return !v1.Subsumes(v2) || !v2.Subsumes(v1)
}

type logicalOperator int

const (
	and logicalOperator = iota
	or
)

type condition interface {
	fulfilled(map[int]*InternalParameter) bool
}

type compositeCondition struct {
	left, right condition
	operator    logicalOperator
}

func NewCompositeCondition(left condition, op logicalOperator, right condition) compositeCondition {
	return compositeCondition{left, right, op}
}

func (composite compositeCondition) fulfilled(config map[int]*InternalParameter) bool {
	switch composite.operator {
	default:
		panic(fmt.Sprintf("unknown operator: %v", composite.operator))
	case and:
		return composite.left.fulfilled(config) && composite.right.fulfilled(config)
	case or:
		return composite.left.fulfilled(config) || composite.right.fulfilled(config)
	}
}

type parameterFinalCondition struct {
	parameterId int
}

func NewFinalCondition(paramId int) parameterFinalCondition {
	return parameterFinalCondition{paramId}
}

func (condition parameterFinalCondition) fulfilled(config map[int]*InternalParameter) bool {
	param := config[condition.parameterId]
	return param.Final()
}

type parameterValueCondition struct {
	parameterId int
	operator    compareOperator
	value       value.Value
}

func NewValueCondition(paramId int, op compareOperator, value value.Value) parameterValueCondition {
	return parameterValueCondition{paramId, op, value}
}

func (condition parameterValueCondition) fulfilled(config map[int]*InternalParameter) bool {
	param := config[condition.parameterId]
	return condition.operator(param.RestrictedValue(), condition.value)
}
