package value

import (
	"fmt"
	"strings"
)

type dRange struct {
	ranges []IntRange
}

func NewDRange(ranges []IntRange) dRange {
	return dRange{ranges}
}

func (v dRange) Subsumes(aValue Value) bool {
	return aValue.subsumedByDRange(v)
}

func (v dRange) subsumedBySet(aValue intValues) bool {
	panic("not yet implemente")
}

func (v dRange) subsumedByRange(other IntRange) bool {
	for _, r := range v.ranges {
		if r.min < other.min || r.max > other.max {
			return false
		}
	}
	return true
}

func (v dRange) subsumedByDRange(aValue dRange) bool {
	panic("not yet implemente")
}

func (v dRange) Sect(other Value) Value {
	return other.sectWithDRange(v)
}

func (v dRange) sectWithSet(aValue intValues) Value {
	panic("not yet implemente")
}

func (v dRange) sectWithRange(other IntRange) Value {
	return sectRangeWithDRange(other, v)
}

func (v dRange) sectWithDRange(aValue dRange) Value {
	panic("not yet implemente")
}

func (v dRange) Diff(other Value) Value {
	return other.diffFromDRange(v)
}

func (v dRange) diffFromSet(aValue intValues) Value {
	panic("not yet implemented")
}

func (v dRange) diffFromRange(aValue IntRange) Value {
	panic("not yet implemented")
}

func (v dRange) diffFromDRange(aValue dRange) Value {
	panic("not yet implemented")
}

func (v dRange) Final() bool {
	return false
}

func (v dRange) String() string {
	rangeStrings := make([]string, 0, len(v.ranges))
	for _, r := range v.ranges {
		rangeStrings = append(rangeStrings, r.String())
	}
	return fmt.Sprintf("[%v]", strings.Join(rangeStrings, ""))
}
