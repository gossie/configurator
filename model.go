package configurator

import "strconv"

type valueType int

const (
	intSetType valueType = iota
)

type ValueModel struct {
	valueType valueType
	values    []int
}

func NewIntSetModel(values []int) ValueModel {
	return ValueModel{
		valueType: intSetType,
		values:    values,
	}
}

func (vModel ValueModel) toInstance() Value {
	switch vModel.valueType {
	default:
		panic("unknown value type " + strconv.Itoa(int(vModel.valueType)))
	case intSetType:
		return intSet{vModel.values}
	}
}

type ParameterModel struct {
	id    int
	name  string
	value ValueModel
}

func NewParameterModel(name string, value ValueModel) ParameterModel {
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
