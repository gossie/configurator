package configurator_test

import (
	"testing"

	"github.com/gossie/configurator"
)

func TestThatConfigurationIsStarted(t *testing.T) {
	model := configurator.Model{}
	model.AddParameter("P1")
	model.AddParameter("P2")

	configuration := configurator.Start(model)

	p1, errP1 := configuration.ParameterById(1)
	if errP1 != nil {
		t.Fatal("parameter with ID 1 could not be found", errP1.Error())
	}
	if p1.Name() != "P1" {
		t.Fatalf("name of parameter with ID 1 should be P1 but was %v", p1.Name())
	}

	p2, errP2 := configuration.ParameterById(2)
	if errP2 != nil {
		t.Fatal("parameter with ID 2 could not be found", errP2.Error())
	}
	if p2.Name() != "P2" {
		t.Fatalf("name of parameter with ID 2 should be P2 but was %v", p2.Name())
	}
}
