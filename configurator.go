package configurator

import (
	"fmt"
	"strconv"

	configuration1 "github.com/gossie/configurator/configuration"
	"github.com/gossie/configurator/internal/configuration"
	"github.com/gossie/configurator/internal/value"
)

func addConstraintsForSetValueIfFinal(cModel constraintModel, parameters map[int]*configuration.InternalParameter) {
	srcParam := parameters[cModel.srcId]
	forwardConstraint := configuration.CreateContraint(configuration.NewFinalCondition(cModel.srcId), configuration.NewSetValueExecution(cModel.srcId, cModel.targetId, cModel.targetValue.toInstance()))
	srcParam.AppendConstraint(forwardConstraint)

	targetParam := parameters[cModel.targetId]
	condition := configuration.NewCompositeCondition(
		configuration.NewValueCondition(cModel.targetId, configuration.IsImpossible, cModel.targetValue.toInstance()),
		configuration.Or,
		configuration.NewDisabledCondition(cModel.targetId),
	)
	backwardConstraint := configuration.CreateContraint(condition, configuration.NewDisableExecution(cModel.srcId))
	targetParam.AppendConstraint(backwardConstraint)
}

func addConstraintsForSetValueIfValue(cModel constraintModel, parameters map[int]*configuration.InternalParameter) {
	srcParam := parameters[cModel.srcId]
	newSrcConstraint := configuration.CreateContraint(configuration.NewValueCondition(cModel.srcId, configuration.Is, cModel.srcValue.toInstance()), configuration.NewSetValueExecution(cModel.srcId, cModel.targetId, cModel.targetValue.toInstance()))
	srcParam.AppendConstraint(newSrcConstraint)

	// TODO: exclude src value
	// targetParam := parameters[cModel.targetId]
	// newTargetConstraint := configuration.CreateContraint(configuration.NewValueCondition(cModel.targetId, configuration.IsNot, cModel.targetValue.toInstance()), configuration.NewDisableExecution(cModel.srcId))
	// targetParam.AppendConstraint(newTargetConstraint)
}

func addConstraintsForExcludeValueIfValue(cModel constraintModel, parameters map[int]*configuration.InternalParameter) {
	srcParam := parameters[cModel.srcId]
	newSrcConstraint := configuration.CreateContraint(configuration.NewValueCondition(cModel.srcId, configuration.Is, cModel.srcValue.toInstance()), configuration.NewExcludeValueExecution(cModel.srcId, cModel.targetId, cModel.targetValue.toInstance()))
	srcParam.AppendConstraint(newSrcConstraint)

	// TODO
}

func Start(model Model) configuration1.Configuration {
	parameters := make(map[int]*configuration.InternalParameter, len(model.parameters))
	for _, pModel := range model.parameters {
		pInstance := pModel.toInstance()
		parameters[pModel.id] = &pInstance
	}

	for _, cModel := range model.constraints {
		switch cModel.constraintType {
		default:
			panic(fmt.Sprintf("unknown constraint type %v", cModel.constraintType))
		case setValueIfFinal:
			addConstraintsForSetValueIfFinal(cModel, parameters)
		case setValueIfValue:
			addConstraintsForSetValueIfValue(cModel, parameters)
		case excludeValueIfValue:
			addConstraintsForExcludeValueIfValue(cModel, parameters)
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
