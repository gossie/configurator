package value

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Value interface {
	Subsumes(aValue Value) bool
	subsumedByRange(other intRange) bool
	subsumedBySet(other intValues) bool
	Final() bool
	String() string
}

type intValues struct {
	values []int
}

func NewIntValues(values []int) Value {
	return intValues{values}
}

func (v intValues) Subsumes(aValue Value) bool {
	return aValue.subsumedBySet(v)
}

func (v intValues) subsumedBySet(aValue intValues) bool {
	for _, intValue := range v.values {
		if !slices.Contains(aValue.values, intValue) {
			return false
		}
	}
	return true
}

func (v intValues) subsumedByRange(aValue intRange) bool {
	for _, intValue := range v.values {
		if intValue < aValue.min || intValue > aValue.max {
			return false
		}
	}
	return true
}

func (v intValues) Final() bool {
	return len(v.values) == 1
}

func (v intValues) String() string {
	if v.Final() {
		return strconv.Itoa(v.values[0])
	}

	strValues := make([]string, 0, len(v.values))
	for _, intValue := range v.values {
		strValues = append(strValues, strconv.Itoa(intValue))
	}
	return "{" + strings.Join(strValues, ",") + "}"
}

type intRange struct {
	min, max         int
	minOpen, maxOpen bool
}

func NewIntRange(min int, minOpen bool, max int, maxOpen bool) Value {
	return intRange{
		min,
		max,
		minOpen,
		maxOpen,
	}
}

func (v intRange) Subsumes(aValue Value) bool {
	return aValue.subsumedByRange(v)
}

func (v intRange) subsumedBySet(aValue intValues) bool {
	return false
}

func (v intRange) subsumedByRange(aValue intRange) bool {
	return v.min >= aValue.min && v.max <= aValue.max
}

func (v intRange) Final() bool {
	return v.min == v.max
}

func (v intRange) String() string {
	if v.Final() {
		return strconv.Itoa(v.min)
	}

	lower := "["
	if v.minOpen {
		lower = "]"
	}

	upper := "]"
	if v.maxOpen {
		upper = "["
	}

	return fmt.Sprintf("%v%v;%v%v", lower, v.min, v.max, upper)
}
