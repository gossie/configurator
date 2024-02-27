package configurator

import "strconv"

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

func NewFinalInt(value int) valueModel {
	return valueModel{
		valueType:  finalInt,
		finalValue: value,
	}
}

func (vModel valueModel) toInstance() value {
	switch vModel.valueType {
	default:
		panic("unknown value type " + strconv.Itoa(int(vModel.valueType)))
	case intSetType:
		return intValues{vModel.values}
	case intRangeType:
		return intRange{
			min:     vModel.min,
			minOpen: vModel.minOpen,
			max:     vModel.max,
			maxOpen: vModel.maxOpen,
		}
	case finalInt:
		return intValues{[]int{vModel.finalValue}}
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

func (pModel parameterModel) toInstance(model Model) Parameter {
	constraints := make([]constraint, 0)
	for _, cModel := range model.constraints {
		if cModel.srcId == pModel.id {
			newConstraint := createContraint(newFinalCondition(cModel.srcId), newSetValueExecution(cModel.targetId, cModel.targetValue.toInstance()))
			constraints = append(constraints, newConstraint)
		}
	}
	return Parameter{
		id:          pModel.id,
		name:        pModel.name,
		value:       pModel.value.toInstance(),
		constraints: constraints,
	}
}

type constraintModel struct {
	srcId, targetId int
	targetValue     valueModel
}

func NewSetValueIfFinalConstraintModel(srcId, targetId int, targetValue valueModel) constraintModel {
	return constraintModel{
		srcId:       srcId,
		targetId:    targetId,
		targetValue: targetValue,
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
