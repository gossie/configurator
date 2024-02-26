package configurator_test

import (
	"testing"

	"github.com/gossie/configurator"
)

func checkOpenParameter(param configurator.Parameter, err error, expectedId int, expectedName string, expectedValue string, t *testing.T) {
	if err != nil {
		t.Fatalf("parameter with ID %v could not be found: %v", expectedId, err.Error())
	}
	if param.Id() != expectedId {
		t.Fatalf("ID of the parameter should be %v but was %v", expectedId, param.Id())
	}
	if param.Name() != expectedName {
		t.Fatalf("name of parameter with ID %v should be P3 but was %v", expectedId, param.Name())
	}
	if param.Final() {
		t.Fatalf("value %v of parameter with ID %v should not be terminal", expectedId, param.Value())
	}
	if param.Value() != expectedValue {
		t.Fatalf("value of parameter with ID %v should be %v but was %v", expectedId, expectedValue, param.Value())
	}
}

func TestThatConfigurationIsStarted(t *testing.T) {
	pModel1 := configurator.NewParameterModel("P1", configurator.NewIntSetModel([]int{1, 2, 3}))
	pModel2 := configurator.NewParameterModel("P2", configurator.NewIntSetModel([]int{1, 2, 3}))
	pModel3 := configurator.NewParameterModel("P3", configurator.NewIntRangeModel(7, false, 17, false))
	pModel4 := configurator.NewParameterModel("P4", configurator.NewIntRangeModel(7, false, 17, true))
	pModel5 := configurator.NewParameterModel("P5", configurator.NewIntRangeModel(7, true, 17, false))
	pModel6 := configurator.NewParameterModel("P6", configurator.NewIntRangeModel(7, true, 17, true))

	model := configurator.Model{}
	model.AddParameter(pModel1)
	model.AddParameter(pModel2)
	model.AddParameter(pModel3)
	model.AddParameter(pModel4)
	model.AddParameter(pModel5)
	model.AddParameter(pModel6)

	configuration := configurator.Start(model)

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
	pModel1 := configurator.NewParameterModel("P1", configurator.NewIntSetModel([]int{1, 2, 3}))

	model := configurator.Model{}
	model.AddParameter(pModel1)

	configuration := configurator.Start(model)

	_, err := configuration.ParameterById(2)
	if err == nil {
		t.Fatal("parameter with ID 2 should not be found")
	}
}

func TestThatValueIsNotSetBecauseTheParameterIsMissing(t *testing.T) {
	pModel1 := configurator.NewParameterModel("P1", configurator.NewIntSetModel([]int{1, 2, 3}))

	model := configurator.Model{}
	model.AddParameter(pModel1)

	configuration := configurator.Start(model)
	configuration, _ = configurator.SetValue(configuration, 2, "2")

	_, err := configuration.ParameterById(2)

	if err == nil {
		t.Fatal("parameter with ID 2 should not be found")
	}
}

func TestThatValueIsSet_intValues(t *testing.T) {
	pModel1 := configurator.NewParameterModel("P1", configurator.NewIntSetModel([]int{1, 2, 3}))
	pModel2 := configurator.NewParameterModel("P2", configurator.NewIntSetModel([]int{1, 2, 3}))

	model := configurator.Model{}
	model.AddParameter(pModel1)
	model.AddParameter(pModel2)

	configuration := configurator.Start(model)
	configuration, _ = configurator.SetValue(configuration, 1, "2")

	p1, _ := configuration.ParameterById(1)
	p2, _ := configuration.ParameterById(2)

	if !p1.Final() {
		t.Fatalf("value %v of parameter with ID 1 should be terminal", p1.Value())
	}
	if p1.Value() != "2" {
		t.Fatalf("value of parameter with ID 1 should be 2 but was %v", p1.Value())
	}

	if p2.Final() {
		t.Fatalf("value %v of parameter with ID 2 should not be terminal", p2.Value())
	}
	if p2.Value() != "{1,2,3}" {
		t.Fatalf("value of parameter with ID 2 should be {1,2,3} but was %v", p2.Value())
	}
}

func TestThatAnImpossibleValueIsNotSet_intValues(t *testing.T) {
	pModel1 := configurator.NewParameterModel("P1", configurator.NewIntSetModel([]int{1, 2, 3}))

	model := configurator.Model{}
	model.AddParameter(pModel1)

	configuration := configurator.Start(model)
	configuration, err := configurator.SetValue(configuration, 1, "4")
	if err == nil {
		t.Fatal("4 is not a valid value and should cause an error")
	}

	p1, _ := configuration.ParameterById(1)

	if p1.Value() != "{1,2,3}" {
		t.Fatalf("value of parameter with ID 1 should be {1,2,3} but was %v", p1.Value())
	}
}

func TestThatValueIsSet_intRange(t *testing.T) {
	pModel1 := configurator.NewParameterModel("P1", configurator.NewIntRangeModel(1, false, 8, false))
	pModel2 := configurator.NewParameterModel("P2", configurator.NewIntSetModel([]int{1, 2, 3}))

	model := configurator.Model{}
	model.AddParameter(pModel1)
	model.AddParameter(pModel2)

	configuration := configurator.Start(model)
	configuration, _ = configurator.SetValue(configuration, 1, "2")

	p1, _ := configuration.ParameterById(1)
	p2, _ := configuration.ParameterById(2)

	if !p1.Final() {
		t.Fatalf("value %v of parameter with ID 1 should be terminal", p1.Value())
	}
	if p1.Value() != "2" {
		t.Fatalf("value of parameter with ID 1 should be 2 but was %v", p1.Value())
	}

	if p2.Final() {
		t.Fatalf("value %v of parameter with ID 2 should not be terminal", p2.Value())
	}
	if p2.Value() != "{1,2,3}" {
		t.Fatalf("value of parameter with ID 2 should be {1,2,3} but was %v", p2.Value())
	}
}

func TestThatAnImpossibleValueIsNotSet_intRange(t *testing.T) {
	pModel1 := configurator.NewParameterModel("P1", configurator.NewIntRangeModel(1, false, 8, false))

	model := configurator.Model{}
	model.AddParameter(pModel1)

	configuration := configurator.Start(model)
	configuration, err := configurator.SetValue(configuration, 1, "9")
	if err == nil {
		t.Fatal("4 is not a valid value and should cause an error")
	}

	p1, _ := configuration.ParameterById(1)

	if p1.Value() != "[1;8]" {
		t.Fatalf("value of parameter with ID 1 should be [1;8] but was %v", p1.Value())
	}
}
