package configurator_test

import (
	"testing"

	model "github.com/gossie/configuration-model"
	"github.com/gossie/configurator"
	"github.com/gossie/configurator/configuration"
)

func checkOpenParameter(param configuration.Parameter, err error, expectedId int, expectedName string, expectedValue string, t *testing.T) {
	if err != nil {
		t.Fatalf("parameter with ID %v could not be found: %v", expectedId, err.Error())
	}
	if param.Id() != expectedId {
		t.Fatalf("ID of the parameter should be %v but was %v", expectedId, param.Id())
	}
	if param.Name() != expectedName {
		t.Fatalf("name of parameter with ID %v should be P3 but was %v", expectedId, param.Name())
	}
	if !param.Selectable() {
		t.Fatalf("parameter with ID %v should be selectable", expectedId)
	}
	if param.Final() {
		t.Fatalf("value %v of parameter with ID %v should not be final", expectedId, param.Value())
	}
	if param.Value() != expectedValue {
		t.Fatalf("value of parameter with ID %v should be %v but was %v", expectedId, expectedValue, param.Value())
	}
}

func checkFinalParameter(param configuration.Parameter, err error, expectedId int, expectedName string, expectedValue string, t *testing.T) {
	if err != nil {
		t.Fatalf("parameter with ID %v could not be found: %v", expectedId, err.Error())
	}
	if param.Id() != expectedId {
		t.Fatalf("ID of the parameter should be %v but was %v", expectedId, param.Id())
	}
	if param.Name() != expectedName {
		t.Fatalf("name of parameter with ID %v should be P3 but was %v", expectedId, param.Name())
	}
	if !param.Selectable() {
		t.Fatalf("parameter with ID %v should be selectable", expectedId)
	}
	if !param.Final() {
		t.Fatalf("value %v of parameter with ID %v should be final", param.Value(), expectedId)
	}
	if param.Value() != expectedValue {
		t.Fatalf("value of parameter with ID %v should be %v but was %v", expectedId, expectedValue, param.Value())
	}
}

func checkUnselectableParameter(param configuration.Parameter, err error, expectedId int, expectedName string, t *testing.T) {
	if err != nil {
		t.Fatalf("parameter with ID %v could not be found: %v", expectedId, err.Error())
	}
	if param.Id() != expectedId {
		t.Fatalf("ID of the parameter should be %v but was %v", expectedId, param.Id())
	}
	if param.Name() != expectedName {
		t.Fatalf("name of parameter with ID %v should be P3 but was %v", expectedId, param.Name())
	}
	if param.Selectable() {
		t.Fatalf("parameter with ID %v should not be selectable", expectedId)
	}
}

func TestThatConfigurationIsStarted(t *testing.T) {
	confModel := model.Model{}
	confModel.AddParameter("P1", model.NewIntSetModel([]int{1, 2, 3}))
	confModel.AddParameter("P2", model.NewIntSetModel([]int{1, 2, 3}))
	confModel.AddParameter("P3", model.NewIntRangeModel(7, false, 17, false))
	confModel.AddParameter("P4", model.NewIntRangeModel(7, false, 17, true))
	confModel.AddParameter("P5", model.NewIntRangeModel(7, true, 17, false))
	confModel.AddParameter("P6", model.NewIntRangeModel(7, true, 17, true))

	configuration := configurator.Start(confModel)

	p1, errP1 := configuration.ParameterById(1)
	checkOpenParameter(p1, errP1, 1, "P1", "{1,2,3}", t)

	p2, errP2 := configuration.ParameterById(2)
	checkOpenParameter(p2, errP2, 2, "P2", "{1,2,3}", t)

	p3, errP3 := configuration.ParameterById(3)
	checkOpenParameter(p3, errP3, 3, "P3", "[7;17]", t)

	p4, errP4 := configuration.ParameterById(4)
	checkOpenParameter(p4, errP4, 4, "P4", "[7;17[", t)

	p5, errP5 := configuration.ParameterById(5)
	checkOpenParameter(p5, errP5, 5, "P5", "]7;17]", t)

	p6, errP6 := configuration.ParameterById(6)
	checkOpenParameter(p6, errP6, 6, "P6", "]7;17[", t)
}

