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
	final() bool
	String() string
}

type intValues struct {
	value []int
}

func (v intValues) possibleValue(aValue string) bool {
	intValue, _ := strconv.Atoi(aValue)
	return slices.Contains(v.value, intValue)
}

func (v intValues) set(aValue string) (value, error) {
	if !v.possibleValue(aValue) {
		return nil, fmt.Errorf("%v is not a possible value for %v", aValue, v)
	}
	intValue, _ := strconv.Atoi(aValue)
	return intValues{value: []int{intValue}}, nil
}

func (v intValues) final() bool {
	return len(v.value) == 1
}

func (v intValues) String() string {
	if len(v.value) == 1 {
		return strconv.Itoa(v.value[0])
	}

	strValues := make([]string, 0, len(v.value))
	for _, intValue := range v.value {
		strValues = append(strValues, strconv.Itoa(intValue))
	}
	return "{" + strings.Join(strValues, ",") + "}"
}
