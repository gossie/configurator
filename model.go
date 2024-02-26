package configurator

var nextParameterId = 1

type ParameterModel struct {
	name string
}

func (pModel ParameterModel) toInstance() Parameter {
	parameterId := nextParameterId
	nextParameterId++
	return Parameter{
		id:   parameterId,
		name: pModel.name,
	}
}

type Model struct {
	parameters []ParameterModel
}

func (pModel *Model) AddParameter(pName string) {
	pModel.parameters = append(pModel.parameters, ParameterModel{pName})
}
