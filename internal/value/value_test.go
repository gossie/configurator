package value_test

import (
	"testing"

	"github.com/gossie/configurator/internal/value"
	"github.com/stretchr/testify/assert"
)

func TestSectDRangeWithRange_simpleIntersectionLeft(t *testing.T) {
	ranges := []value.IntRange{value.NewIntRange(-10, false, 10, false), value.NewIntRange(20, false, 30, false)}
	dRange := value.NewDRange(ranges)

	otherRange := value.NewIntRange(-15, false, 5, false)

	expected := value.NewIntRange(-10, false, 5, false)
	assert.Equal(t, expected, value.SectRangeWithDRange(otherRange, dRange))
}

func TestSectDRangeWithRange_simpleIntersectionMiddle1(t *testing.T) {
	ranges := []value.IntRange{value.NewIntRange(-10, false, 10, false), value.NewIntRange(20, false, 30, false)}
	dRange := value.NewDRange(ranges)

	otherRange := value.NewIntRange(0, false, 15, false)

	expected := value.NewIntRange(0, false, 10, false)
	assert.Equal(t, expected, value.SectRangeWithDRange(otherRange, dRange))
}

func TestSectDRangeWithRange_simpleIntersectionMiddle2(t *testing.T) {
	ranges := []value.IntRange{value.NewIntRange(-10, false, 10, false), value.NewIntRange(20, false, 30, false)}
	dRange := value.NewDRange(ranges)

	otherRange := value.NewIntRange(15, false, 25, false)

	expected := value.NewIntRange(20, false, 25, false)
	assert.Equal(t, expected, value.SectRangeWithDRange(otherRange, dRange))
}

func TestSectDRangeWithRange_simpleIntersectionRight(t *testing.T) {
	ranges := []value.IntRange{value.NewIntRange(-10, false, 10, false), value.NewIntRange(20, false, 30, false)}
	dRange := value.NewDRange(ranges)

	otherRange := value.NewIntRange(25, false, 35, false)

	expected := value.NewIntRange(25, false, 30, false)
	assert.Equal(t, expected, value.SectRangeWithDRange(otherRange, dRange))
}

func TestSectDRangeWithRange_dRangeIsSubsumed(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(10, false, 14, false), value.NewIntRange(16, false, 20, false)})

	otherRange := value.NewIntRange(10, false, 20, false)

	expected := value.NewDRange([]value.IntRange{value.NewIntRange(10, false, 14, false), value.NewIntRange(16, false, 20, false)})
	assert.Equal(t, expected, value.SectRangeWithDRange(otherRange, dRange))
}

func TestSectDRangeWithRange_rangeOverlapsWithMultipleDRangeRanges(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(-10, false, 10, false), value.NewIntRange(20, false, 30, false)})

	otherRange := value.NewIntRange(-5, false, 25, false)

	expected := value.NewDRange([]value.IntRange{value.NewIntRange(-5, false, 10, false), value.NewIntRange(20, false, 25, false)})
	assert.Equal(t, expected, value.SectRangeWithDRange(otherRange, dRange))
}

func TestSectRangeWithSet(t *testing.T) {
	intRange := value.NewIntRange(7, false, 17, false)

	set := value.NewIntValues([]int{5, 9, 15, 17, 19})

	expected := value.NewIntValues([]int{9, 15, 17})
	assert.Equal(t, expected, value.SectRangeWithSet(intRange, set))
}

func TestSectDRangeWithSet(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(7, false, 10, false), value.NewIntRange(20, false, 30, false)})

	set := value.NewIntValues([]int{5, 9, 15, 17, 19, 25, 30, 35})

	expected := value.NewIntValues([]int{9, 25, 30})
	assert.Equal(t, expected, value.SectDRangeWithSet(dRange, set))
}

func TestSectSetWithSet(t *testing.T) {
	set1 := value.NewIntValues([]int{5, 9, 15, 17, 19, 25, 30, 35})
	set2 := value.NewIntValues([]int{4, 9, 16, 17, 18, 25, 31, 37})

	expected := value.NewIntValues([]int{9, 17, 25})
	assert.Equal(t, expected, value.SectSetWithSet(set1, set2))
}

