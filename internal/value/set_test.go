package value_test

import (
	"testing"

	"github.com/gossie/configurator/internal/value"
	"github.com/stretchr/testify/assert"
)

func TestThatSetSubsumesSet(t *testing.T) {
	set1 := value.NewIntValues([]int{1, 2, 3, 4})
	set2 := value.NewIntValues([]int{2, 3})

	assert.True(t, set1.Subsumes(set2))
}

func TestThatSetDoesNotSubsumeSet(t *testing.T) {
	set1 := value.NewIntValues([]int{1, 2, 3, 4})
	set2 := value.NewIntValues([]int{2, 3, 5})

	assert.False(t, set1.Subsumes(set2))
}

func TestThatSetSubsumesDRange(t *testing.T) {
	set := value.NewIntValues([]int{1, 2, 3, 4})
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(1, false, 2, false), value.NewIntRange(3, false, 3, false)})

	assert.True(t, set.Subsumes(dRange))
}

func TestThatSetDoesNotSubsumeDRange(t *testing.T) {
	set := value.NewIntValues([]int{1, 2, 4})
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(1, false, 2, false), value.NewIntRange(3, false, 3, false)})

	assert.False(t, set.Subsumes(dRange))
}

func TestThatSetSubsumesRange(t *testing.T) {
	set := value.NewIntValues([]int{1, 2, 3, 4})
	aRange := value.NewIntRange(2, false, 4, false)

	assert.True(t, set.Subsumes(aRange))
}

func TestThatSetDoesNotSubsumeRange(t *testing.T) {
	set := value.NewIntValues([]int{1, 2, 4})

	aRange := value.NewIntRange(2, false, 4, false)

	assert.False(t, set.Subsumes(aRange))
}

func TestThatRangeIsSubtractedFromSet(t *testing.T) {
	set := value.NewIntValues([]int{1, 2, 3, 4, 5, 6, 7})
	r := value.NewIntRange(3, false, 5, false)

	expected := value.NewIntValues([]int{1, 2, 6, 7})

	assert.Equal(t, expected, set.Diff(r))
}

func TestThatRangeIsSubtractedFromSet_noIntersection(t *testing.T) {
	set := value.NewIntValues([]int{1, 2, 3, 4, 5, 6, 7})
	r := value.NewIntRange(8, false, 16, false)

	expected := value.NewIntValues([]int{1, 2, 3, 4, 5, 6, 7})

	assert.Equal(t, expected, set.Diff(r))
}

func TestThatDRangeIsSubtractedFromSet(t *testing.T) {
	set := value.NewIntValues([]int{1, 2, 3, 4, 5, 6, 7})
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(-10, false, 2, false), value.NewIntRange(6, false, 16, false)})

	expected := value.NewIntValues([]int{3, 4, 5})

	assert.Equal(t, expected, set.Diff(dRange))
}

func TestThatDRangeIsSubtractedFromSet_noIntersection(t *testing.T) {
	set := value.NewIntValues([]int{1, 2, 3, 4, 5, 6, 7})
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(-10, false, 0, false), value.NewIntRange(8, false, 16, false)})

	expected := value.NewIntValues([]int{1, 2, 3, 4, 5, 6, 7})

	assert.Equal(t, expected, set.Diff(dRange))
}

func TestSetString(t *testing.T) {
	set := value.NewIntValues([]int{1, 2, 3, 4, 5, 6, 7})
	assert.Equal(t, "{1,2,3,4,5,6,7}", set.String())
}
