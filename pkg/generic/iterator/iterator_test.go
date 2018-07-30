package iterator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	expected := []T0{"abc", "foo", "bar"}
	iter := Generate(func(iter chan<- T0) error {
		for _, v := range expected {
			iter <- v
		}
		return nil
	})
	actuals := []T0{}
	for actual := range iter.Next() {
		actuals = append(actuals, actual)
	}
	assert.NoError(t, iter.Err())
	assert.Equal(t, expected, actuals)

}

func TestGenerateWithError(t *testing.T) {
	expected := []T0{"abc", "foo", "bar"}
	expectedErr := errors.New("boom")
	iter := Generate(func(iter chan<- T0) error {
		for i, v := range expected {
			if i > 1 {
				return expectedErr
			}
			iter <- v
		}
		return nil
	})
	actuals := []T0{}
	for actual := range iter.Next() {
		actuals = append(actuals, actual)
	}
	assert.Equal(t, expectedErr, iter.Err())
	assert.Equal(t, expected[0:2], actuals)
}
