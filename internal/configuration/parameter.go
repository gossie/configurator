package configuration

import (
	"fmt"

	"github.com/gossie/configurator/internal/value"
)

type InternalParameter struct {
	id          int
	name        string
	value       value.Value
	selectable  bool
	Constraints []Constraint
}

func NewInternalParameter(id int, name string, value value.Value) InternalParameter {
	return InternalParameter{
		id:         id,
		name:       name,
		value:      value,
		selectable: true,
	}
}

func (p InternalParameter) Id() int {
	return p.id
}

func (p InternalParameter) Name() string {
	return p.name
}

func (p InternalParameter) Selectable() bool {
	return p.selectable
}

func (p InternalParameter) Final() bool {
	return p.value.Final()
}

func (p InternalParameter) Value() string {
	return p.value.String()
}

func (p *InternalParameter) AppendConstraint(newConstraints Constraint) {
	p.Constraints = append(p.Constraints, newConstraints)
}

func (p *InternalParameter) SetValue(aValue value.Value) (bool, error) {
	if !p.selectable {
		return false, fmt.Errorf("parameter with ID %v is not selectable", p.id)
	}

	if !p.value.Subsumes(aValue) {
		return false, fmt.Errorf("parameter with ID %v: %v is not a possible value for %v", p.id, aValue, p.value)
	}

	p.value = aValue
	return true, nil
}
