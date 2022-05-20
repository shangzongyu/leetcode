package main

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	tables := []struct {
		target   int
		want_one int
		want_two int
	}{
		{9, 0, 1},
		{18, 1, 2},
		{26, 2, 3},
	}

	for _, table := range tables {
		result := twoSum(nums, table.target)
		if result == nil {
			t.Errorf("twoSum(%v, %d) == nil", nums, table.target)
		} else {
			if result[0] != table.want_one || result[1] != table.want_two {
				t.Errorf("twoSum(%v, %d) == %v, want %v", nums, table.target, result, []int{table.want_one, table.want_two})
			}
		}
	}
}
