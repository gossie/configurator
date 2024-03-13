package configurator

import (
	"log"

	"github.com/gossie/configurator/internal/configuration"
	"github.com/gossie/configurator/internal/value"
)

type ValueType int

const (
	IntSetType ValueType = iota
	IntRangeType
	FinalInt
	StringSetType
)

type valueModel struct {
	valueType        ValueType
	intValues        []int
	stringValues     []string
	min, max         int
	minOpen, maxOpen bool
	finalValue       int
}

func NewIntSetModel(values []int) valueModel {
	return valueModel{
		valueType: IntSetType,
		intValues: values,
	}
}

func NewIntRangeModel(min int, minOpen bool, max int, maxOpen bool) valueModel {
	return valueModel{
		valueType: IntRangeType,
		min:       min,
		minOpen:   minOpen,
		max:       max,
		maxOpen:   maxOpen,
	}
}

func NewFinalIntModel(value int) valueModel {
	return valueModel{
		valueType:  FinalInt,
		finalValue: value,
	}
}

func (vModel valueModel) toInstance() value.Value {
	switch vModel.valueType {
	case IntSetType:
		return value.NewIntValues(vModel.intValues)
	case StringSetType:
		confValues := make([]int, len(vModel.stringValues))
		for index := range len(vModel.stringValues) {
			confValues[index] = index
		}
		return value.NewIntValues(confValues)
	case IntRangeType:
		return value.NewIntRange(vModel.min, vModel.minOpen, vModel.max, vModel.maxOpen)
	case FinalInt:
		return value.NewIntValues([]int{vModel.finalValue})
	default:
		log.Default().Println("unknown value type", vModel.valueType)
		return nil
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

type ConstraintType int

const (
	SetValueIfFinal ConstraintType = iota
	SetValueIfValue
	ExcludeValueIfValue
)

type constraintModel struct {
	constraintType        ConstraintType
	srcId, targetId       int
	srcValue, targetValue valueModel
}

func NewSetValueIfFinalConstraintModel(srcId, targetId int, targetValue valueModel) constraintModel {
	return constraintModel{
		constraintType: SetValueIfFinal,
		srcId:          srcId,
		targetId:       targetId,
		targetValue:    targetValue,
	}
}

func NewSetValueIfValueConstraintModel(srcId int, srcValue valueModel, targetId int, targetValue valueModel) constraintModel {
	return constraintModel{
		constraintType: SetValueIfValue,
		srcId:          srcId,
		targetId:       targetId,
		srcValue:       srcValue,
		targetValue:    targetValue,
	}
}

func NewExcludeValueIfValueConstraintModel(srcId int, srcValue valueModel, targetId int, targetValue valueModel) constraintModel {
	return constraintModel{
		constraintType: ExcludeValueIfValue,
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
