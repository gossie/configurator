package value_test

import (
	"testing"

	"github.com/gossie/configurator/internal/value"
	"github.com/stretchr/testify/assert"
)

func TestThatDRangeSubsumesDRange(t *testing.T) {
	dRange1 := value.NewDRange([]value.IntRange{value.NewIntRange(8, false, 12, false), value.NewIntRange(14, false, 16, false)})
	dRange2 := value.NewDRange([]value.IntRange{value.NewIntRange(9, false, 11, false), value.NewIntRange(12, false, 15, false)})

	assert.False(t, dRange1.Subsumes(dRange2))
}

func TestThatDRangeDoesNotSubsumeDRange(t *testing.T) {
	dRange1 := value.NewDRange([]value.IntRange{value.NewIntRange(8, false, 12, false), value.NewIntRange(14, false, 16, false)})
	dRange2 := value.NewDRange([]value.IntRange{value.NewIntRange(9, false, 11, false), value.NewIntRange(12, false, 17, false)})

	assert.False(t, dRange1.Subsumes(dRange2))
}

func TestThatDRangeSubsumesRange(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(8, false, 12, false), value.NewIntRange(14, false, 16, false)})
	r := value.NewIntRange(9, false, 11, false)

	assert.True(t, dRange.Subsumes(r))
}

func TestThatDRangeDoesNotSubsumeRange(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(8, false, 12, false), value.NewIntRange(14, false, 16, false)})
	r := value.NewIntRange(11, false, 15, false)

	assert.False(t, dRange.Subsumes(r))
}

func TestThatDRangeSubsumesSet(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(8, false, 12, false), value.NewIntRange(14, false, 16, false)})
	r := value.NewIntValues([]int{9, 10, 14, 15})

	assert.True(t, dRange.Subsumes(r))
}

func TestThatDRangeDoesNotSubsumeSet(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(8, false, 12, false), value.NewIntRange(14, false, 16, false)})
	r := value.NewIntValues([]int{9, 10, 13, 14, 15})

	assert.False(t, dRange.Subsumes(r))
}

func TestThatSetIsSubtractedFromDRange(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(3, false, 7, false), value.NewIntRange(12, false, 20, false)})
	set := value.NewIntValues([]int{2, 4, 8, 16})

	expected := value.NewDRange([]value.IntRange{value.NewIntRange(3, false, 3, false), value.NewIntRange(5, false, 7, false), value.NewIntRange(12, false, 15, false), value.NewIntRange(17, false, 20, false)})

	assert.Equal(t, expected, dRange.Diff(set))
}

func TestThatSetIsSubtractedFromDRange_noIntersection(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(3, false, 7, false), value.NewIntRange(12, false, 20, false)})
	set := value.NewIntValues([]int{2, 10, 25})

	expected := value.NewDRange([]value.IntRange{value.NewIntRange(3, false, 7, false), value.NewIntRange(12, false, 20, false)})

	assert.Equal(t, expected, dRange.Diff(set))
}

func TestThatRangeIsSubtractedFromDRange(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(3, false, 7, false), value.NewIntRange(12, false, 20, false)})
	r := value.NewIntRange(6, false, 15, false)

	expected := value.NewDRange([]value.IntRange{value.NewIntRange(3, false, 5, false), value.NewIntRange(16, false, 20, false)})

	assert.Equal(t, expected, dRange.Diff(r))
}

func TestThatRangeIsSubtractedFromDRange_noIntersection(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(3, false, 7, false), value.NewIntRange(12, false, 20, false)})
	r := value.NewIntRange(9, false, 11, false)

	expected := value.NewDRange([]value.IntRange{value.NewIntRange(3, false, 7, false), value.NewIntRange(12, false, 20, false)})

	assert.Equal(t, expected, dRange.Diff(r))
}

// func TestThatDRangeIsSubtractedFromDRange(t *testing.T) {
// 	dRange1 := value.NewDRange([]value.IntRange{value.NewIntRange(3, false, 7, false), value.NewIntRange(12, false, 20, false)})
// 	dRange2 := value.NewDRange([]value.IntRange{value.NewIntRange(-2, false, 4, false), value.NewIntRange(18, false, 25, false)})

// 	expected := value.NewDRange([]value.IntRange{value.NewIntRange(5, false, 7, false), value.NewIntRange(12, false, 17, false)})

// 	assert.Equal(t, expected, dRange1.Diff(dRange2))
// }

// func TestThatDRangeIsSubtractedFromDRange_noIntersection(t *testing.T) {
// 	dRange1 := value.NewDRange([]value.IntRange{value.NewIntRange(3, false, 7, false), value.NewIntRange(12, false, 20, false)})
// 	dRange2 := value.NewDRange([]value.IntRange{value.NewIntRange(-2, false, 1, false), value.NewIntRange(8, false, 10, false)})

// 	expected := value.NewDRange([]value.IntRange{value.NewIntRange(3, false, 7, false), value.NewIntRange(12, false, 20, false)})

// 	assert.Equal(t, expected, dRange1.Diff(dRange2))
// }

func TestDRangeString(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(8, false, 12, false), value.NewIntRange(14, false, 16, false)})
	assert.Equal(t, "[[8;12][14;16]]", dRange.String())
}
