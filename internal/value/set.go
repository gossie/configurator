package value

import (
	"slices"
	"strconv"
	"strings"
)

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

func (v intValues) subsumedByRange(aValue IntRange) bool {
	for _, intValue := range v.values {
		if intValue < aValue.min || intValue > aValue.max {
			return false
		}
	}
	return true
}

func (v intValues) subsumedByDRange(aValue dRange) bool {
	panic("not yet implemented")
}

func (v intValues) Sect(other Value) Value {
	return other.sectWithSet(v)
}

func (v intValues) sectWithSet(aValue intValues) Value {
	values := make([]int, 0)
	for _, intValue := range v.values {
		if slices.Contains(aValue.values, intValue) {
			values = append(values, intValue)
		}
	}
	return NewIntValues(values)
}

func (v intValues) sectWithRange(aValue IntRange) Value {
	values := make([]int, 0)
	for _, intValue := range v.values {
		if aValue.min <= intValue && aValue.max >= intValue {
			values = append(values, intValue)
		}
	}
	return NewIntValues(values)
}

func (v intValues) sectWithDRange(aValue dRange) Value {
	panic("not yet implemented")
}

func (v intValues) Diff(other Value) Value {
	return other.diffFromSet(v)
}

func (v intValues) diffFromSet(aValue intValues) Value {
	values := make([]int, 0)
	for _, intValue := range aValue.values {
		if !slices.Contains(v.values, intValue) {
			values = append(values, intValue)
		}
	}
	return NewIntValues(values)
}

func (v intValues) diffFromRange(aValue IntRange) Value {
	ranges := make([]IntRange, 0)
	lowerBound := aValue.min
	for _, intValue := range v.values {
		ranges = append(ranges, NewIntRange(lowerBound, false, intValue-1, false))
		lowerBound = intValue + 1
	}
	ranges = append(ranges, NewIntRange(lowerBound, false, aValue.max, false))
	return NewDRange(ranges)
}

func (v intValues) diffFromDRange(aValue dRange) Value {
	panic("not yet implemented")
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