func TestThatParameterIsNotFound(t *testing.T) {
	confModel := model.Model{}
	confModel.AddParameter("P1", model.NewIntSetModel([]int{1, 2, 3}))

	configuration := configurator.Start(confModel)

	_, err := configuration.ParameterById(2)
	if err == nil {
		t.Fatal("parameter with ID 2 should not be found")
	}
}

func TestThatValueIsNotSetBecauseTheParameterIsMissing(t *testing.T) {
	confModel := model.Model{}
	confModel.AddParameter("P1", model.NewIntSetModel([]int{1, 2, 3}))

	configuration := configurator.Start(confModel)
	configuration, _ = configurator.SetValue(configuration, 2, "2")

	_, err := configuration.ParameterById(2)

	if err == nil {
		t.Fatal("parameter with ID 2 should not be found")
	}
}

func TestThatValueIsSet_intValues(t *testing.T) {
	confModel := model.Model{}
	confModel.AddParameter("P1", model.NewIntSetModel([]int{1, 2, 3}))
	confModel.AddParameter("P2", model.NewIntSetModel([]int{1, 2, 3}))

	configuration := configurator.Start(confModel)
	configuration, _ = configurator.SetValue(configuration, 1, "2")

	p1, errP1 := configuration.ParameterById(1)
	p2, errP2 := configuration.ParameterById(2)

	checkFinalParameter(p1, errP1, 1, "P1", "2", t)
	checkOpenParameter(p2, errP2, 2, "P2", "{1,2,3}", t)
}

func TestThatAnImpossibleValueIsNotSet_intValues(t *testing.T) {
	confModel := model.Model{}
	confModel.AddParameter("P1", model.NewIntSetModel([]int{1, 2, 3}))

	configuration := configurator.Start(confModel)
	configuration, err := configurator.SetValue(configuration, 1, "4")
	if err == nil {
		t.Fatal("4 is not a valid value and should cause an error")
	}

	p1, errP1 := configuration.ParameterById(1)

	checkOpenParameter(p1, errP1, 1, "P1", "{1,2,3}", t)
}

func TestThatValueIsSet_intRange(t *testing.T) {
	confModel := model.Model{}
	confModel.AddParameter("P1", model.NewIntRangeModel(1, false, 8, false))
	confModel.AddParameter("P2", model.NewIntSetModel([]int{1, 2, 3}))

	configuration := configurator.Start(confModel)
	configuration, _ = configurator.SetValue(configuration, 1, "2")

	p1, errP1 := configuration.ParameterById(1)
	p2, errP2 := configuration.ParameterById(2)

	checkFinalParameter(p1, errP1, 1, "P1", "2", t)
	checkOpenParameter(p2, errP2, 2, "P2", "{1,2,3}", t)
}

func TestThatAnImpossibleValueIsNotSet_intRange(t *testing.T) {
	confModel := model.Model{}
	confModel.AddParameter("P1", model.NewIntRangeModel(1, false, 8, false))

	configuration := configurator.Start(confModel)
	configuration, err := configurator.SetValue(configuration, 1, "9")
	if err == nil {
		t.Fatal("4 is not a valid value and should cause an error")
	}

	p1, errP1 := configuration.ParameterById(1)

	checkOpenParameter(p1, errP1, 1, "P1", "[1;8]", t)
}

