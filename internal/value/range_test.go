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

func TestThatSetIsSubtractedFromRange_lowerBound(t *testing.T) {
	r := value.NewIntRange(7, false, 17, false)
	set := value.NewIntValues([]int{7})

	expected := value.NewIntRange(8, false, 17, false)

	assert.Equal(t, expected, r.Diff(set))
}

func TestThatSetIsSubtractedFromRange_upperBound(t *testing.T) {
	r := value.NewIntRange(7, false, 17, false)
	set := value.NewIntValues([]int{17})

	expected := value.NewIntRange(7, false, 16, false)

	assert.Equal(t, expected, r.Diff(set))
}

func TestThatSetIsSubtractedFromRange_noIntersection(t *testing.T) {
	r := value.NewIntRange(7, false, 17, false)
	set := value.NewIntValues([]int{3, 19})

	expected := value.NewIntRange(7, false, 17, false)

	assert.Equal(t, expected, r.Diff(set))
}

func TestThatSetIsSubtractedFromRange(t *testing.T) {
	r := value.NewIntRange(7, false, 17, false)
	set := value.NewIntValues([]int{3, 9, 15, 19})

	expected := value.NewDRange([]value.IntRange{value.NewIntRange(7, false, 8, false), value.NewIntRange(10, false, 14, false), value.NewIntRange(16, false, 17, false)})

	assert.Equal(t, expected, r.Diff(set))
}

func TestThatRangeIsSubtractedFromRange_Subsume(t *testing.T) {
	r1 := value.NewIntRange(7, false, 17, false)
	r2 := value.NewIntRange(9, false, 12, false)

	expected := value.NewDRange([]value.IntRange{value.NewIntRange(7, false, 8, false), value.NewIntRange(13, false, 17, false)})

	assert.Equal(t, expected, r1.Diff(r2))
}

func TestThatRangeIsSubtractedFromRange_SubsumeIncludingLowerBound(t *testing.T) {
	r1 := value.NewIntRange(7, false, 17, false)
	r2 := value.NewIntRange(7, false, 12, false)

	expected := value.NewIntRange(13, false, 17, false)

	assert.Equal(t, expected, r1.Diff(r2))
}

func TestThatRangeIsSubtractedFromRange_SubsumeIncludingUpperBound(t *testing.T) {
	r1 := value.NewIntRange(7, false, 17, false)
	r2 := value.NewIntRange(10, false, 17, false)

	expected := value.NewIntRange(7, false, 9, false)

	assert.Equal(t, expected, r1.Diff(r2))
}

func TestThatRangeIsSubtractedFromRange_lowerIntersection(t *testing.T) {
	r1 := value.NewIntRange(7, false, 17, false)
	r2 := value.NewIntRange(3, false, 10, false)

	expected := value.NewIntRange(11, false, 17, false)

	assert.Equal(t, expected, r1.Diff(r2))
}

func TestThatRangeIsSubtractedFromRange_upperIntersection(t *testing.T) {
	r1 := value.NewIntRange(7, false, 17, false)
	r2 := value.NewIntRange(15, false, 20, false)

	expected := value.NewIntRange(7, false, 14, false)

	assert.Equal(t, expected, r1.Diff(r2))
}

func TestThatRangeIsSubtractedFromRange_noIntersection(t *testing.T) {
	r1 := value.NewIntRange(7, false, 17, false)
	r2 := value.NewIntRange(20, false, 30, false)

	expected := value.NewIntRange(7, false, 17, false)

	assert.Equal(t, expected, r1.Diff(r2))
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
