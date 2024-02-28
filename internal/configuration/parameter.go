package configuration

import (
	"fmt"
	"slices"

	"github.com/gossie/configurator/internal/value"
)

type restrictionType int

const (
	user restrictionType = iota
	system
)

type restriction struct {
	srcId           int
	restrictionType restrictionType
	value           value.Value
}

type InternalParameter struct {
	id            int
	name          string
	restrictions  []restriction
	originalValue value.Value
	selectable    bool
	Constraints   []Constraint
}

func NewInternalParameter(id int, name string, value value.Value) InternalParameter {
	return InternalParameter{
		id:            id,
		name:          name,
		originalValue: value,
		selectable:    true,
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

func (p InternalParameter) RestrictedValue() value.Value {
	currentValue := p.originalValue
	for _, r := range p.restrictions {
		currentValue = value.Sect(currentValue, r.value)
	}
	return currentValue
}

func (p InternalParameter) Final() bool {
	return p.RestrictedValue().Final()
}

func (p InternalParameter) Value() string {
	return p.RestrictedValue().String()
}

func (p *InternalParameter) AppendConstraint(newConstraints Constraint) {
	p.Constraints = append(p.Constraints, newConstraints)
}

func (p *InternalParameter) RestrictValue(aValue value.Value) (bool, error) {
	if !p.selectable {
		return false, fmt.Errorf("parameter with ID %v is not selectable", p.id)
	}

	p.restrictions = slices.DeleteFunc(p.restrictions, func(r restriction) bool {
		return r.restrictionType == user
	})

	if !p.RestrictedValue().Subsumes(aValue) {
		return false, fmt.Errorf("parameter with ID %v: %v is not a possible value for %v", p.id, aValue, p.RestrictedValue())
	}

	p.restrictions = append(p.restrictions, restriction{restrictionType: user, value: aValue})
	return true, nil
}

func (p *InternalParameter) RestrictValueFromConstraint(srcId int, aValue value.Value) (bool, error) {
	if !p.selectable {
		return false, fmt.Errorf("parameter with ID %v is not selectable", p.id)
	}

	p.restrictions = slices.DeleteFunc(p.restrictions, func(r restriction) bool {
		return r.restrictionType == system && r.srcId == srcId
	})

	if !p.RestrictedValue().Subsumes(aValue) {
		return false, fmt.Errorf("parameter with ID %v: %v is not a possible value for %v", p.id, aValue, p.RestrictedValue())
	}

	p.restrictions = append(p.restrictions, restriction{
		restrictionType: system,
		srcId:           srcId,
		value:           aValue,
	})
	return true, nil
}

func (p *InternalParameter) RemoveRestrictionsFromConstraint(srcId int) {
	p.restrictions = slices.DeleteFunc(p.restrictions, func(r restriction) bool {
		return r.restrictionType == system && r.srcId == srcId
	})
}
