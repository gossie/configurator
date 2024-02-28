package value

import (
	"fmt"
	"strconv"
)

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

func (v intRange) Sect(other Value) Value {
	return other.sectWithRange(v)
}

func (v intRange) sectWithSet(aValue intValues) Value {
	values := make([]int, 0)
	for _, intValue := range aValue.values {
		if v.min <= intValue && v.max >= intValue {
			values = append(values, intValue)
		}
	}
	return NewIntValues(values)
}

func (v intRange) sectWithRange(aValue intRange) Value {
	return NewIntRange(max(v.min, aValue.min), false, min(v.min, aValue.min), false)
}

func (v intRange) Diff(other Value) Value {
	return other.diffFromRange(v)
}

func (v intRange) diffFromSet(aValue intValues) Value {
	// TODO
	panic("not yet implemented")
}

func (v intRange) diffFromRange(aValue intRange) Value {
	// TODO
	panic("not yet implemented")
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
