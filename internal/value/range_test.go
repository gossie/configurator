package value_test

import (
	"testing"

	"github.com/gossie/configurator/internal/value"
	"github.com/stretchr/testify/assert"
)

func TestThatRangeSubsumesSet(t *testing.T) {
	r := value.NewIntRange(7, false, 17, false)
	set := value.NewIntValues([]int{7, 10, 17})

	assert.True(t, r.Subsumes(set))
}

func TestThatRangeDoesNotSubsumeSet(t *testing.T) {
	r := value.NewIntRange(7, false, 17, false)
	set := value.NewIntValues([]int{6, 7, 10, 17})

	assert.False(t, r.Subsumes(set))
}

func TestThatRangeSubsumesRange(t *testing.T) {
	r1 := value.NewIntRange(7, false, 17, false)
	r2 := value.NewIntRange(8, false, 16, false)

	assert.True(t, r1.Subsumes(r2))
}

func TestThatRangeSubsumesRange_equal(t *testing.T) {
	r1 := value.NewIntRange(7, false, 17, false)
	r2 := value.NewIntRange(7, false, 17, false)

	assert.True(t, r1.Subsumes(r2))
}

func TestThatRangeDoesNotSubsumeRange(t *testing.T) {
	r1 := value.NewIntRange(7, false, 17, false)
	r2 := value.NewIntRange(8, false, 18, false)

	assert.False(t, r1.Subsumes(r2))
}

func TestThatRangeSubsumesDRange(t *testing.T) {
	r1 := value.NewIntRange(7, false, 17, false)
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(8, false, 12, false), value.NewIntRange(14, false, 16, false)})

	assert.True(t, r1.Subsumes(dRange))
}

func TestThatRangeDoesNotSubsumeDRange(t *testing.T) {
	r1 := value.NewIntRange(7, false, 17, false)
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(8, false, 12, false), value.NewIntRange(14, false, 18, false)})

	assert.False(t, r1.Subsumes(dRange))
}

func TestRangeString_closed(t *testing.T) {
	r := value.NewIntRange(7, false, 17, false)
	assert.Equal(t, "[7;17]", r.String())
}

func TestRangeString_lowerOpen(t *testing.T) {
	r := value.NewIntRange(7, true, 17, false)
	assert.Equal(t, "]7;17]", r.String())
}

func TestRangeString_upperOpen(t *testing.T) {
	r := value.NewIntRange(7, false, 17, true)
	assert.Equal(t, "[7;17[", r.String())
}

func TestRangeString_bothOpen(t *testing.T) {
	r := value.NewIntRange(7, true, 17, true)
	assert.Equal(t, "]7;17[", r.String())
}
