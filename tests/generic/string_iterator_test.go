package main

import (
	"errors"
	"testing"
)

func TestStringIterator(t *testing.T) {
	list := []string{"foo", "bar", "bazz"}

	iter := Generate(func(generator chan<- string) error {
		for _, s := range list {
			generator <- s
		}
		return nil
	})

	i := 0
	for received := range iter.Next() {
		if received != list[i] {
			t.Errorf("#%d %s != %s", i, received, list[i])
		}
		i++
	}

	if i != len(list) {
		t.Errorf("list lenghtes are different: %d != %d", i, len(list))
	}
	if err := iter.Err(); err != nil {
		t.Errorf("Error from iterator not nil: %v", err)
	}
}

func TestStringIteratorReturnsError(t *testing.T) {
	list := []string{"foo", "bar", "bazz", "fuzz"}
	expectedErr := errors.New("boom")
	failAt := 2

	iter := Generate(func(generator chan<- string) error {
		for i, s := range list {
			generator <- s
			if i == failAt {
				return expectedErr
			}
		}
		return nil
	})

	i := 0
	for received := range iter.Next() {
		if received != list[i] {
			t.Errorf("#%d %s != %s", i, received, list[i])
		}
		i++
	}
	if err := iter.Err(); err != expectedErr {
		t.Errorf("Error from iterator is not as expected: %v != %v", err, expectedErr)
	}

	i--
	if i != failAt {
		t.Errorf("iterator returned more items that expected: %d != %d", i, failAt)
	}
}
