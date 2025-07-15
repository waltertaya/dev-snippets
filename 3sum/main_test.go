package main

import "testing"

func TestThreeSome(t *testing.T) {
	var tests = []struct {
		nums []int
		expect [][]int
	} {
		{[]int{-1,0,1,2,-1,-4}, [][]int{{-1,-1,2},{-1,0,1}}},
		{[]int{0,1,1}, [][]int{}},
		{[]int{0,0,0}, [][]int{{0,0,0}}},
	}

	for _, test := range tests {
		result := threeSum(test.nums)
		if len(result) != len(test.expect) {
			t.Errorf("Expected %d results, got %d", len(test.expect), len(result))
			continue
		}

		for i, triplet := range result {
			if len(triplet) != 3 {
				t.Errorf("Expected triplet of length 3, got %d", len(triplet))
				continue
			}
			if triplet[0] != test.expect[i][0] || triplet[1] != test.expect[i][1] || triplet[2] != test.expect[i][2] {
				t.Errorf("Expected triplet %v, got %v", test.expect[i], triplet)
			}
		}
	}
}
