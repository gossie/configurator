package configurator

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type value interface {
	possibleValue(aValue value) bool
	subsumedByRange(other intRange) bool
	subsumedBySet(other intValues) bool
	//set(aValue string) (value, error)
	final() bool
	String() string
}

type intValues struct {
	value []int
}

func (v intValues) possibleValue(aValue value) bool {
	return aValue.subsumedBySet(v)
}

func (v intValues) subsumedBySet(aValue intValues) bool {
	for _, intValue := range v.value {
		if !slices.Contains(aValue.value, intValue) {
			return false
		}
	}
	return true
}

func (v intValues) subsumedByRange(aValue intRange) bool {
	for _, intValue := range v.value {
		if intValue < aValue.min || intValue > aValue.max {
			return false
		}
	}
	return true
}

/*
	func (v intValues) set(aValue string) (value, error) {
		if !v.possibleValue(aValue) {
			return nil, fmt.Errorf("%v is not a possible value for %v", aValue, v)
		}
		intValue, _ := strconv.Atoi(aValue)
		return intValues{value: []int{intValue}}, nil
	}
*/
func (v intValues) final() bool {
	return len(v.value) == 1
}

func (v intValues) String() string {
	if v.final() {
		return strconv.Itoa(v.value[0])
	}

	strValues := make([]string, 0, len(v.value))
	for _, intValue := range v.value {
		strValues = append(strValues, strconv.Itoa(intValue))
	}
	return "{" + strings.Join(strValues, ",") + "}"
}

type intRange struct {
	min, max         int
	minOpen, maxOpen bool
}

func (v intRange) possibleValue(aValue value) bool {
	return aValue.subsumedByRange(v)
}

func (v intRange) subsumedBySet(aValue intValues) bool {
	return false
}

func (v intRange) subsumedByRange(aValue intRange) bool {
	return v.min >= aValue.min && v.max <= aValue.max
}

/*
	func (v intRange) set(aValue string) (value, error) {
		if !v.possibleValue(aValue) {
			return nil, fmt.Errorf("%v is not a possible value for %v", aValue, v)
		}
		intValue, _ := strconv.Atoi(aValue)
		return intRange{min: intValue, minOpen: false, max: intValue, maxOpen: false}, nil
	}
*/

func (v intRange) final() bool {
	return v.min == v.max
}

func (v intRange) String() string {
	if v.final() {
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
