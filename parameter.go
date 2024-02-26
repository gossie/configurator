package configurator

import "fmt"

type Parameter struct {
	id    int
	name  string
	value value
}

func (p Parameter) Id() int {
	return p.id
}

func (p Parameter) Name() string {
	return p.name
}

func (p Parameter) Terminal() bool {
	return p.value.terminal()
}

func (p Parameter) Value() string {
	return p.value.String()
}

func (p *Parameter) SetValue(aValue string) error {
	if p.value.possibleValue(aValue) {
		p.value, _ = p.value.set(aValue)
		return nil
	}
	return fmt.Errorf("%v is not a possible value for %v", aValue, p.value)
}