func TestSectRangeWithRange1(t *testing.T) {
	range1 := value.NewIntRange(1, false, 10, false)
	range2 := value.NewIntRange(-5, false, 5, false)

	expected := value.NewIntRange(1, false, 5, false)
	actual, err := value.SectRangeWithRange(range1, range2)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestSectRangeWithRange2(t *testing.T) {
	range1 := value.NewIntRange(1, false, 10, false)
	range2 := value.NewIntRange(5, false, 15, false)

	expected := value.NewIntRange(5, false, 10, false)
	actual, err := value.SectRangeWithRange(range1, range2)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestSectRangeWithRange_equal(t *testing.T) {
	range1 := value.NewIntRange(1, false, 10, false)
	range2 := value.NewIntRange(1, false, 10, false)

	expected := value.NewIntRange(1, false, 10, false)
	actual, err := value.SectRangeWithRange(range1, range2)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestSectRangeWithRange_oneSubsumesTheOther(t *testing.T) {
	range1 := value.NewIntRange(1, false, 10, false)
	range2 := value.NewIntRange(3, false, 7, false)

	expected := value.NewIntRange(3, false, 7, false)
	actual, err := value.SectRangeWithRange(range1, range2)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestSectRangeWithRange_noResult(t *testing.T) {
	range1 := value.NewIntRange(1, false, 10, false)
	range2 := value.NewIntRange(11, false, 20, false)

	_, err := value.SectRangeWithRange(range1, range2)

	assert.Error(t, err)
}

func TestSectDrangeWithDRange_oneIntersectionEach(t *testing.T) {
	dRange1 := value.NewDRange([]value.IntRange{value.NewIntRange(1, false, 10, false), value.NewIntRange(20, false, 30, false)})
	dRange2 := value.NewDRange([]value.IntRange{value.NewIntRange(-5, false, 5, false), value.NewIntRange(25, false, 35, false)})

	expected := value.NewDRange([]value.IntRange{value.NewIntRange(1, false, 5, false), value.NewIntRange(25, false, 30, false)})
	actual, err := value.SectDRangeWithDRange(dRange1, dRange2)

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestSectDrangeWithDRange_intersectionWithWithTwoRanges(t *testing.T) {
	dRange1 := value.NewDRange([]value.IntRange{value.NewIntRange(1, false, 10, false), value.NewIntRange(20, false, 30, false)})
	dRange2 := value.NewDRange([]value.IntRange{value.NewIntRange(5, false, 21, false), value.NewIntRange(25, false, 35, false)})

	expected := value.NewDRange([]value.IntRange{value.NewIntRange(5, false, 10, false), value.NewIntRange(20, false, 21, false), value.NewIntRange(25, false, 30, false)})
	actual, err := value.SectDRangeWithDRange(dRange1, dRange2)

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestSectDrangeWithDRange_oneRangeSubsumesAnother(t *testing.T) {
	dRange1 := value.NewDRange([]value.IntRange{value.NewIntRange(1, false, 10, false), value.NewIntRange(20, false, 30, false)})
	dRange2 := value.NewDRange([]value.IntRange{value.NewIntRange(3, false, 7, false), value.NewIntRange(25, false, 35, false)})

	expected := value.NewDRange([]value.IntRange{value.NewIntRange(3, false, 7, false), value.NewIntRange(25, false, 30, false)})
	actual, err := value.SectDRangeWithDRange(dRange1, dRange2)

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestSectDrangeWithDRange_noIntersection(t *testing.T) {
	dRange1 := value.NewDRange([]value.IntRange{value.NewIntRange(1, false, 10, false), value.NewIntRange(20, false, 30, false)})
	dRange2 := value.NewDRange([]value.IntRange{value.NewIntRange(12, false, 15, false), value.NewIntRange(35, false, 40, false)})

	_, err := value.SectDRangeWithDRange(dRange1, dRange2)
	assert.Error(t, err)
}
