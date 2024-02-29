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
	assert.Equal(t, expected, dRange.Sect(otherRange))
}

func TestSectDRangeWithRange_simpleIntersectionMiddle1(t *testing.T) {
	ranges := []value.IntRange{value.NewIntRange(-10, false, 10, false), value.NewIntRange(20, false, 30, false)}
	dRange := value.NewDRange(ranges)

	otherRange := value.NewIntRange(0, false, 15, false)

	expected := value.NewIntRange(0, false, 10, false)
	assert.Equal(t, expected, dRange.Sect(otherRange))
}

func TestSectDRangeWithRange_simpleIntersectionMiddle2(t *testing.T) {
	ranges := []value.IntRange{value.NewIntRange(-10, false, 10, false), value.NewIntRange(20, false, 30, false)}
	dRange := value.NewDRange(ranges)

	otherRange := value.NewIntRange(15, false, 25, false)

	expected := value.NewIntRange(20, false, 25, false)
	assert.Equal(t, expected, dRange.Sect(otherRange))
}

func TestSectDRangeWithRange_simpleIntersectionRight(t *testing.T) {
	ranges := []value.IntRange{value.NewIntRange(-10, false, 10, false), value.NewIntRange(20, false, 30, false)}
	dRange := value.NewDRange(ranges)

	otherRange := value.NewIntRange(25, false, 35, false)

	expected := value.NewIntRange(25, false, 30, false)
	assert.Equal(t, expected, dRange.Sect(otherRange))
}

func TestSectDRangeWithRange_dRangeIsSubsumed(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(10, false, 14, false), value.NewIntRange(16, false, 20, false)})

	otherRange := value.NewIntRange(10, false, 20, false)

	expected := value.NewDRange([]value.IntRange{value.NewIntRange(10, false, 14, false), value.NewIntRange(16, false, 20, false)})
	assert.Equal(t, expected, dRange.Sect(otherRange))
}

func TestSectDRangeWithRange_rangeOverlapsWithMultipleDRangeRanges(t *testing.T) {
	dRange := value.NewDRange([]value.IntRange{value.NewIntRange(-10, false, 10, false), value.NewIntRange(20, false, 30, false)})

	otherRange := value.NewIntRange(-5, false, 25, false)

	expected := value.NewDRange([]value.IntRange{value.NewIntRange(-5, false, 10, false), value.NewIntRange(20, false, 25, false)})
	assert.Equal(t, expected, dRange.Sect(otherRange))
}
