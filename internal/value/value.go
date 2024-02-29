package value

import "slices"

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

func SectRangeWithDRange(v1 IntRange, v2 dRange) Value {
	ranges := make([]IntRange, 0)
	var currentLowerBound int
	searchingLowerBound := true
	for _, r := range v2.ranges {
		if searchingLowerBound {
			if v1.min >= r.min && v1.min <= r.max {
				currentLowerBound = v1.min
				searchingLowerBound = false
			}

			if r.min >= v1.min && r.min <= v1.max {
				currentLowerBound = r.min
				searchingLowerBound = false
			}
		}
		if !searchingLowerBound {
			switch {
			case v1.max >= r.min && v1.max <= r.max:
				ranges = append(ranges, NewIntRange(currentLowerBound, false, v1.max, false))
				searchingLowerBound = true
			case r.max >= v1.min && r.max <= v1.max:
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
		if intValue >= v1.min && intValue <= v1.max {
			values = append(values, intValue)
		}
	}
	return intValues{values}
}

func SectDRangeWithSet(v1 dRange, v2 intValues) Value {
	values := make([]int, 0)
	for _, intValue := range v2.values {
		for _, r := range v1.ranges {
			if intValue >= r.min && intValue <= r.max {
				values = append(values, intValue)
				break
			}
		}
	}
	return intValues{values}
}

func SectRangeWithRange(v1, v2 IntRange) Value {
	return NewIntRange(max(v1.min, v2.min), false, min(v1.max, v2.max), false)
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
