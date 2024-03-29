package value

import (
	"fmt"
	"slices"
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

func (v IntRange) Subsumes(other Value) bool {
	return other.subsumedByRange(v)
}

func (v IntRange) subsumedBySet(other intValues) bool {
	for i := v.min; i <= v.max; i++ {
		if !slices.Contains(other.values, i) {
			return false
		}
	}
	return true
}

func (v IntRange) subsumedByRange(other IntRange) bool {
	return v.min >= other.min && v.max <= other.max
}

func (v IntRange) subsumedByDRange(other dRange) bool {
	for _, r := range other.ranges {
		if v.subsumedByRange(r) {
			return true
		}
	}
	return false
}

func (v IntRange) Sect(other Value) Value {
	return other.sectWithRange(v)
}

func (v IntRange) sectWithSet(other intValues) Value {
	return SectRangeWithSet(v, other)
}

func (v IntRange) sectWithRange(other IntRange) Value {
	intersection, _ := SectRangeWithRange(v, other)
	return intersection
}

func (v IntRange) sectWithDRange(other dRange) Value {
	return SectRangeWithDRange(v, other)
}

func (v IntRange) Diff(other Value) Value {
	return other.diffFromRange(v)
}

func (v IntRange) diffFromSet(aValue intValues) Value {
	values := make([]int, 0)
	for _, intValue := range aValue.values {
		if !InRange(v, intValue) {
			values = append(values, intValue)
		}
	}
	return NewIntValues(values)
}

func (v IntRange) diffFromRange(other IntRange) Value {
	if !intersect(v, other) {
		return other
	}

	if other.Subsumes(v) {
		switch {
		case v.min == other.min:
			return NewIntRange(v.max+1, false, other.max, false)
		case v.max == other.max:
			return NewIntRange(other.min, false, v.min-1, false)
		default:
			return NewDRange([]IntRange{NewIntRange(other.min, false, v.min-1, false), NewIntRange(v.max+1, false, other.max, false)})
		}
	}

	newLowerBound := other.min
	newUpperBound := v.min - 1
	if InRange(v, other.min) {
		newLowerBound = v.max + 1
		newUpperBound = other.max
	}

	return NewIntRange(newLowerBound, false, newUpperBound, false)
}

func (v IntRange) diffFromDRange(other dRange) Value {
	result := make([]IntRange, 0)
	for _, r := range other.ranges {
		tmp := r.Diff(v)
		if tmpRange, ok := tmp.(IntRange); ok {
			result = append(result, tmpRange)
		} else if tmpDRange, ok := tmp.(dRange); ok {
			result = append(result, tmpDRange.ranges...)
		}
	}

	if len(result) == 1 {
		return result[0]
	}
	return NewDRange(result)
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
