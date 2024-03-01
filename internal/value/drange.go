package value

import (
	"fmt"
	"slices"
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

func (v dRange) subsumedBySet(other intValues) bool {
	for _, r := range v.ranges {
		for i := r.min; i <= r.max; i++ {
			if !slices.Contains(other.values, i) {
				return false
			}
		}
	}
	return true
}

func (v dRange) subsumedByRange(other IntRange) bool {
	for _, r := range v.ranges {
		if r.min < other.min || r.max > other.max {
			return false
		}
	}
	return true
}

func (v dRange) subsumedByDRange(other dRange) bool {
	for _, r := range other.ranges {
		if !r.subsumedByDRange(v) {
			return false
		}
	}
	return true
}

func (v dRange) Sect(other Value) Value {
	return other.sectWithDRange(v)
}

func (v dRange) sectWithSet(other intValues) Value {
	return SectDRangeWithSet(v, other)
}

func (v dRange) sectWithRange(other IntRange) Value {
	return SectRangeWithDRange(other, v)
}

func (v dRange) sectWithDRange(other dRange) Value {
	intersection, _ := SectDRangeWithDRange(v, other)
	return intersection
}

func (v dRange) Diff(other Value) Value {
	return other.diffFromDRange(v)
}

func (v dRange) diffFromSet(other intValues) Value {
	panic("not yet implemented")
}

func (v dRange) diffFromRange(other IntRange) Value {
	result := make([]IntRange, 0)
	rangeToProcess := other
	for _, r := range v.ranges {
		tmp := rangeToProcess.Diff(r)
		if tmpRange, ok := tmp.(IntRange); ok {
			rangeToProcess = tmpRange
		} else if tmpDRange, ok := tmp.(dRange); ok {
			result = append(result, tmpDRange.ranges[0])
			rangeToProcess = tmpDRange.ranges[1]
		}
	}
	result = append(result, rangeToProcess)

	if len(result) == 1 {
		return result[0]
	}
	return NewDRange(result)
}

func (v dRange) diffFromDRange(other dRange) Value {
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
