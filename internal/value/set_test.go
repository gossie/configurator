package value_test

import (
	"testing"

	"github.com/gossie/configurator/internal/value"
	"github.com/stretchr/testify/assert"
)

func TestSetString(t *testing.T) {
	set := value.NewIntValues([]int{1, 2, 3, 4, 5, 6, 7})
	assert.Equal(t, "{1,2,3,4,5,6,7}", set.String())
}
