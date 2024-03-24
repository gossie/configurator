package configurator

import (
	"fmt"
	"log"
	"strconv"

	model "github.com/gossie/configuration-model"
	configuration1 "github.com/gossie/configurator/configuration"
	"github.com/gossie/configurator/internal/configuration"
	"github.com/gossie/configurator/internal/value"
)

func addConstraintsForSetValueIfFinal(cModel model.ConstraintModel, parameters map[int]*configuration.InternalParameter) {
	srcParam := parameters[cModel.SrcId()]
	forwardConstraint := configuration.CreateContraint(configuration.NewFinalCondition(cModel.SrcId()), configuration.NewSetValueExecution(cModel.SrcId(), cModel.TargetId(), valueToInstance(cModel.TargetValue())))
	srcParam.AppendConstraint(forwardConstraint)

	targetParam := parameters[cModel.TargetId()]
	condition := configuration.NewCompositeCondition(
		configuration.NewValueCondition(cModel.TargetId(), configuration.IsImpossible, valueToInstance(cModel.TargetValue())),
		configuration.Or,
		configuration.NewDisabledCondition(cModel.TargetId()),
	)
	backwardConstraint := configuration.CreateContraint(condition, configuration.NewDisableExecution(cModel.SrcId()))
	targetParam.AppendConstraint(backwardConstraint)
}

func addConstraintsForSetValueIfValue(cModel model.ConstraintModel, parameters map[int]*configuration.InternalParameter) {
	srcParam := parameters[cModel.SrcId()]
	newSrcConstraint := configuration.CreateContraint(configuration.NewValueCondition(cModel.SrcId(), configuration.Is, valueToInstance(cModel.SrcValue())), configuration.NewSetValueExecution(cModel.SrcId(), cModel.TargetId(), valueToInstance(cModel.TargetValue())))
	srcParam.AppendConstraint(newSrcConstraint)

	// TODO: exclude src value
	// targetParam := parameters[cModel.targetId]
	// newTargetConstraint := configuration.CreateContraint(configuration.NewValueCondition(cModel.targetId, configuration.IsNot, cModel.targetValue.toInstance()), configuration.NewDisableExecution(cModel.srcId))
	// targetParam.AppendConstraint(newTargetConstraint)
}

func addConstraintsForExcludeValueIfValue(cModel model.ConstraintModel, parameters map[int]*configuration.InternalParameter) {
	srcParam := parameters[cModel.SrcId()]
	newSrcConstraint := configuration.CreateContraint(configuration.NewValueCondition(cModel.SrcId(), configuration.Is, valueToInstance(cModel.SrcValue())), configuration.NewExcludeValueExecution(cModel.SrcId(), cModel.TargetId(), valueToInstance(cModel.TargetValue())))
	srcParam.AppendConstraint(newSrcConstraint)

	// TODO
}

func Start(confModel model.Model) configuration1.Configuration {
	parameters := make(map[int]*configuration.InternalParameter, len(confModel.Parameters()))
	for _, pModel := range confModel.Parameters() {
		pInstance := parameterToInstance(pModel)
		parameters[pModel.Id()] = &pInstance
	}

	for _, cModel := range confModel.Constraints() {
		switch cModel.ConstraintType() {
		case model.SetValueIfFinal:
			addConstraintsForSetValueIfFinal(cModel, parameters)
		case model.SetValueIfValue:
			addConstraintsForSetValueIfValue(cModel, parameters)
		case model.ExcludeValueIfValue:
			addConstraintsForExcludeValueIfValue(cModel, parameters)
		default:
			log.Default().Println("unknown constraint type", cModel.ConstraintType())
		}
	}

	config := configuration.NewInternalConfiguration(parameters)
	for _, p := range config.Parameters {
		for _, c := range p.Constraints {
			c(config.Parameters)
		}
	}

	return config
}

func SetValue(config configuration1.Configuration, parameterId int, valueToSet string) (configuration1.Configuration, error) {
	internalConfig := config.(configuration.InternalConfiguration)
	parameter, ok := internalConfig.Parameters[parameterId]
	if !ok {
		return config, fmt.Errorf("parameter with ID %v could not be found", parameterId)
	}

	intValue, _ := strconv.Atoi(valueToSet)
	_, err := parameter.RestrictValue(value.NewIntValues([]int{intValue}))
	for _, c := range parameter.Constraints {
		c(internalConfig.Parameters)
	}
	return config, err
}

func parameterToInstance(pModel model.ParameterModel) configuration.InternalParameter {
	return configuration.NewInternalParameter(
		pModel.Id(),
		pModel.Name(),
		valueToInstance(pModel.Value()),
	)
}

func valueToInstance(vModel model.ValueModel) value.Value {
	switch vModel.ValueType() {
	case model.IntSetType:
		return value.NewIntValues(vModel.IntValues())
	case model.StringSetType:
		confValues := make([]int, len(vModel.StringValues()))
		for index := range len(vModel.StringValues()) {
			confValues[index] = index
		}
		return value.NewIntValues(confValues)
	case model.IntRangeType:
		return value.NewIntRange(vModel.Min(), vModel.MinOpen(), vModel.Max(), vModel.MaxOpen())
	case model.FinalInt:
		return value.NewIntValues([]int{vModel.FinalValue()})
	default:
		log.Default().Println("unknown value type", vModel.ValueType())
		return nil
	}
}
