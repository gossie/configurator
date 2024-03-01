package value

import (
	"fmt"
	"slices"
)

type Value interface {
	Subsumes(aValue Value) bool
	subsumedByRange(other IntRange) bool
	subsumedBySet(other intValues) bool
	subsumedByDRange(other dRange) bool
	Sect(other Value) Value
	sectWithRange(other IntRange) Value
	sectWithSet(other intValues) Value
	sectWithDRange(other dRange) Value
	Diff(other Value) Value
	diffFromRange(other IntRange) Value
	diffFromSet(other intValues) Value
	diffFromDRange(other dRange) Value
	Final() bool
	String() string
}

func InRange(v1 IntRange, v2 int) bool {
	return v2 >= v1.min && v2 <= v1.max
}

func SectRangeWithDRange(v1 IntRange, v2 dRange) Value {
	ranges := make([]IntRange, 0)
	var currentLowerBound int
	searchingLowerBound := true
	for _, r := range v2.ranges {
		if searchingLowerBound {
			if InRange(r, v1.min) {
				currentLowerBound = v1.min
				searchingLowerBound = false
			}

			if InRange(v1, r.min) {
				currentLowerBound = r.min
				searchingLowerBound = false
			}
		}
		if !searchingLowerBound {
			switch {
			case InRange(r, v1.max):
				ranges = append(ranges, NewIntRange(currentLowerBound, false, v1.max, false))
				searchingLowerBound = true
			case InRange(v1, r.max):
				ranges = append(ranges, NewIntRange(currentLowerBound, false, r.max, false))
				searchingLowerBound = true
			}

		}
	}

	if len(ranges) == 1 {
		return ranges[0]
	}
	return NewDRange(ranges)
}

func SectRangeWithSet(v1 IntRange, v2 intValues) Value {
	values := make([]int, 0)
	for _, intValue := range v2.values {
		if InRange(v1, intValue) {
			values = append(values, intValue)
		}
	}
	return intValues{values}
}

func SectDRangeWithSet(v1 dRange, v2 intValues) Value {
	values := make([]int, 0)
	for _, intValue := range v2.values {
		for _, r := range v1.ranges {
			if InRange(r, intValue) {
				values = append(values, intValue)
				break
			}
		}
	}
	return intValues{values}
}

func intersect(v1, v2 IntRange) bool {
	return InRange(v2, v1.min) || InRange(v2, v1.max) || InRange(v1, v2.min) || InRange(v1, v2.max)
}

func SectRangeWithRange(v1, v2 IntRange) (IntRange, error) {
	if intersect(v1, v2) {
		intersection := NewIntRange(max(v1.min, v2.min), false, min(v1.max, v2.max), false)
		return intersection, nil
	}

	var empty IntRange
	return empty, fmt.Errorf("%v and %v do not intersect", v1, v2)
}

func SectSetWithSet(v1, v2 intValues) Value {
	values := make([]int, 0)
	for _, intValue := range v1.values {
		if slices.Contains(v2.values, intValue) {
			values = append(values, intValue)
		}
	}
	return NewIntValues(values)
}

func SectDRangeWithDRange(v1, v2 dRange) (Value, error) {
	result := make([]IntRange, 0)
	i, j := 0, 0
	for i < len(v1.ranges) && j < len(v2.ranges) {
		r1 := v1.ranges[i]
		r2 := v2.ranges[j]

		newRange, err := SectRangeWithRange(r1, r2)
		if err == nil {
			result = append(result, newRange)
			if r1.max <= newRange.max {
				i++
			}
			if r2.max <= newRange.max {
				j++
			}
		} else {
			if r1.max < r2.max {
				i++
			} else {
				j++
			}
		}
	}
	if len(result) == 0 {
		var empty dRange
		return empty, fmt.Errorf("%v and %v do not intersect", v1, v2)
	}
	return NewDRange(result), nil
}