func TestThatForwardRuleIsExecuted(t *testing.T) {
	confModel := model.Model{}
	pModel1 := confModel.AddParameter("P1", model.NewIntRangeModel(1, false, 8, false))
	pModel2 := confModel.AddParameter("P2", model.NewIntSetModel([]int{1, 2, 3}))
	confModel.AddConstraint(model.NewSetValueIfFinalConstraintModel(pModel1.Id(), pModel2.Id(), model.NewFinalIntModel(3)))

	configuration := configurator.Start(confModel)
	configuration, _ = configurator.SetValue(configuration, 1, "2")

	p1, errP1 := configuration.ParameterById(1)
	p2, errP2 := configuration.ParameterById(2)

	checkFinalParameter(p1, errP1, 1, "P1", "2", t)
	checkFinalParameter(p2, errP2, 2, "P2", "3", t)
}

func TestThatBackwardRuleIsExecuted(t *testing.T) {
	confModel := model.Model{}
	pModel1 := confModel.AddParameter("P1", model.NewIntRangeModel(1, false, 8, false))
	pModel2 := confModel.AddParameter("P2", model.NewIntSetModel([]int{1, 2, 3}))
	confModel.AddConstraint(model.NewSetValueIfFinalConstraintModel(pModel1.Id(), pModel2.Id(), model.NewFinalIntModel(3)))

	configuration := configurator.Start(confModel)
	configuration, _ = configurator.SetValue(configuration, 2, "1")

	p1, errP1 := configuration.ParameterById(1)
	p2, errP2 := configuration.ParameterById(2)

	checkUnselectableParameter(p1, errP1, 1, "P1", t)
	checkFinalParameter(p2, errP2, 2, "P2", "1", t)
}

func TestThatDependingForwardRulesAreExecuted(t *testing.T) {
	confModel := model.Model{}
	pModel1 := confModel.AddParameter("P1", model.NewIntRangeModel(1, false, 8, false))
	pModel2 := confModel.AddParameter("P2", model.NewIntSetModel([]int{1, 2, 3}))
	pModel3 := confModel.AddParameter("P3", model.NewIntSetModel([]int{1, 2, 3}))
	confModel.AddConstraint(model.NewSetValueIfFinalConstraintModel(pModel1.Id(), pModel2.Id(), model.NewFinalIntModel(3)))
	confModel.AddConstraint(model.NewSetValueIfFinalConstraintModel(pModel2.Id(), pModel3.Id(), model.NewIntSetModel([]int{1, 2})))

	configuration := configurator.Start(confModel)
	configuration, _ = configurator.SetValue(configuration, 1, "2")

	p1, errP1 := configuration.ParameterById(1)
	p2, errP2 := configuration.ParameterById(2)
	p3, errP3 := configuration.ParameterById(3)

	checkFinalParameter(p1, errP1, 1, "P1", "2", t)
	checkFinalParameter(p2, errP2, 2, "P2", "3", t)
	checkOpenParameter(p3, errP3, 3, "P3", "{1,2}", t)
}

func TestThatDependingBackwardRulesAreExecuted(t *testing.T) {
	confModel := model.Model{}
	pModel1 := confModel.AddParameter("P1", model.NewIntRangeModel(1, false, 8, false))
	pModel2 := confModel.AddParameter("P2", model.NewIntSetModel([]int{1, 2, 3}))
	pModel3 := confModel.AddParameter("P3", model.NewIntSetModel([]int{1, 2, 3}))
	confModel.AddConstraint(model.NewSetValueIfFinalConstraintModel(pModel1.Id(), pModel2.Id(), model.NewFinalIntModel(3)))
	confModel.AddConstraint(model.NewSetValueIfFinalConstraintModel(pModel2.Id(), pModel3.Id(), model.NewIntSetModel([]int{1, 2})))

	configuration := configurator.Start(confModel)
	configuration, _ = configurator.SetValue(configuration, 3, "3")

	p1, errP1 := configuration.ParameterById(1)
	p2, errP2 := configuration.ParameterById(2)
	p3, errP3 := configuration.ParameterById(3)

	checkUnselectableParameter(p1, errP1, 1, "P1", t)
	checkUnselectableParameter(p2, errP2, 2, "P2", t)
	checkFinalParameter(p3, errP3, 3, "P3", "3", t)
}

