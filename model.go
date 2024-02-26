package configurator

import "strconv"

type valueType int

const (
	intSetType valueType = iota
	intRangeType
)

type valueModel struct {
	valueType        valueType
	values           []int
	min, max         int
	minOpen, maxOpen bool
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
	}
}

type ParameterModel struct {
	id    int
	name  string
	value valueModel
}

func NewParameterModel(name string, value valueModel) ParameterModel {
	return ParameterModel{
		name:  name,
		value: value,
	}
}

func (pModel ParameterModel) toInstance() Parameter {
	return Parameter{
		id:    pModel.id,
		name:  pModel.name,
		value: pModel.value.toInstance(),
	}
}

type Model struct {
	nextParameterId int
	parameters      []ParameterModel
}

func (pModel *Model) AddParameter(param ParameterModel) {
	pModel.nextParameterId++
	param.id = pModel.nextParameterId
	pModel.parameters = append(pModel.parameters, param)
}
