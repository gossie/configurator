package value

import (
	"fmt"
	"strconv"
)

type IntRange struct {
	min, max         int
	minOpen, maxOpen bool
}

func NewIntRange(min int, minOpen bool, max int, maxOpen bool) IntRange {
	return IntRange{
		min,
		max,
		minOpen,
		maxOpen,
	}
}

func (v IntRange) Subsumes(aValue Value) bool {
	return aValue.subsumedByRange(v)
}

func (v IntRange) subsumedBySet(aValue intValues) bool {
	return false
}

func (v IntRange) subsumedByRange(aValue IntRange) bool {
	return v.min >= aValue.min && v.max <= aValue.max
}

func (v IntRange) subsumedByDRange(aValue dRange) bool {
	panic("not yet implemented")
}

func (v IntRange) Sect(other Value) Value {
	return other.sectWithRange(v)
}

func (v IntRange) sectWithSet(aValue intValues) Value {
	values := make([]int, 0)
	for _, intValue := range aValue.values {
		if v.min <= intValue && v.max >= intValue {
			values = append(values, intValue)
		}
	}
	return NewIntValues(values)
}

func (v IntRange) sectWithRange(aValue IntRange) Value {
	return NewIntRange(max(v.min, aValue.min), false, min(v.min, aValue.min), false)
}

func (v IntRange) sectWithDRange(other dRange) Value {
	return sectRangeWithDRange(v, other)
}

func (v IntRange) Diff(other Value) Value {
	return other.diffFromRange(v)
}

func (v IntRange) diffFromSet(aValue intValues) Value {
	panic("not yet implemented")
}

func (v IntRange) diffFromRange(aValue IntRange) Value {
	panic("not yet implemented")
}

func (v IntRange) diffFromDRange(aValue dRange) Value {
	panic("not yet implemented")
}

func (v IntRange) Final() bool {
	return v.min == v.max
}

func (v IntRange) String() string {
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