func TestThatConstraintImpactIsReverted(t *testing.T) {
	confModel := model.Model{}
	pModel1 := confModel.AddParameter("P1", model.NewIntRangeModel(1, false, 8, false))
	pModel2 := confModel.AddParameter("P2", model.NewIntSetModel([]int{1, 2, 3}))
	confModel.AddConstraint(model.NewSetValueIfValueConstraintModel(pModel1.Id(), model.NewFinalIntModel(7), pModel2.Id(), model.NewFinalIntModel(3)))

	configuration := configurator.Start(confModel)
	configuration, _ = configurator.SetValue(configuration, 1, "7")

	p1, errP1 := configuration.ParameterById(1)
	p2, errP2 := configuration.ParameterById(2)

	checkFinalParameter(p1, errP1, 1, "P1", "7", t)
	checkFinalParameter(p2, errP2, 2, "P2", "3", t)

	configuration, _ = configurator.SetValue(configuration, 1, "5")

	p1, errP1 = configuration.ParameterById(1)
	p2, errP2 = configuration.ParameterById(2)

	checkFinalParameter(p1, errP1, 1, "P1", "5", t)
	checkOpenParameter(p2, errP2, 2, "P2", "{1,2,3}", t)
}

func TestThatConstraintAreCheckedWhenConfigurationStarts(t *testing.T) {
	confModel := model.Model{}
	pModel1 := confModel.AddParameter("P1", model.NewFinalIntModel(7))
	pModel2 := confModel.AddParameter("P2", model.NewIntSetModel([]int{1, 2, 3}))
	confModel.AddConstraint(model.NewSetValueIfValueConstraintModel(pModel1.Id(), model.NewFinalIntModel(7), pModel2.Id(), model.NewFinalIntModel(3)))

	configuration := configurator.Start(confModel)

	p1, errP1 := configuration.ParameterById(1)
	p2, errP2 := configuration.ParameterById(2)

	checkFinalParameter(p1, errP1, 1, "P1", "7", t)
	checkFinalParameter(p2, errP2, 2, "P2", "3", t)
}

func TestThatConstraintExcludesValues(t *testing.T) {
	confModel := model.Model{}
	pModel1 := confModel.AddParameter("P1", model.NewIntRangeModel(1, false, 8, false))
	pModel2 := confModel.AddParameter("P2", model.NewIntSetModel([]int{1, 2, 3}))
	pModel3 := confModel.AddParameter("P3", model.NewIntRangeModel(10, false, 20, false))
	confModel.AddConstraint(model.NewExcludeValueIfValueConstraintModel(pModel1.Id(), model.NewFinalIntModel(7), pModel2.Id(), model.NewFinalIntModel(2)))
	confModel.AddConstraint(model.NewExcludeValueIfValueConstraintModel(pModel1.Id(), model.NewFinalIntModel(7), pModel3.Id(), model.NewFinalIntModel(15)))

	configuration := configurator.Start(confModel)
	configuration, _ = configurator.SetValue(configuration, 1, "7")

	p1, errP1 := configuration.ParameterById(1)
	p2, errP2 := configuration.ParameterById(2)
	p3, errP3 := configuration.ParameterById(3)

	checkFinalParameter(p1, errP1, 1, "P1", "7", t)
	checkOpenParameter(p2, errP2, 2, "P2", "{1,3}", t)
	checkOpenParameter(p3, errP3, 3, "P3", "[[10;14][16;20]]", t)

	configuration, _ = configurator.SetValue(configuration, 1, "5")

	p1, errP1 = configuration.ParameterById(1)
	p2, errP2 = configuration.ParameterById(2)
	p3, errP3 = configuration.ParameterById(3)

	checkFinalParameter(p1, errP1, 1, "P1", "5", t)
	checkOpenParameter(p2, errP2, 2, "P2", "{1,2,3}", t)
	checkOpenParameter(p3, errP3, 3, "P3", "[10;20]", t)
}
