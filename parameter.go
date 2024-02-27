package configurator

import "fmt"

type Parameter struct {
	id          int
	name        string
	value       value
	constraints []constraint
}

func (p Parameter) Id() int {
	return p.id
}

func (p Parameter) Name() string {
	return p.name
}

func (p Parameter) Final() bool {
	return p.value.final()
}

func (p Parameter) Value() string {
	return p.value.String()
}

func (p *Parameter) SetValue(aValue value) error {
	if p.value.possibleValue(aValue) {
		p.value = aValue
		return nil
	}
	return fmt.Errorf("%v is not a possible value for %v", aValue, p.value)
}
