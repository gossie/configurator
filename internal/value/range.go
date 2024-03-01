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
