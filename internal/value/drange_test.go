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

func TestDRangeString(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(8, false, 12, false), value.NewIntRange(14, false, 16, false)})
	assert.Equal(t, "[[8;12][14;16]]", dRange.String())
}
