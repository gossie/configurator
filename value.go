package configurator

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type value interface {
	possibleValue(aValue string) bool
	set(aValue string) (value, error)
	terminal() bool
	String() string
}

type intSet struct {
	value []int
}

func (v intSet) possibleValue(aValue string) bool {
	intValue, _ := strconv.Atoi(aValue)
	return slices.Contains(v.value, intValue)
}

func (v intSet) set(aValue string) (value, error) {
	if !v.possibleValue(aValue) {
		return nil, fmt.Errorf("%v is not a possible value for %v", aValue, v)
	}
	intValue, _ := strconv.Atoi(aValue)
	return intSet{value: []int{intValue}}, nil
}

func (v intSet) terminal() bool {
	return len(v.value) == 1
}

func (v intSet) String() string {
	if len(v.value) == 1 {
		return strconv.Itoa(v.value[0])
	}

	strValues := make([]string, 0, len(v.value))
	for _, intValue := range v.value {
		strValues = append(strValues, strconv.Itoa(intValue))
	}
	return "{" + strings.Join(strValues, ",") + "}"
}
