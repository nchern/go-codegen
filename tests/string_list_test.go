package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"testing"
)

func TestStringList(t *testing.T) {
	original := []string{"foo", "bar", "buzz"}
	l := NewStringListFromSlice(original)

	i := 0
	for s := range l.Iter() {
		if original[i] != s {
			t.Errorf("Expected %s at %d in %v", s, i, original)
		}
		i++
	}

	if !l.Any(func(s string) bool { return s == "foo" }) {
		t.Errorf("Any() must find 'foo'")
	}
	if l.Any(func(s string) bool { return s == "foobar" }) {
		t.Errorf("Any() must NOT find 'foobar'")
	}

	filtered := l.Filter(func(s string) bool { return strings.HasPrefix(s, "b") })
	if filtered.Get(0) != "bar" {
		t.Errorf("Expected 'bar' at 0 in %v", filtered)
	}
	if filtered.Get(1) != "buzz" {
		t.Errorf("Expected 'buzz' at 1 in %v", filtered)
	}

	l.Append("brr")
	if l.Get(3) != "brr" {
		t.Errorf("Expected 'brr' at 3 in %v", l)
	}
	if l.Pop("default") != "brr" {
		t.Errorf("Expected pop 'brr' out %v", l)
	}
	if l.Len() != 3 {
		t.Errorf("Expected have 3 elements after pop %v", l)
	}

	mustBeSorted := []string{}
	l.ByFunc(func(s1, s2 string) bool { return s1 < s2 }).Sort()
	for s := range l.Iter() {
		mustBeSorted = append(mustBeSorted, s)
	}
	if !sort.StringsAreSorted(mustBeSorted) {
		t.Errorf("%v must be sorted", mustBeSorted)
	}

}

func TestStringListConcurrency(t *testing.T) {
	l := NewStringList()
	for i := 0; i < 1000; i++ {
		l.Append(fmt.Sprintf("%d", i))
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	n := 100
	done := make(chan bool, n)
	for i := 0; i < n; i++ {
		go func() {
			wg.Wait()
			l.ByFunc(func(s1, s2 string) bool { return s1 < s2 }).Sort()
			done <- true
		}()
		go func(j int) {
			wg.Wait()
			l.Pop("")
			done <- true
		}(i)
	}
	wg.Done()

	for i := 0; i < 2*n; i++ {
		<-done
	}

	if l.Len() != 900 {
		t.Errorf("Expected to have 900 elements instead of %d", l.Len())
	}

}
