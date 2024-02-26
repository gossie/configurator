package configurator_test

import (
	"testing"

	"github.com/gossie/configurator"
)

func TestThatConfigurationIsStarted(t *testing.T) {
	pModel1 := configurator.NewParameterModel("P1", configurator.NewIntSetModel([]int{1, 2, 3}))
	pModel2 := configurator.NewParameterModel("P2", configurator.NewIntSetModel([]int{1, 2, 3}))

	model := configurator.Model{}
	model.AddParameter(pModel1)
	model.AddParameter(pModel2)

	configuration := configurator.Start(model)

	p1, errP1 := configuration.ParameterById(1)
	if errP1 != nil {
		t.Fatal("parameter with ID 1 could not be found", errP1.Error())
	}
	if p1.Name() != "P1" {
		t.Fatalf("name of parameter with ID 1 should be P1 but was %v", p1.Name())
	}
	if p1.Value() != "{1,2,3}" {
		t.Fatalf("value of parameter with ID 1 should be {1,2,3} but was %v", p1.Value())
	}

	p2, errP2 := configuration.ParameterById(2)
	if errP2 != nil {
		t.Fatal("parameter with ID 2 could not be found", errP2.Error())
	}
	if p2.Name() != "P2" {
		t.Fatalf("name of parameter with ID 2 should be P2 but was %v", p2.Name())
	}
	if p2.Value() != "{1,2,3}" {
		t.Fatalf("value of parameter with ID 2 should be {1,2,3} but was %v", p2.Value())
	}
}

func TestThatValueIsSet(t *testing.T) {
	pModel1 := configurator.NewParameterModel("P1", configurator.NewIntSetModel([]int{1, 2, 3}))
	pModel2 := configurator.NewParameterModel("P2", configurator.NewIntSetModel([]int{1, 2, 3}))

	model := configurator.Model{}
	model.AddParameter(pModel1)
	model.AddParameter(pModel2)

	configuration := configurator.Start(model)
	configuration, _ = configurator.SetValue(configuration, 1, "2")

	p1, _ := configuration.ParameterById(1)
	p2, _ := configuration.ParameterById(2)

	if p1.Value() != "2" {
		t.Fatalf("value of parameter with ID 2 should be 2 but was %v", p1.Value())
	}

	if p2.Value() != "{1,2,3}" {
		t.Fatalf("value of parameter with ID 2 should be {1,2,3} but was %v", p2.Value())
	}
}
