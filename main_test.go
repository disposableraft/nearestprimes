package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name   string
		x      int
		limit  int
		expect []int
	}{
		{"getNearestPrime for 2 with limit 1", 2, 1, []int{1, 2, 3}},
		{"getNearestPrime for 5 with limit 2", 5, 2, []int{2, 3, 5, 7, 11}},
		{"getNearestPrime for 7 with limit 3", 7, 3, []int{2, 3, 5, 7, 11, 13, 17}},
		{"getNearestPrime for 8 with limit 1", 8, 1, []int{7, 11}},
		{"getNearestPrime for 4 with limit 1", 4, 1, []int{3, 5}},
		{"getNearestPrime for 3 with limit 0", 3, 0, []int{3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getNearestPrimes(tt.x, tt.limit)

			if equal := reflect.DeepEqual(tt.expect, got); !equal {
				t.Errorf("%s failed: x=%v; limit=%v; got=%v; expected=%d", tt.name, tt.x, tt.limit, got, tt.expect)
			}
		})
	}

	t.Run("getNearestPrime for 0 with limit 0", func(t *testing.T) {
		got := getNearestPrimes(0, 0)
		if len(got) != 0 {
			t.Errorf("getNearestPrime for 0 with limit 0 failed: got=%v; expected=%v", got, []int{})
		}
	})

	t.Run("getNearestPrime for 4 with limit 0", func(t *testing.T) {
		got := getNearestPrimes(4, 0)
		if len(got) != 0 {
			t.Errorf("getNearestPrime for 4 with limit 0 failed: got=%v; expected=%v", got, []int{})
		}
	})
}
