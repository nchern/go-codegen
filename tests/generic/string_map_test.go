package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestStringMap(t *testing.T) {
	m := NewStringStringMap()

	m.Set("foo", "bar")

	v, found := m.Get("foo")
	if !found {
		t.Errorf("foo not found")
	}
	if v != "bar" {
		t.Errorf("foo does not correspond to bar")
	}

	m.Remove("foo")
	if _, found = m.Get("foo"); found {
		t.Errorf("foo removed but found")
	}

	m.Set("1", "a")
	m.Set("2", "b")
	m.Set("3", "c")

	keys := []string{}
	values := []string{}
	m.Each(func(k, v string) bool {
		keys = append(keys, k)
		values = append(values, v)

		return true
	})

}

func TestStringMapConcurrency(t *testing.T) {
	m := NewStringStringMapSyncronized()
	for i := 0; i < 1000; i++ {
		k := fmt.Sprintf("-%d", i)
		m.Set("k-"+k, "v-"+k)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	n := 100
	done := make(chan bool, n)
	for i := 0; i < n; i++ {
		go func() {
			wg.Wait()
			m.Each(func(string, string) bool { return true })
			done <- true
		}()
		go func(j int) {
			wg.Wait()
			m.Remove(fmt.Sprintf("%d", j))
		}(i)
	}
	wg.Done()

	for i := 0; i < n; i++ {
		<-done
	}
}
