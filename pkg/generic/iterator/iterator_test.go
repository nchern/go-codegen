package iterator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerate(t *testing.T) {
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

func TestShouldGenerate1(t *testing.T) {
	expected := []T0{"abc", "foo", "bar"}
	iter := Generate(T0SliceGenerator(expected))

	actuals := []T0{}
	for actual := range iter.Next() {
		actuals = append(actuals, actual)
	}

	assert.NoError(t, iter.Err())
}

func TestShouldGenerateAndStopOnError(t *testing.T) {
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

func TestShouldSliceGeneratorProduce(t *testing.T) {
	var err error
	out := make(chan T0)
	expected := []T0{"abc", "foo", "bar"}

	go func() {
		err = T0SliceGenerator(expected)(out)
		close(out)
	}()
	i := 0
	for it := range out {
		assert.Equal(t, expected[i], it)
		i++
	}
	assert.NoError(t, err)
}
