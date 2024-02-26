package configurator

import "strconv"

type valueType int

const (
	intSetType valueType = iota
)

type valueModel struct {
	valueType valueType
	values    []int
}

func NewIntSetModel(values []int) valueModel {
	return valueModel{
		valueType: intSetType,
		values:    values,
	}
}

func (vModel valueModel) toInstance() value {
	switch vModel.valueType {
	default:
		panic("unknown value type " + strconv.Itoa(int(vModel.valueType)))
	case intSetType:
		return intValues{vModel.values}
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
