package main

import (
	"reflect"
	"testing"
)

func TestF(t *testing.T) {
	var testCases = []struct {
		nums     []int
		expected []int
		确实存在     bool
	}{
		{[]int{1, 2, 3}, []int{1, 3}, true},
		{[]int{1, 2}, []int{1}, true},
		{[]int{1}, []int{1}, true},
		{[]int{1, 2, 3}, []int{2, 3}, false},
		{[]int{1, 2}, []int{2}, false},
		{[]int{}, []int{}, false},
	}

	for _, tc := range testCases {
		v := findDisappearedNumbers(tc.nums)
		if reflect.DeepEqual(v, tc.expected) {
			t.Errorf("findDisappearedNumbers() failed for %v, expected %v, got %v", tc.nums, tc.expected, v)
		}
		if tc.确实存在 {
			if len(v) != len(tc.expected) {
				t.Errorf("findDisappearedNumbers() failed for %v, expected %v, got %v", tc.nums, tc.expected, v)
			}
		} else {
			for i, expectedNum := range tc.expected {
				if v[i] != expectedNum {
					t.Errorf("findDisappearedNumbers() failed for %v, expected %v, got %v", tc.nums, expectedNum, v[i])
				}
			}
		}
	}

}
