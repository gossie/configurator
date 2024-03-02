package configurator

import (
	"strconv"

	"github.com/gossie/configurator/internal/configuration"
	"github.com/gossie/configurator/internal/value"
)

type valueType int

const (
	intSetType valueType = iota
	intRangeType
	finalInt
)

type valueModel struct {
	valueType        valueType
	values           []int
	min, max         int
	minOpen, maxOpen bool
	finalValue       int
}

func NewIntSetModel(values []int) valueModel {
	return valueModel{
		valueType: intSetType,
		values:    values,
	}
}

func NewIntRangeModel(min int, minOpen bool, max int, maxOpen bool) valueModel {
	return valueModel{
		valueType: intRangeType,
		min:       min,
		minOpen:   minOpen,
		max:       max,
		maxOpen:   maxOpen,
	}
}

func NewFinalIntModel(value int) valueModel {
	return valueModel{
		valueType:  finalInt,
		finalValue: value,
	}
}

func (vModel valueModel) toInstance() value.Value {
	switch vModel.valueType {
	case intSetType:
		return value.NewIntValues(vModel.values)
	case intRangeType:
		return value.NewIntRange(vModel.min, vModel.minOpen, vModel.max, vModel.maxOpen)
	case finalInt:
		return value.NewIntValues([]int{vModel.finalValue})
	default:
		panic("unknown value type " + strconv.Itoa(int(vModel.valueType)))
	}
}

type parameterModel struct {
	id    int
	name  string
	value valueModel
}

func (pModel parameterModel) Id() int {
	return pModel.id
}

func (pModel parameterModel) toInstance() configuration.InternalParameter {
	return configuration.NewInternalParameter(
		pModel.id,
		pModel.name,
		pModel.value.toInstance(),
	)
}

type constraintType int

const (
	setValueIfFinal constraintType = iota
	setValueIfValue
	excludeValueIfValue
)

type constraintModel struct {
	constraintType        constraintType
	srcId, targetId       int
	srcValue, targetValue valueModel
}

func NewSetValueIfFinalConstraintModel(srcId, targetId int, targetValue valueModel) constraintModel {
	return constraintModel{
		constraintType: setValueIfFinal,
		srcId:          srcId,
		targetId:       targetId,
		targetValue:    targetValue,
	}
}

func NewSetValueIfValueConstraintModel(srcId int, srcValue valueModel, targetId int, targetValue valueModel) constraintModel {
	return constraintModel{
		constraintType: setValueIfValue,
		srcId:          srcId,
		targetId:       targetId,
		srcValue:       srcValue,
		targetValue:    targetValue,
	}
}

func NewExcludeValueIfValueConstraintModel(srcId int, srcValue valueModel, targetId int, targetValue valueModel) constraintModel {
	return constraintModel{
		constraintType: excludeValueIfValue,
		srcId:          srcId,
		targetId:       targetId,
		srcValue:       srcValue,
		targetValue:    targetValue,
	}
}

type Model struct {
	nextParameterId int
	parameters      []parameterModel
	constraints     []constraintModel
}

func (pModel *Model) AddParameter(name string, value valueModel) parameterModel {
	pModel.nextParameterId++
	newParameter := parameterModel{
		id:    pModel.nextParameterId,
		name:  name,
		value: value,
	}
	pModel.parameters = append(pModel.parameters, newParameter)
	return newParameter
}

func (pModel *Model) AddConstraint(constraint constraintModel) {
	pModel.constraints = append(pModel.constraints, constraint)
}
