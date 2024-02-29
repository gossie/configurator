package value_test

import (
	"testing"

	"github.com/gossie/configurator/internal/value"
	"github.com/stretchr/testify/assert"
)

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

func TestDRangeString(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(8, false, 12, false), value.NewIntRange(14, false, 16, false)})
	assert.Equal(t, "[[8;12][14;16]]", dRange.String())
}
