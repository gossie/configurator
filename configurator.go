package configurator

import (
	"fmt"
	"strconv"

	configuration1 "github.com/gossie/configurator/configuration"
	"github.com/gossie/configurator/internal/configuration"
	"github.com/gossie/configurator/internal/value"
)

func Start(model Model) configuration1.Configuration {
	parameters := make(map[int]*configuration.InternalParameter, len(model.parameters))
	for _, pModel := range model.parameters {
		pInstance := pModel.toInstance()
		parameters[pModel.id] = &pInstance
	}

	for _, cModel := range model.constraints {
		srcParam := parameters[cModel.srcId]
		newSrcConstraint := configuration.CreateContraint(configuration.NewFinalCondition(cModel.srcId), configuration.NewSetValueExecution(cModel.targetId, cModel.targetValue.toInstance()))
		srcParam.AppendConstraint(newSrcConstraint)

		targetParam := parameters[cModel.targetId]
		newTargetConstraint := configuration.CreateContraint(configuration.NewValueCondition(cModel.targetId, configuration.IsNot, cModel.targetValue.toInstance()), configuration.NewDisableExecution(cModel.srcId))
		targetParam.AppendConstraint(newTargetConstraint)

	}

	return configuration.NewInternalConfiguration(parameters)
}

func SetValue(config configuration1.Configuration, parameterId int, valueToSet string) (configuration1.Configuration, error) {
	internalConfig := config.(configuration.InternalConfiguration)
	parameter, ok := internalConfig.Parameters[parameterId]
	if !ok {
		return config, fmt.Errorf("parameter with ID %v could not be found", parameterId)
	}

	intValue, _ := strconv.Atoi(valueToSet)
	_, err := parameter.SetValue(value.NewIntValues([]int{intValue}))
	for _, c := range parameter.Constraints {
		c(internalConfig.Parameters)
	}
	return config, err
}
