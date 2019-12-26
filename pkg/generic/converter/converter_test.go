package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceConversion(t *testing.T) {
	given := []T0{"foo", "bar", "bazz"}
	expected := []T1{3, 3, 4}

	actual := T0Slice(given).ToT1Slice(func(x T0) T1 {
		return len(x)
	})
	assert.Equal(t, expected, actual)
}
